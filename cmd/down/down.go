package down

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/docker/distribution/digest"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/joshwget/lay/utils"
)

func Action(c *cli.Context) error {
	image := c.Args()[0]
	registryURL := c.GlobalString("registry")
	dir := c.String("dir")

	hub, err := registry.New(registryURL, "", "")
	if err != nil {
		return err
	}
	hub.Logf = registry.Quiet

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
		isPackage, err := utils.IsPackage(reader)
		if err != nil {
			return err
		}
		if !isPackage {
			continue
		}
		reader.Close()

		reader, err = hub.DownloadLayer(image, digest)
		if err != nil {
			return err
		}
		if err = utils.ExtractTar(reader, dir); err != nil {
			return err
		}
		reader.Close()
	}

	return nil
}
