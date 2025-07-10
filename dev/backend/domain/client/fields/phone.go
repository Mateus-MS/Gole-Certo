package fields

import (
	"errors"
)

type Phone struct {
	value string
}

var ErrorInvalidPhone = errors.New("invalid phone")

func NewPhone(value string) (phone Phone, err error) {
	phone.set(value)

	if !phone.validate() {
		return phone, ErrorInvalidPhone
	}

	return phone, nil
}

func (c *Phone) Get() string {
	return c.value
}

// Since the value can't change it will be validate only on this package
func (c *Phone) validate() bool {
	return true
}

// The set method is only available on this package for immutability
func (c *Phone) set(value string) {
	c.value = value
}
