package config

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
}
