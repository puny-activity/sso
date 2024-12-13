package errs

import (
	"errors"
	"fmt"
)

func Wrap(msg string, err error) error {
	if err == nil {
		return errors.New(msg)
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func Detail(err error, detail string) error {
	if err == nil {
		return errors.New(detail)
	}
	return fmt.Errorf("%w: %s", err, detail)
}
