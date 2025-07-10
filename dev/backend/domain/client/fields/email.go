package fields

import "errors"

type EmailTag struct{}
type Email = Field[EmailTag]

var ErrInvalidEmail = errors.New("invalid email")

// Constructor
func NewEmail(value string) (email Email, err error) {
	return NewField[EmailTag](value, email.Tag.Validate, ErrInvalidEmail)
}

func (e EmailTag) Validate(value string) bool {
	return true
}
