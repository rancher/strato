package xf

import (
	"os"

	"github.com/urfave/cli"

	"github.com/rancher/strato/src/utils"
)

var Command = cli.Command{
	Name:            "xf",
	HideHelp:        true,
	SkipFlagParsing: true,
	Action:          Action,
	Hidden:          true,
}

func Action(c *cli.Context) error {
	filename := c.Args().Get(0)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return utils.ExtractTar(f, "/", nil, nil)
}
