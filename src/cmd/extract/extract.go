package extract

import (
	"github.com/rancher/strato/src/extract"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
)

type info struct {
	Layers []string `json:"Layers"`
}

var Command = cli.Command{
	Name:            "extract",
	Usage:           "Extract the last layer from the built image",
	HideHelp:        true,
	SkipFlagParsing: true,
	Action:          Action,
}

func Action(c *cli.Context) error {
	inDir := c.Args().Get(0)
	outDir, err := utils.GetOutDir(c.Args().Get(1))
	if err != nil {
		return err
	}
	return extract.Extract(inDir, outDir)
}
