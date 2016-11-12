package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/joshwget/lay/cmd/add"
	"github.com/joshwget/lay/cmd/inspect"
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
			Name:  "registry",
			Value: "https://registry-1.docker.io/",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
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
				cli.StringFlag{
					Name: "skip",
				},
			},
		},
		{
			Name:            "inspect",
			HideHelp:        true,
			SkipFlagParsing: true,
			Action:          inspect.Action,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
