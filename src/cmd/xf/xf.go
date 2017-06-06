package xf

import (
	"os"

	"github.com/urfave/cli"

	"github.com/rancher/strato/src/utils"
)

func Action(c *cli.Context) error {
	filename := c.Args()[0]

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return utils.ExtractTar(f, "/", nil, nil)
}
