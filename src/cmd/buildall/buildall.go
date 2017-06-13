package buildall

import (
	"fmt"
	"os"
	"path"

	"github.com/rancher/strato/src/build"
	"github.com/rancher/strato/src/buildorder"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
)

// Command definition
var Command = cli.Command{
	Name:     "build-all",
	Usage:    "Managing building a collection of packages",
	HideHelp: true,
	Action:   buildAllAction,
	Flags:    []cli.Flag{},
}

func buildAllAction(c *cli.Context) error {
	inDir := c.Args().Get(0)
	if _, err := os.Stat(inDir); os.IsNotExist(err) {
		return err
	}
	outDir, err := utils.GetOutDir(c.Args().Get(1))
	if err != nil {
		return err
	}

	buildOrder, err := buildorder.Get(inDir)
	if err != nil {
		return err
	}

	fmt.Println("Build order")
	for _, pkg := range buildOrder {
		fmt.Println(pkg)
	}

	for _, pkg := range buildOrder {
		if err = build.Build(path.Join(inDir, pkg), outDir); err != nil {
			return err
		}
	}

	return nil
}
