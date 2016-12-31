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
	b, err := ioutil.ReadFile(os.Args[1])
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

	whitelist, blacklist, err := config.GenerateWhiteAndBlackLists(&pkg, "")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	gzipWriter := gzip.NewWriter(f)
	packageWriter := tar.NewWriter(gzipWriter)

	layerReader := bytes.NewReader(buf.Bytes())
	if err := utils.TarForEach(layerReader, whitelist, blacklist, func(tarReader io.Reader, header *tar.Header) error {
		fmt.Println(header.Name)
		packageWriter.WriteHeader(header)
		buf = new(bytes.Buffer)
		io.Copy(buf, tarReader)
		packageWriter.Write(buf.Bytes())
		return nil
	}); err != nil {
		panic(err)
	}

	packageWriter.Close()
	gzipWriter.Close()
	f.Close()
}
