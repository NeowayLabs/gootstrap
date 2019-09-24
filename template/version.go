package template

const Version = `package version

import "fmt"

// This string will be overwritten during the build process.
var (
	GitVersion = ""
)

// Version returns a newline-terminated string describing the current
// version of the build.
func Version() string {
	if GitVersion == "" {
		return "devel\n"
	}
	return fmt.Sprintf("Version: %s\n", GitVersion)
}
`
