package state

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

var Size float64

var (
	stateDir         = "/var/lib/strato"
	packagesListFile = path.Join(stateDir, "packages")
)

var (
	packageListMutex = sync.Mutex{}
)

func AddToPackageList(pkg, dir string) error {
	return withPackageListLock(dir, func(packages []string) ([]string, error) {
		return append(packages, pkg), nil
	})
}

func InPackageList(pkg, dir string) (bool, error) {
	inPackageList := false

	if err := withPackageListLock(dir, func(packages []string) ([]string, error) {
		for _, existingPackage := range packages {
			if pkg == existingPackage {
				inPackageList = true
				break
			}
		}
		return nil, nil
	}); err != nil {
		return false, err
	}

	return inPackageList, nil
}

func withPackageListLock(dir string, f func([]string) ([]string, error)) error {
	packagesListFile := path.Join(dir, packagesListFile)

	packageListMutex.Lock()

	bytes, err := ioutil.ReadFile(packagesListFile)
	if err != nil {
		return err
	}
	packages := strings.Split(string(bytes), "\n")

	packages, err = f(packages)
	if err != nil {
		return err
	}

	if packages != nil {
		if err := ioutil.WriteFile(packagesListFile, []byte(strings.Join(packages, "\n")), os.ModePerm); err != nil {
			return err
		}
	}

	packageListMutex.Unlock()

	return nil
}
