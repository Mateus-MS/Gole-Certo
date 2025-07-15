package fields

import "errors"

type CNPJTag struct{}
type CNPJ = Field[CNPJTag]

var ErrInvalidCNPJ = errors.New("invalid cnpj")

// Constructor
func NewCNPJ(value string) (cnpj CNPJ, err error) {
	return NewField[CNPJTag](value, cnpj.Tag.Validate, ErrInvalidCNPJ)
}

func (c CNPJTag) Validate(value string) bool {
	return true
}
