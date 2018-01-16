package config

import (
	"regexp"
)

const (
	Filename  = "strato.yml"
	IndexName = "index.json"
)

type Package struct {
	License      string   `yaml:"license,omitempty" json:"license,omitempty"`
	Version      string   `yaml:"version,omitempty" json:"license,omitempty"`
	Description  string   `yaml:"description,omitempty" json:"description,omitempty"`
	Dependencies []string `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`
	// TODO: implement include this for Package
	Include       []string              `yaml:"include,omitempty" json:"include,omitempty"`
	Exclude       []string              `yaml:"exclude,omitempty" json:"exclude,omitempty"`
	Subpackages   map[string]Subpackage `yaml:"subpackages,omitempty" json:"subpackages,omitempty"`
	Precmd        string                `yaml:"precmd,omitempty" json:"precmd,omitempty"`
	Postcmd       string                `yaml:"postcmd,omitempty" json:"postcmd,omitempty"`
	ExtractFolder string                `yaml:"extract_folder,omitempty" json:"extract_folder,omitempty"`
}

type Subpackage struct {
	License      string   `yaml:"license,omitempty" json:"license,omitempty"`
	Version      string   `yaml:"version,omitempty" json:"version,omitempty"`
	Description  string   `yaml:"description,omitempty" json:"description,omitempty"`
	Dependencies []string `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`
	Include      []string `yaml:"include,omitempty" json:"include,omitempty"`
	Exclude      []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
	Precmd       string   `yaml:"precmd,omitempty" json:"precmd,omitempty"`
	Postcmd      string   `yaml:"postcmd,omitempty" json:"postcmd,omitempty"`
}

func GenerateWhiteAndBlackLists(pkg *Package, subpackage string) ([]*regexp.Regexp, []*regexp.Regexp, error) {
	var whitelist []*regexp.Regexp
	var blacklist []*regexp.Regexp
	if subpackage, ok := pkg.Subpackages[subpackage]; ok {
		whitelistItems := subpackage.Include
		// Only install whitelisted for subpackages
		for _, whitelistItem := range whitelistItems {
			whitelistRegex, err := regexp.Compile(whitelistItem)
			if err != nil {
				return nil, nil, err
			}
			whitelist = append(whitelist, whitelistRegex)
		}
	} else {
		// Blacklist the union of all subpackage whitelists for regular packages
		var union []*regexp.Regexp
		for _, subpackage := range pkg.Subpackages {
			whitelistItems := subpackage.Include
			for _, whitelistItem := range whitelistItems {
				whitelistRegex, err := regexp.Compile(whitelistItem)
				if err != nil {
					return nil, nil, err
				}
				union = append(union, whitelistRegex)
			}
		}
		blacklist = union
	}
	for _, exclude := range pkg.Exclude {
		excludeRegex, err := regexp.Compile(exclude)
		if err != nil {
			return nil, nil, err
		}
		blacklist = append(blacklist, excludeRegex)
	}
	return whitelist, blacklist, nil
}
