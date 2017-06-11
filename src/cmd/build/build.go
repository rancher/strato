package build

import (
	"io"
	"log"
	"os/exec"
	"path"

	"os"

	"strings"

	"path/filepath"

	"fmt"

	"github.com/rancher/strato/src/cmd/extract"
	"github.com/rancher/strato/src/config"
	"github.com/rancher/strato/src/utils"
	"github.com/urfave/cli"
)

// Command definition
var Command = cli.Command{
	Name:     "build",
	Usage:    "Build image from source and extract the last layer from the built image",
	HideHelp: true,
	Action:   buildAction,
	Flags:    []cli.Flag{},
}

// Build the package from Dockerfile, strato.yml and prebuild.sh
func buildAction(c *cli.Context) error {
	inDir := c.Args().Get(0)
	if _, err := os.Stat(filepath.Join(inDir, "Dockerfile")); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(filepath.Join(inDir, config.Filename)); os.IsNotExist(err) {
		return err
	}
	outDir, err := utils.GetOutDir(c.Args().Get(1))
	if err != nil {
		return err
	}

	packageName := path.Base(inDir)
	if strings.Contains(packageName, ".") {
		packageName = strings.SplitN(packageName, ".", 2)[1]
	}

	// TODO: don't overwrite an existing buildlog (it might mean we're just using the cache)
	f, err := os.OpenFile(filepath.Join(outDir, packageName+".buildlog"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	mwriter := io.MultiWriter(f, os.Stdout)

	// TODO: once there is a version number for the :tag, try pulling the image and only build when needed.
	// TODO: also need to add ARCH!
	cmd := exec.Command("docker", "build", "-t", "build/"+packageName, inDir)
	fmt.Printf("Running: %v\n", cmd.Args)
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Extracting into packages")
	if err := extract.Extract(inDir, outDir); err != nil {
		return err
	}

	return nil
}
