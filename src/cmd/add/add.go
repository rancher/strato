package add

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/urfave/cli"

	"github.com/joshwget/strato/src/config"
	"github.com/joshwget/strato/src/state"
	"github.com/joshwget/strato/src/utils"
	"github.com/joshwget/strato/src/version"
	"gopkg.in/yaml.v2"
)

const (
	// TODO: move to different package
	repositoriesFile = "/etc/strato/repositories"
)

func Action(c *cli.Context) error {
	dir := c.String("dir")
	source := c.GlobalString("source")

	if source == "" {
		repositoriesFileBytes, err := ioutil.ReadFile(repositoriesFile)
		if err != nil {
			return err
		}
		source = strings.Trim(string(repositoriesFileBytes), "\n")
	}

	var indexBytes []byte
	var err error
	if path.IsAbs(source) {
		indexBytes, err = ioutil.ReadFile(path.Join(source, "index.yml"))
		if err != nil {
			return err
		}
	} else {
		resp, err := http.Get(source + "index.yml")
		if err != nil {
			return err
		}
		indexBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	index := map[string]config.Package{}
	if err := yaml.Unmarshal(indexBytes, &index); err != nil {
		return err
	}

	packageMap := map[string]bool{}
	for _, pkg := range c.Args() {
		packageMap[pkg] = true
		// TODO: should be recursive
		for _, dependency := range index[pkg].Dependencies {
			packageMap[dependency] = true
		}
	}

	var installs sync.WaitGroup
	for pkg := range packageMap {
		installs.Add(1)
		go func(pkg string) {
			defer installs.Done()
			if err := add(dir, source+pkg+".tar.gz"); err != nil {
				panic(err)
			}
		}(pkg)
	}
	installs.Wait()

	fmt.Println(utils.Size/1000000.0, "mb")

	return nil
}

func add(dir string, packages ...string) error {
	for _, image := range packages {
		inPackageList, err := state.InPackageList(image, dir)
		if err != nil {
			return err
		}
		if inPackageList {
			continue
		}

		var packageReader io.ReadCloser
		if path.IsAbs(image) {
			packageReader, err = os.Open(image)
			if err != nil {
				return err
			}
		} else {
			resp, err := http.Get(image)
			if err != nil {
				return err
			}
			packageReader = resp.Body
		}

		fmt.Printf("Installing package %s", fmt.Sprintf("%s:%s", image, version.Tag))
		if err = utils.ExtractGzipTar(packageReader, dir, nil, nil); err != nil {
			return err
		}
		packageReader.Close()

		if err = state.AddToPackageList(image, dir); err != nil {
			return err
		}
	}

	return nil
}
