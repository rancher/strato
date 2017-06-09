package extract

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/docker/docker/client"
	"github.com/rancher/strato/src/config"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type info struct {
	Layers []string `json:"Layers"`
}

func Action(c *cli.Context) error {
	inDir := c.Args().Get(0)
	outDir, err := utils.GetOutDir(c.Args().Get(1))
	if err != nil {
		return err
	}
	return Extract(inDir, outDir)
}

func Extract(inDir, outDir string) error {
	configPath := path.Join(inDir, "strato.yml")

	packageName := path.Base(inDir)
	if strings.Contains(packageName, ".") {
		packageName = strings.SplitN(packageName, ".", 2)[1]
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var pkg config.Package
	if err := yaml.Unmarshal(b, &pkg); err != nil {
		return err
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	reader, err := cli.ImageSave(context.Background(), []string{"build/" + packageName})
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := utils.TarForEach(reader, nil, nil, func(tarReader io.Reader, header *tar.Header) error {
		if header.Name == "manifest.json" {
			io.Copy(buf, tarReader)
		}
		return nil
	}); err != nil {
		return err
	}

	var infos []info
	if err := json.Unmarshal(buf.Bytes(), &infos); err != nil {
		return err
	}

	layers := infos[0].Layers
	layer := layers[len(layers)-1]

	reader.Close()

	reader, err = cli.ImageSave(context.Background(), []string{"build/" + packageName})
	if err != nil {
		return err
	}

	buf = new(bytes.Buffer)
	if err := utils.TarForEach(reader, nil, nil, func(tarReader io.Reader, header *tar.Header) error {
		if header.Name == layer {
			io.Copy(buf, tarReader)
			w, err := os.Create(filepath.Join(outDir, packageName+"-all.tar"))
			if err != nil {
				panic(err)
			}
			defer w.Close()

			layerReader := bytes.NewReader(buf.Bytes())
			_, err = io.Copy(w, layerReader)
			if err != nil {
				panic(err)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(outDir, packageName+".extractlog"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	logFile := bufio.NewWriter(f)
	defer func() {
		logFile.Flush()
		f.Close()
	}()

	b = buf.Bytes()
	if err = generatePackage(b, outDir, packageName, &pkg, logFile); err != nil {
		return err
	}
	for subpackageName := range pkg.Subpackages {
		if err = generatePackage(b, outDir, subpackageName, &pkg, logFile); err != nil {
			return err
		}
	}

	return nil
}

func generatePackage(b []byte, outDir, name string, pkg *config.Package, logFile *bufio.Writer) error {
	// TODO: make the default package code more obvious
	whitelist, blacklist, err := config.GenerateWhiteAndBlackLists(pkg, name)
	if err != nil {
		return err
	}

	tgzFileName := path.Join(outDir, name) + ".tar.gz"
	f, err := os.Create(tgzFileName)
	if err != nil {
		return err
	}
	gzipWriter := gzip.NewWriter(f)
	packageWriter := tar.NewWriter(gzipWriter)

	layerReader := bytes.NewReader(b)
	if err := utils.TarForEach(layerReader, whitelist, blacklist, func(tarReader io.Reader, header *tar.Header) error {
		fmt.Printf("%s | %s\n", name, header.Name)
		fmt.Fprintf(logFile, "%s | %s\n", name, header.Name)
		packageWriter.WriteHeader(header)
		buf := new(bytes.Buffer)
		io.Copy(buf, tarReader)
		packageWriter.Write(buf.Bytes())
		return nil
	}); err != nil {
		return err
	}

	packageWriter.Close()
	gzipWriter.Close()
	f.Close()

	cmd := exec.Command("sh", "-c", "zcat "+tgzFileName+" | docker import - stratopkg/"+name)
	fmt.Printf("Running: %v\n", cmd.Args)
	fmt.Fprintf(logFile, "Running: %v\n", cmd.Args)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
