package fields

import (
	"errors"
)

type CNPJ struct {
	value string
}

var ErrorInvalidCNPJ = errors.New("invalid cnpj")

func NewCNPJ(value string) (cnpj CNPJ, err error) {
	cnpj.set(value)

	if !cnpj.validate() {
		return cnpj, ErrorInvalidCNPJ
	}

	return cnpj, nil
}

func (c *CNPJ) Get() string {
	return c.value
}

// Since the value can't change it will be validate only on this package
func (c *CNPJ) validate() bool {
	return true
}

// The set method is only available on this package for immutability
func (c *CNPJ) set(value string) {
	c.value = value
}
