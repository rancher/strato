package main

import (
	"os"

	"github.com/joshwget/strato/cmd/add"
	"github.com/joshwget/strato/cmd/inspect"
	"github.com/joshwget/strato/cmd/xf"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
		},
		cli.StringFlag{
			Name: "source",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "add",
			HideHelp: true,
			Action:   add.Action,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir",
					Value: "/",
				},
			},
		},
		{
			Name:            "inspect",
			HideHelp:        true,
			SkipFlagParsing: true,
			Action:          inspect.Action,
		},
		{
			Name:            "xf",
			HideHelp:        true,
			SkipFlagParsing: true,
			Action:          xf.Action,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
