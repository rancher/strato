package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetOutDir returns the output dir for this arch (makes it if it needs to)
func GetOutDir(dir string) (string, error) {
	outDir := filepath.Join(dir, runtime.GOARCH)
	err := os.MkdirAll(outDir, 0755)

	return outDir, err
}
