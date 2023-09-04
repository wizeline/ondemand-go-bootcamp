package configuration

import (
	"fmt"
	"strconv"
	"strings"
)

// SemanticVersion is the versioning pattern used to handle the application's versions.
// Ref: https://semver.org/
type SemanticVersion string

var _ fmt.Stringer = SemanticVersion("")

func (s SemanticVersion) String() string {
	return string(s)
}

// MajorVersion returns the number of the major version
func (s SemanticVersion) MajorVersion() int {
	strVer := s.String()
	strVer = strings.Trim(strVer, "v")
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
