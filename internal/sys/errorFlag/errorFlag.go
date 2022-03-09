package errorFlag

import (
	"errors"
)

const (
	InvalidData   = iota + 1 // Invalid input data provided
	NotFound                 // Target not found
	AlreadyExists            // Entity already exists
	Internal                 // Internal error or inconsistency
)

type ErrorFlag int

type flagged struct {
	error
	flag ErrorFlag
}

type Flagged interface {
	Unwrap() error
	Flag() ErrorFlag
}

// New wraps err with an error that will return true from HasFlag(err, flag).
func New(err error, flag ErrorFlag) error {
	if err == nil {
		return nil
	}
	return flagged{error: err, flag: flag}
}

// HasFlag reports if err has been flagged with the given flag.
func HasFlag(err error, flag ErrorFlag) bool {
	for {
		if f, ok := err.(flagged); ok && f.flag == flag {
			return true
		}
		if err = errors.Unwrap(err); err == nil {
			return false
		}
	}
}

func (f flagged) Unwrap() error {
	return f.error
}

func (f flagged) Flag() ErrorFlag {
	return f.flag
}
