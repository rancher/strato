package config

const (
	Filename = "_package.yml"
)

type Package struct {
	License      string
	Version      string
	Description  string
	Dependencies []string
}
