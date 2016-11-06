package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/joshwget/lay/cmd/down"
	"github.com/joshwget/lay/cmd/inspect"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "registry",
			Value: "https://registry-1.docker.io/",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "down",
			HideHelp: true,
			Action:   down.Action,
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
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
