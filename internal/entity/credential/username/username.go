package username

import (
	"errors"
	"fmt"
)

type Username struct {
	value string
}

func New(value string, rules ...ValidationRule) (Username, error) {
	username := Username{
		value: value,
	}

	for _, rule := range commonRules {
		if err := rule(username); err != nil {
			return Username{}, err
		}
	}

	for _, rule := range rules {
		if err := rule(username); err != nil {
			return Username{}, err
		}
	}

	return username, nil
}

func (e Username) String() string {
	return e.value
}

type ValidationRule func(username Username) error

var commonRules = []ValidationRule{
	notEmpty(),
}

func notEmpty() ValidationRule {
	return func(username Username) error {
		if len(username.value) == 0 {
			return errors.New("username must not be empty")
		}
		return nil
	}
}

func MinLength(min int) ValidationRule {
	return func(username Username) error {
		if len([]rune(username.value)) < min {
			return fmt.Errorf("username must be at least %d characters", min)
		}
		return nil
	}
}

func MaxLength(max int) ValidationRule {
	return func(username Username) error {
		if len([]rune(username.value)) > max {
			return fmt.Errorf("username must be at most %d characters", max)
		}
		return nil
	}
}
