package fields

import (
	"errors"
	"regexp"
	"strings"
)

type PhoneTag struct{}
type Phone = Field[PhoneTag]

var ErrInvalidPhone = errors.New("invalid phone")

// Constructor
func NewPhone(value string) (phone Phone, err error) {
	value = phone.Tag.sanitize(value)

	return NewField[PhoneTag](value, phone.Tag.Validate, ErrInvalidPhone)
}

func (p PhoneTag) Validate(value string) bool {
	var re = regexp.MustCompile(`^(\+351|00351)?(9[1236]\d{7}|2\d{8})$`)
	return re.MatchString(value)
}

func (e PhoneTag) sanitize(value string) (clean string) {
	// Remove all whitespace characters (space, tabs, newlines, etc.)
	noSpaces := strings.ReplaceAll(value, " ", "")

	// Remove "+351" prefix if it exists at the start
	noSpaces = strings.TrimPrefix(noSpaces, "+351")

	return noSpaces
}
