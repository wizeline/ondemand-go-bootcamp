package repository

import (
	"errors"
	"fmt"
)

var (
	ErrFilePathEmpty   = errors.New("data file path empty")
	ErrFilePathIsDir   = errors.New("data file path is a directory")
	ErrFileModeInvalid = errors.New("invalid data file mode")
)

type CsvErr struct {
	Err error
}

func (e CsvErr) Error() string {
	return fmt.Sprintf("csv: %s", e.Err)
}

func (e CsvErr) Unwrap() error {
	return e.Err
}
