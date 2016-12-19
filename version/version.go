package version

import "runtime"

var (
	Arch    string
	Suffix  string
	Tag     string
	Version string
)

func init() {
	if Arch == "" {
		Arch = runtime.GOARCH
	}
	if Suffix == "" && Arch != "amd64" {
		Suffix = "_" + Arch
	}
	if Version == "" {
		Version = "dev"
	}
	Tag = Version + Suffix
}
