package fields

import (
	"errors"
	"regexp"
)

type PhoneTag struct{}
type Phone = Field[PhoneTag]

var ErrInvalidPhone = errors.New("invalid phone")

// Constructor
func NewPhone(value string) (phone Phone, err error) {
	return NewField[PhoneTag](value, phone.Tag.Validate, ErrInvalidPhone)
}

func (p PhoneTag) Validate(value string) bool {
	var re = regexp.MustCompile(`^(\+351|00351)?(9[1236]\d{7}|2\d{8})$`)
	return re.MatchString(value)
}
