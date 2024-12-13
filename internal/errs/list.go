package errs

import (
	"errors"
)

type appError struct {
	code int
	msg  string
}

func (e *appError) Error() string {
	return e.msg
}

var (
	unexpected = appError{code: 0, msg: "unexpected error"}
	Validation = appError{code: 1, msg: "validation error"}
)

var errorList = []appError{
	Validation,
}

func ExtractCode(err error) int {
	if err == nil {
		return 0
	}

	var appErr *appError
	if errors.As(err, &appErr) {
		return appErr.code
	}
	return unexpected.code
}

func ExtractText(err error) string {
	if err == nil {
		return ""
	}

	var appErr *appError
	if errors.As(err, &appErr) {
		return appErr.msg
	}
	return unexpected.msg
}

func ExtractTextAsError(err error) error {
	if err == nil {
		return nil
	}

	var appErr *appError
	if errors.As(err, &appErr) {
		return errors.New(appErr.msg)
	}
	return errors.New(unexpected.msg)
}
