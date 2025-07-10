package fields

import "errors"

type PhoneTag struct{}
type Phone = Field[PhoneTag]

var ErrInvalidPhone = errors.New("invalid phone")

// Constructor
func NewPhone(value string) (phone Phone, err error) {
	return NewField[PhoneTag](value, validatePhone, ErrInvalidPhone)
}

func validatePhone(value string) bool {
	return true
}
