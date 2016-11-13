package add

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/urfave/cli"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/distribution/digest"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/joshwget/strato/utils"
)

func Action(c *cli.Context) error {
	registryURL := c.GlobalString("registry")
	dir := c.String("dir")

	hub, err := registry.New(registryURL, "", "")
	if err != nil {
		return err
	}
	hub.Logf = registry.Quiet

	var skip *regexp.Regexp
	if c.String("skip") != "" {
		skip, err = regexp.Compile(c.String("skip"))
		if err != nil {
			return err
		}
	}

	var installs sync.WaitGroup
	for _, image := range c.Args() {
		installs.Add(1)
		go func(image string) {
			defer installs.Done()
			if err = add(hub, dir, skip, image); err != nil {
				panic(err)
			}
		}(image)
	}
	installs.Wait()

	return nil
}

func add(hub *registry.Registry, dir string, skip *regexp.Regexp, images ...string) error {
	for _, image := range images {
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
			reader.Close()

			for _, dependency := range pkg.Dependencies {
				if err = add(hub, dir, skip, dependency); err != nil {
					return err
				}
			}

			reader, err = hub.DownloadLayer(image, digest)
			if err != nil {
				return err
			}

			log.Infof("Installing package %s", image)
			if err = utils.ExtractTar(reader, dir, skip); err != nil {
				return err
			}
			reader.Close()
		}
	}

	return nil
}
