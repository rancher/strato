package buildorder

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/rancher/strato/src/config"
	yaml "gopkg.in/yaml.v2"
)

func Get(packagesPath string) ([]string, error) {
	files, err := ioutil.ReadDir(packagesPath)
	if err != nil {
		return nil, err
	}

	dependencies := map[string]map[string]bool{}
	subpackages := map[string]map[string]bool{}
	for _, file := range files {
		dockerfilePath := path.Join(packagesPath, file.Name(), "Dockerfile")
		contents, err := ioutil.ReadFile(dockerfilePath)
		if err != nil {
			return nil, err
		}

		stratoPath := path.Join(packagesPath, file.Name(), config.Filename)
		stratoContents, err := ioutil.ReadFile(stratoPath)
		if err != nil {
			return nil, err
		}

		packageName := file.Name()

		dependencies[packageName] = parsePackageDependencies(string(contents))
		subpackages[packageName], err = parseSubpackages(stratoContents)
		if err != nil {
			return nil, err
		}
	}

	var order []string
	covered := map[string]bool{}
	for len(dependencies) > 0 {
		for pkg, pkgDependencies := range dependencies {
			dependenciesMet := true
			for pkgDependency := range pkgDependencies {
				if _, ok := covered[pkgDependency]; !ok {
					dependenciesMet = false
				}
			}
			if dependenciesMet {
				covered[pkg] = true
				for subpackage := range subpackages[pkg] {
					covered[subpackage] = true
				}
				order = append(order, pkg)
				delete(dependencies, pkg)
			}
		}
	}

	return order, nil
}

func parsePackageDependencies(contents string) map[string]bool {
	packages := map[string]bool{}

	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		if len(split) < 4 {
			continue
		}
		if split[0] != "RUN" || split[1] != "strato" || split[2] != "add" {
			continue
		}
		for _, pkg := range split[3:] {
			packages[pkg] = true
		}
	}

	return packages
}

func parseSubpackages(contents []byte) (map[string]bool, error) {
	var pkg config.Package
	if err := yaml.Unmarshal(contents, &pkg); err != nil {
		return nil, err
	}

	subpackages := map[string]bool{}
	for subpackage := range pkg.Subpackages {
		subpackages[subpackage] = true
	}

	return subpackages, nil
}
