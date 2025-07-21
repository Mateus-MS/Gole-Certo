package fields

import (
	"errors"
	"regexp"
	"strings"
)

type EmailTag struct{}
type Email = Field[EmailTag]

var ErrInvalidEmail = errors.New("invalid email")

// Constructor
func NewEmail(value string) (email Email, err error) {
	value = email.Tag.sanitize(value)

	return NewField[EmailTag](value, email.Tag.Validate, ErrInvalidEmail)
}

func (e EmailTag) Validate(value string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(value)
}

func (e EmailTag) sanitize(value string) (clean string) {
	return strings.TrimSpace(value)
}
