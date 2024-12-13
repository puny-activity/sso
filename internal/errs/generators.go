package errs

import "fmt"

func GenInvalidParameterValue(name string, value string) error {
	return fmt.Errorf("%w: invalid value \"%v\" for parameter \"%s\"", &Validation, value, name)
}
