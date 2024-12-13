package password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Unhashed struct {
	value string
}

type Hashed struct {
	value string
}

func NewUnhashed(value string) (Unhashed, error) {
	return Unhashed{
		value: value,
	}, nil
}

func NewHashed(value string) (Hashed, error) {
	return Hashed{
		value: value,
	}, nil
}

func (e Unhashed) Hash() (Hashed, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(e.value), bcrypt.DefaultCost)
	if err != nil {
		return Hashed{}, fmt.Errorf("failed to hash password: %w", err)
	}
	return NewHashed(string(hashedBytes))
}

func IsMatch(unhashed Unhashed, hashed Hashed) bool {
	return bcrypt.CompareHashAndPassword([]byte(unhashed.value), []byte(hashed.value)) == nil
}
