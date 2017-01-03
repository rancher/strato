package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/joshwget/strato/config"
)

// TODO: not exactly thread safe...
var Size float64

func ExtractTar(reader io.Reader, target string, whitelist, blacklist []*regexp.Regexp) error {
	return GzipTarForEach(reader, whitelist, blacklist, writeFile(target))
}

func writeFile(target string) func(io.Reader, *tar.Header) error {
	return func(tarReader io.Reader, header *tar.Header) error {
		filename := path.Join(target, header.Name)
		fmt.Println(filename)
		Size += float64(header.FileInfo().Size())

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filename, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
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
			if err = os.Chmod(filename, header.FileInfo().Mode()); err != nil {
				return err
			}
			writer.Close()
		case tar.TypeLink:
			if _, err := os.Stat(filename); err == nil {
				if err := os.Remove(filename); err != nil {
					return err
				}
			}
			if err := os.Link(header.Linkname, filename); err != nil {
				return err
			}
		case tar.TypeSymlink:
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
		return nil
	}
}

func GzipTarForEach(reader io.Reader, whitelist, blacklist []*regexp.Regexp, f func(io.Reader, *tar.Header) error) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	return TarForEach(gzipReader, whitelist, blacklist, f)
}

func TarForEach(reader io.Reader, whitelist, blacklist []*regexp.Regexp, f func(io.Reader, *tar.Header) error) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		filename := header.Name
		if filename == config.Filename {
			continue
		}
		if len(whitelist) > 0 {
			passes := false
			for _, r := range whitelist {
				if r.MatchString(filename) {
					passes = true
				}
			}
			if !passes {
				continue
			}
		}
		passes := true
		for _, r := range blacklist {
			if r.MatchString(filename) {
				passes = false
			}
		}
		if !passes {
			continue
		}
		// Temporarily ignored conditions
		if strings.HasSuffix(filename, ".a") {
			continue
		}
		if strings.HasPrefix(filename, "tmp/") {
			continue
		}
		if strings.HasPrefix(filename, "usr/src/") {
			continue
		}

		if err := f(tarReader, header); err != nil {
			return err
		}
	}

	return nil
}
