package fields

import "errors"

type PhoneTag struct{}
type Phone = Field[PhoneTag]

var ErrInvalidPhone = errors.New("invalid phone")

// Constructor
func NewPhone(value string) (phone Phone, err error) {
	return NewField[PhoneTag](value, phone.Tag.Validate, ErrInvalidPhone)
}

func (p PhoneTag) Validate(value string) bool {
	return true
}
