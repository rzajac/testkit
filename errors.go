package testkit

import (
	"errors"
)

// ErrTestError is general error used in tests,it is returned by some helpers.
var ErrTestError = errors.New("testkit test error")
