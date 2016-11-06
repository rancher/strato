package inspect

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	log "github.com/Sirupsen/logrus"
	"github.com/heroku/docker-registry-client/registry"
)

func Action(c *cli.Context) error {
	image := c.Args()[0]
	registryURL := c.GlobalString("registry")

	hub, err := registry.New(registryURL, "", "")
	if err != nil {
		return err
	}

	tags, err := hub.Tags(image)
	if err != nil {
		return err
	}
	log.Infof("Tags: %s", tags)

	manifest, err := hub.Manifest(image, "latest")
	if err != nil {
		return err
	}
	for _, layer := range manifest.FSLayers {
		split := strings.Split(fmt.Sprint(layer.BlobSum), ":")[1]
		log.Infoln(split)
	}
	return nil
}
