package errors

import "github.com/pkg/errors"

var (
	ErrNotPresentInCache = errors.New("Rate is not present in cache")
	ErrIsNotFloat64      = errors.New("Can't convert to float64")
)
