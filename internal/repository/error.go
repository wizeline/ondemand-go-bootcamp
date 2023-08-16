package repository

import (
	"fmt"
)

const (
	ErrFilterEmpty   Err = "filter empty"
	ErrInvalidFilter Err = "invalid filter"
)

type Err string

func (e Err) Error() string {
	return fmt.Sprintf("repository: %v", e.String())
}

func (e Err) String() string {
	return string(e)
}

func newCsvErr(err error) Err {
	return Err(fmt.Sprintf("csv: %s", err))
}
