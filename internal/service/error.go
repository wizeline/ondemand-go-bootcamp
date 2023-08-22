package service

import (
	"errors"
	"fmt"
)

var (
	ErrFilterValueEmpty     = errors.New("filter empty")
	ErrFilterNotSupported   = errors.New("filter not supported")
	ErrFilterNotImplemented = errors.New("filter not implemented")
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
