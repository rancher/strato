package add

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/urfave/cli"

	"github.com/rancher/strato/src/config"
	"github.com/rancher/strato/src/state"
	"github.com/rancher/strato/src/utils"
	"github.com/rancher/strato/src/version"
)

const (
	// TODO: move to different package
	repositoriesFile = "/etc/strato/repositories"
)

var Command = cli.Command{
	Name:     "add",
	Usage:    "add/install a package",
	HideHelp: true,
	Action:   Action,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Value: "/",
		},
	},
}

func Action(c *cli.Context) error {
	dir := c.String("dir")
	source := c.GlobalString("source")

	if source == "" {
		repositoriesFileBytes, err := ioutil.ReadFile(repositoriesFile)
		if err != nil {
			log.Panic(err)
			return err
		}
		source = strings.Trim(string(repositoriesFileBytes), "\n")
	}
	source = source + "/" + runtime.GOARCH + "/"

	var indexBytes []byte
	var err error
	if path.IsAbs(source) {
		indexBytes, err = ioutil.ReadFile(path.Join(source, config.IndexName))
		if err != nil {
			log.Panic(err)
			return err
		}
	} else {
		u, err := url.Parse(source)
		u.Path = path.Join(u.Path, config.IndexName)
		resp, err := http.Get(u.String())
		if err != nil {
			log.Panic(err)
			return err
		}
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			err = fmt.Errorf("Error getting %s: %v", u.String(), resp.Status)
			log.Panic(err)
			return err
		}
		indexBytes, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Panic(err)
			return err
		}
	}

	index := map[string]config.Package{}
	if err := json.Unmarshal(indexBytes, &index); err != nil {
		log.Panic(err)
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
				log.Panic(err)
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

		fmt.Printf("Installing package %s\n", fmt.Sprintf("%s:%s", image, version.Tag))
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
