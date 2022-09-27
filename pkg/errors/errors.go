package errors

import "github.com/pkg/errors"

var (
	ErrNotPresentInCache = errors.New("rate is not present in cache")
	ErrIsNotFloat64      = errors.New("can't convert to float64")
	ErrAlreadySubscribed = errors.New("email address is already subscribed")
)
