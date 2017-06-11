package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/strato/src/cmd/add"
	"github.com/rancher/strato/src/cmd/build"
	"github.com/rancher/strato/src/cmd/extract"
	"github.com/rancher/strato/src/cmd/index"
	"github.com/rancher/strato/src/cmd/inspect"
	"github.com/rancher/strato/src/cmd/xf"
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
		add.Command,
		build.Command,
		extract.Command,
		index.Command,
		inspect.Command,
		xf.Command,
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
