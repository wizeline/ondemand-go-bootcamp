package config

import (
	"fmt"
	"strconv"
	"strings"
)

const apiPrefixFmt = "/api/v%d"

// Application holds and retrieves the application's configuration properties.
type Application struct {
	version string
}

// Version returns the semantic version of the application. Ref: https://semver.org/
func (a Application) Version() string {
	return a.version
}

// MajorVersion returns the  major version's number of the application.
func (a Application) MajorVersion() int {
	strVer := strings.Trim(a.version, "v")
	if len(strVer) == 0 {
		return 0
	}

	const separatorPattern byte = '.'
	buf := strings.Builder{}
	for i := 0; i < len(strVer); i++ {
		if separatorPattern == strVer[i] {
			break
		}
		buf.WriteByte(strVer[i])
	}
	ver, _ := strconv.Atoi(buf.String())
	return ver
}

// BasePath returns the base path in the form "/api/v{MajorVersion}"
func (a Application) BasePath() string {
	return fmt.Sprintf(apiPrefixFmt, a.MajorVersion())
}
