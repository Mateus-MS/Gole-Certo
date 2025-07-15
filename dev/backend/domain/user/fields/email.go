package fields

import (
	"errors"
	"regexp"
)

type EmailTag struct{}
type Email = Field[EmailTag]

var ErrInvalidEmail = errors.New("invalid email")

// Constructor
func NewEmail(value string) (email Email, err error) {
	return NewField[EmailTag](value, email.Tag.Validate, ErrInvalidEmail)
}

func (e EmailTag) Validate(value string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(value)
}
