package index

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"strings"

	"github.com/rancher/strato/src/config"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var Command = cli.Command{
	Name:            "index",
	Usage:           "Generate index.json file from XXXXX",
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

	packageMap := map[string]config.Package{}

	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path.Join(inDir, file.Name(), config.Filename))
		if err != nil {
			return err
		}

		var pkg config.Package
		if err := yaml.Unmarshal(b, &pkg); err != nil {
			return err
		}

		packageName := file.Name()
		if strings.Contains(packageName, ".") {
			packageName = strings.SplitN(packageName, ".", 2)[1]
		}

		packageMap[packageName] = config.Package{
			Dependencies: pkg.Dependencies,
		}

		for subpackageName, subpackage := range pkg.Subpackages {
			packageMap[subpackageName] = config.Package{
				Dependencies: subpackage.Dependencies,
			}
		}
	}

	b, err := json.Marshal(packageMap)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(outDir, config.IndexName), b, 0644)
}
