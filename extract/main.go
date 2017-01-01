package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/docker/docker/client"
	"github.com/joshwget/strato/config"
	"github.com/joshwget/strato/utils"
	"gopkg.in/yaml.v2"
)

const (
	imageName = "package"
)

type info struct {
	Layers []string `json:"Layers"`
}

func main() {
	inDir := os.Args[1]
	outDir := os.Args[2]
	packageName := path.Base(inDir)
	configPath := path.Join(inDir, "strato.yml")

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var pkg config.Package
	if err := yaml.Unmarshal(b, &pkg); err != nil {
		panic(err)
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImageSave(context.Background(), []string{imageName})
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := utils.TarForEach(reader, nil, nil, func(tarReader io.Reader, header *tar.Header) error {
		if header.Name == "manifest.json" {
			io.Copy(buf, tarReader)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var infos []info
	if err := json.Unmarshal(buf.Bytes(), &infos); err != nil {
		panic(err)
	}

	layers := infos[0].Layers
	layer := layers[len(layers)-1]

	reader.Close()

	reader, err = cli.ImageSave(context.Background(), []string{imageName})
	if err != nil {
		panic(err)
	}

	buf = new(bytes.Buffer)
	if err := utils.TarForEach(reader, nil, nil, func(tarReader io.Reader, header *tar.Header) error {
		if header.Name == layer {
			io.Copy(buf, tarReader)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	b = buf.Bytes()
	if err = generatePackage(b, outDir, packageName, &pkg); err != nil {
		panic(err)
	}
	for subpackageName := range pkg.Subpackages {
		if err = generatePackage(b, outDir, subpackageName, &pkg); err != nil {
			panic(err)
		}
	}
}

func generatePackage(b []byte, outDir, name string, pkg *config.Package) error {
	// TODO: make the default package code more obvious
	whitelist, blacklist, err := config.GenerateWhiteAndBlackLists(pkg, name)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(outDir, name) + ".tar.gz")
	if err != nil {
		return err
	}
	gzipWriter := gzip.NewWriter(f)
	packageWriter := tar.NewWriter(gzipWriter)

	layerReader := bytes.NewReader(b)
	if err := utils.TarForEach(layerReader, whitelist, blacklist, func(tarReader io.Reader, header *tar.Header) error {
		fmt.Printf("%s | %s\n", name, header.Name)
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

	return nil
}
