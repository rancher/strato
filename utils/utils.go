package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func ExtractTar(reader io.Reader, target string) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		filename := header.Name
		filename = path.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			log.Debugf("Dir: %s", filename)
			if err = os.MkdirAll(filename, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			log.Debugf("File: %s", filename)
			writer, err := os.Create(filename)
			if err != nil {
				return err
			}
			io.Copy(writer, tarReader)
			if err = os.Chmod(filename, os.FileMode(header.Mode)); err != nil {
				return err
			}
			writer.Close()
		case tar.TypeLink:
			log.Debugf("Hard link: %s", filename)
			if err := os.Link(header.Linkname, filename); err != nil {
				return err
			}
		case tar.TypeSymlink:
			log.Debugf("Soft link: %s", filename)
			if err := os.Symlink(header.Linkname, filename); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Failed to untar %s (%c)", filename, header.Typeflag)
		}
	}

	return nil
}

func IsPackage(reader io.Reader) (bool, error) {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return false, err
	}
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}
		filename := header.Name
		if strings.Contains(filename, "_magic") {
			return true, nil
		}
		if filename > "_magic" {
			return false, nil
		}
	}
	return false, nil
}
