package inspect

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/distribution/digest"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/joshwget/strato/utils"
)

func Action(c *cli.Context) error {
	image := c.Args()[0]
	registryURL := c.GlobalString("registry")

	hub, err := registry.New(registryURL, "", "")
	if err != nil {
		return err
	}

	manifest, err := hub.Manifest(image, "latest")
	if err != nil {
		return err
	}

	layers := []string{}
	for _, layer := range manifest.FSLayers {
		split := strings.Split(fmt.Sprint(layer.BlobSum), ":")[1]
		layers = append(layers, split)
	}

	for _, layer := range layers {
		digest := digest.NewDigestFromHex(
			"sha256",
			layer,
		)
		reader, err := hub.DownloadLayer(image, digest)
		if err != nil {
			return err
		}
		pkg, err := utils.FindPackage(reader)
		if err != nil {
			return err
		}
		if pkg == nil {
			continue
		}
		log.Infof("License: %s", pkg.License)
		log.Infof("Version: %s", pkg.Version)
		log.Infof("License: %s", pkg.Description)
		log.Infof("Dependencies: %s", strings.Join(pkg.Dependencies, ","))

		reader, err = hub.DownloadLayer(image, digest)
		if err != nil {
			return err
		}

		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return err
		}
		tarReader := tar.NewReader(gzipReader)
		for {
			header, err := tarReader.Next()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			log.Infoln(header.Name)
		}

		reader.Close()
	}
	return nil
}
