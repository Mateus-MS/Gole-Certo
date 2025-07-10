package fields

import "errors"

type EmailTag struct{}
type Email = Field[EmailTag]

var ErrInvalidEmail = errors.New("invalid email")

// Constructor
func NewEmail(value string) (email Email, err error) {
	return NewField[EmailTag](value, validateEmail, ErrInvalidEmail)
}

func validateEmail(value string) bool {
	return true
}
