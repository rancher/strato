package inspect

import "github.com/urfave/cli"

var Command = cli.Command{
	Name:            "inspect",
	HideHelp:        true,
	SkipFlagParsing: true,
	Action:          Action,
}

func Action(c *cli.Context) error {
	/*user := c.GlobalString("user")
	image := c.Args()[0]

	if !strings.Contains(image, "/") {
		image = fmt.Sprintf("%s/%s", user, image)
	}

	var subpackage string
	imageSplit := strings.SplitN(image, "%", 2)
	if len(imageSplit) > 1 {
		image = imageSplit[0]
		subpackage = imageSplit[1]
	}
	registryURL := c.GlobalString("registry")

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

		whitelist, blacklist, err := config.GenerateWhiteAndBlackLists(pkg, subpackage)
		if err != nil {
			return err
		}

		if err := utils.GzipTarForEach(reader, whitelist, blacklist, func(tarReader io.Reader, header *tar.Header) error {
			log.Infoln(header.Name)
			return nil
		}); err != nil {
			return err
		}

		reader.Close()
	}*/
	return nil
}
