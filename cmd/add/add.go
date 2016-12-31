package add

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/urfave/cli"

	"github.com/joshwget/strato/state"
	"github.com/joshwget/strato/utils"
	"github.com/joshwget/strato/version"
)

func Action(c *cli.Context) error {
	source := c.GlobalString("source")
	dir := c.String("dir")

	var installs sync.WaitGroup
	for _, image := range c.Args() {
		installs.Add(1)
		go func(image string) {
			defer installs.Done()
			if err := add(dir, source+image+".tar.gz"); err != nil {
				panic(err)
			}
		}(image)
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
		if err = utils.ExtractTar(packageReader, dir, nil, nil); err != nil {
			return err
		}
		packageReader.Close()

		if err = state.AddToPackageList(image, dir); err != nil {
			return err
		}
	}

	return nil
}
