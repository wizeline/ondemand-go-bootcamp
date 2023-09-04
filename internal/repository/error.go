package repository

import (
	"errors"
	"fmt"
)

var (
	ErrFilePathEmpty   = errors.New("validData file path empty")
	ErrFilePathIsDir   = errors.New("validData file path is a directory")
	ErrFileModeInvalid = errors.New("invalid validData file mode")

	ErrCSVRecordEmpty    = errors.New("csv record empty")
	ErrTimeStringEmpty   = errors.New("unix epoch time string empty")
	ErrInvalidTimeString = errors.New("invalid unix epoch time string")
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
