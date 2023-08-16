package utils

import (
	"fmt"
)

const (
	ErrDataFilePathEmpty   Err = "data file path empty"
	ErrDataFileIsDir       Err = "data file is a directory"
	ErrDataFileInvalidMode Err = "invalid data file mode"
)

var _ error = Err("")

type Err string

func (e Err) Error() string {
	return fmt.Sprintf("utils: %v", e.String())
}

func (e Err) String() string {
	return string(e)
}
