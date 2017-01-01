package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/joshwget/strato/config"
	"gopkg.in/yaml.v2"
)

func main() {
	inDir := os.Args[1]
	outDir := os.Args[2]

	packageMap := map[string]config.Package{}

	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path.Join(inDir, file.Name(), config.Filename))
		if err != nil {
			panic(err)
		}

		var pkg config.Package
		if err := yaml.Unmarshal(b, &pkg); err != nil {
			panic(err)
		}

		pkg.Exclude = nil
		// TODO: split out subpackages
		pkg.Subpackages = nil

		packageMap[file.Name()] = pkg
	}

	b, err := yaml.Marshal(packageMap)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(path.Join(outDir, "index.yml"), b, 0644); err != nil {
		panic(err)
	}
}
