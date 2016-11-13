package inspect

import (
	"fmt"
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
		reader.Close()
		if pkg != nil {
			log.Infof("License: %s", pkg.License)
			log.Infof("Version: %s", pkg.Version)
			log.Infof("License: %s", pkg.Description)
			log.Infof("Dependencies: %s", strings.Join(pkg.Dependencies, ","))
			return nil
		}
	}
	return nil
}
