package fields

import (
	"errors"
)

type Email struct {
	value string
}

var ErrorInvalidEmail = errors.New("invalid email")

func NewEmail(value string) (email Email, err error) {
	email.set(value)

	if !email.validate() {
		return email, ErrorInvalidEmail
	}

	return email, nil
}

func (e *Email) Get() string {
	return e.value
}

// Since the value can't change it will be validate only on this package
func (e *Email) validate() bool {
	return true
}

// The set method is only available on this package for immutability
func (e *Email) set(value string) {
	e.value = value
}
