package fields

import "errors"

type CPFTag struct{}
type CPF = Field[CPFTag]

var ErrInvalidCPF = errors.New("invalid cpf")

// Constructor
func NewCPF(value string) (cpf CPF, err error) {
	return NewField[CPFTag](value, cpf.Tag.Validate, ErrInvalidCPF)
}

func (c CPFTag) Validate(value string) bool {
	return true
}
