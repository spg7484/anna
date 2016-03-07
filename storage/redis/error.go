package redisstorage

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var queryExecutionFailedError = errgo.New("query execution failed")

// IsQueryExecutionFailed asserts queryExecutionFailedError.
func IsQueryExecutionFailed(err error) bool {
	return errgo.Cause(err) == queryExecutionFailedError
}