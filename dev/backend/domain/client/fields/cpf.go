package fields

import "errors"

type CPF struct {
	value string
}

var ErrorInvalidCPF = errors.New("invalid cpf")

func NewCPF(value string) (cpf CPF, err error) {
	cpf.set(value)

	if !cpf.validate() {
		return cpf, ErrorInvalidCPF
	}

	return cpf, nil
}

func (c *CPF) Get() string {
	return c.value
}

// Since the value can't change it will be validate only on this package
func (c *CPF) validate() bool {
	return true
}

// The set method is only available on this package for immutability
func (c *CPF) set(value string) {
	c.value = value
}
