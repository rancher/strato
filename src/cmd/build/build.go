package build

import (
	"github.com/rancher/strato/src/build"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
)

// Command definition
var Command = cli.Command{
	Name:     "build",
	Usage:    "Build image from source and extract the last layer from the built image",
	HideHelp: true,
	Action:   buildAction,
	Flags:    []cli.Flag{},
}

// Build the package from Dockerfile, strato.yml and prebuild.sh
func buildAction(c *cli.Context) error {
	inDir := c.Args().Get(0)
	outDir, err := utils.GetOutDir(c.Args().Get(1))
	if err != nil {
		return err
	}
	return build.Build(inDir, outDir)
}
