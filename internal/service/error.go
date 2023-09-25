package service

import (
	"errors"
	"fmt"
)

var (
	ErrFltrTypeEmpty  = errors.New("filter type empty")
	ErrFltrValueEmpty = errors.New("filter value empty")
	ErrFltrInvalid    = errors.New("invalid filter")
)

type FilterErr struct {
	Err error
}

func (e FilterErr) Error() string {
	return fmt.Sprintf("service filter: %s", e.Err)
}

func (e FilterErr) Unwrap() error {
	return e.Err
}
