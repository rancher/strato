package add

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/urfave/cli"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/distribution/digest"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/joshwget/strato/config"
	"github.com/joshwget/strato/utils"
	"github.com/joshwget/strato/version"
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

	fmt.Println(utils.Size/1000000.0, "mb")

	return nil
}

func add(hub *registry.Registry, dir string, images ...string) error {
	for _, image := range images {
		var subpackage string
		imageSplit := strings.SplitN(image, "#", 2)
		if len(imageSplit) > 1 {
			image = imageSplit[0]
			subpackage = imageSplit[1]
		}

		manifest, err := hub.Manifest(image, version.Tag)
		if err != nil {
			return err
		}

		layers := []string{}
		for _, layer := range manifest.FSLayers {
			split := strings.Split(fmt.Sprint(layer.BlobSum), ":")[1]
			layers = append([]string{split}, layers...)
		}

		var packageLayer string
		var pkg *config.Package
		for i, layer := range layers {
			digest := digest.NewDigestFromHex(
				"sha256",
				layer,
			)
			reader, err := hub.DownloadLayer(image, digest)
			if err != nil {
				return err
			}
			pkg, err = utils.FindPackage(reader)
			if err != nil {
				return err
			}
			reader.Close()
			if pkg != nil {
				packageLayer = layers[i+1]
				break
			}
		}

		for _, dependency := range pkg.Dependencies {
			if err = add(hub, dir, dependency); err != nil {
				return err
			}
		}

		whitelist, blacklist, err := config.GenerateWhiteAndBlackLists(pkg, subpackage)
		if err != nil {
			return err
		}

		digest := digest.NewDigestFromHex(
			"sha256",
			packageLayer,
		)

		reader, err := hub.DownloadLayer(image, digest)
		if err != nil {
			return err
		}

		log.Infof("Installing package %s", fmt.Sprintf("%s:%s", image, version.Tag))
		if err = utils.ExtractTar(reader, dir, whitelist, blacklist); err != nil {
			return err
		}
		reader.Close()

		if pkg.Postcmd != "" {
			log.Infof("Running command %s", pkg.Postcmd)
			cmd := exec.Command("sh", "-c", pkg.Postcmd)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}

	return nil
}
