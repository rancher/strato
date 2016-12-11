package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/joshwget/strato/config"
	"gopkg.in/yaml.v2"
)

func ExtractTar(reader io.Reader, target string, skip *regexp.Regexp) error {
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
		if skip != nil && skip.MatchString(filename) {
			continue
		}
		filename = path.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			log.Debugf("Dir: %s", filename)
			if err = os.MkdirAll(filename, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			log.Debugf("File: %s", filename)
			if _, err := os.Stat(filename); err == nil {
				if err := os.Remove(filename); err != nil {
					return err
				}
			}
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
			if _, err := os.Stat(filename); err == nil {
				if err := os.Remove(filename); err != nil {
					return err
				}
			}
			if err := os.Link(header.Linkname, filename); err != nil {
				return err
			}
		case tar.TypeSymlink:
			log.Debugf("Soft link: %s", filename)
			if _, err := os.Stat(filename); err == nil {
				if err := os.Remove(filename); err != nil {
					return err
				}
			}
			if err := os.Symlink(header.Linkname, filename); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Failed to untar %s (%c)", filename, header.Typeflag)
		}
	}

	return nil
}

func FindPackage(reader io.Reader) (*config.Package, error) {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		filename := header.Name
		if strings.Contains(filename, config.Filename) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(tarReader)
			var pkg config.Package
			if err := yaml.Unmarshal(buf.Bytes(), &pkg); err != nil {
				return nil, err
			}
			return &pkg, nil
		}
		if filename > config.Filename {
			return nil, nil
		}
	}
	return nil, nil
}
