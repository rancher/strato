package index

import (
	"io/ioutil"
	"path"

	"github.com/joshwget/strato/config"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func Action(c *cli.Context) error {
	inDir := c.Args()[0]
	outDir := c.Args()[1]

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

		pkg.Exclude = nil
		// TODO: split out subpackages
		pkg.Subpackages = nil

		packageMap[file.Name()] = pkg
	}

	b, err := yaml.Marshal(packageMap)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(outDir, "index.yml"), b, 0644)
}
