package publickey

import "errors"

type Key struct {
	value string
}

func New(value string) (Key, error) {
	username := Key{
		value: value,
	}

	for _, rule := range commonRules {
		if err := rule(username); err != nil {
			return Key{}, err
		}
	}

	return username, nil
}

func (e *Key) String() string {
	return e.value
}

type ValidationRule func(key Key) error

var commonRules = []ValidationRule{
	notEmpty(),
}

func notEmpty() ValidationRule {
	return func(key Key) error {
		if len(key.value) == 0 {
			return errors.New("key must not be empty")
		}
		return nil
	}
}
