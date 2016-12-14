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
	user := c.GlobalString("user")
	dir := c.String("dir")

	hub, err := registry.New(registryURL, "", "")
	if err != nil {
		return err
	}
	hub.Logf = registry.Quiet

	var installs sync.WaitGroup
	for _, image := range c.Args() {
		if !strings.Contains(image, "/") {
			image = fmt.Sprintf("%s/%s", user, image)
		}
		installs.Add(1)
		go func(image string) {
			defer installs.Done()
			if err = add(hub, dir, image); err != nil {
				panic(err)
			}
		}(image)
	}
	installs.Wait()

	return nil
}

func add(hub *registry.Registry, dir string, images ...string) error {
	for _, image := range images {
		var subpackage string
		imageSplit := strings.SplitN(image, "%", 2)
		if len(imageSplit) > 1 {
			image = imageSplit[0]
			subpackage = imageSplit[1]
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
			reader.Close()

			for _, dependency := range pkg.Dependencies {
				if err = add(hub, dir, dependency); err != nil {
					return err
				}
			}

			var whitelist []*regexp.Regexp
			var blacklist []*regexp.Regexp
			if whitelistItems, ok := pkg.Subpackages[subpackage]; ok {
				// Only install whitelisted for subpackages
				for _, whitelistItem := range whitelistItems {
					whitelistRegex, err := regexp.Compile(whitelistItem)
					if err != nil {
						return err
					}
					whitelist = append(whitelist, whitelistRegex)
				}
			} else {
				// Blacklist the union of all subpackage whitelists for regular packages
				var union []*regexp.Regexp
				for _, whitelistItems := range pkg.Subpackages {
					for _, whitelistItem := range whitelistItems {
						whitelistRegex, err := regexp.Compile(whitelistItem)
						if err != nil {
							return err
						}
						union = append(union, whitelistRegex)
					}
				}
				blacklist = union
			}
			for _, exclude := range pkg.Exclude {
				excludeRegex, err := regexp.Compile(exclude)
				if err != nil {
					return err
				}
				blacklist = append(blacklist, excludeRegex)
			}

			reader, err = hub.DownloadLayer(image, digest)
			if err != nil {
				return err
			}

			log.Infof("Installing package %s", image)
			if err = utils.ExtractTar(reader, dir, whitelist, blacklist); err != nil {
				return err
			}
			reader.Close()
		}
	}

	return nil
}
