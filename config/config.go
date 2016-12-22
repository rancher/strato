package config

import (
	"regexp"
)

const (
	Filename = "strato.yml"
)

type Package struct {
	License      string
	Version      string
	Description  string
	Dependencies []string
	Exclude      []string
	Subpackages  map[string][]string
	Precmd       string
	Postcmd      string
}

func GenerateWhiteAndBlackLists(pkg *Package, subpackage string) ([]*regexp.Regexp, []*regexp.Regexp, error) {
	var whitelist []*regexp.Regexp
	var blacklist []*regexp.Regexp
	if whitelistItems, ok := pkg.Subpackages[subpackage]; ok {
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
		for _, whitelistItems := range pkg.Subpackages {
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
