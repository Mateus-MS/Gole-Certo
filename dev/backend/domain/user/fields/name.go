package fields

import "errors"

type NameTag struct{}
type Name = Field[NameTag]

var ErrInvalidName = errors.New("invalid name")

/*
I know it's kinda funny having a custom type for `name` since the validation is just a simple length check
but because of the validation flow i choose to stick with
i found that was easier to read and to maintain if it has the custom type like the ones that really makes sense to have a custom type like `email` or `phone`.
*/

// Constructor
func NewName(value string) (name Name, err error) {
	return NewField[NameTag](value, name.Tag.Validate, ErrInvalidName)
}

func (c NameTag) Validate(value string) bool {
	return len(value) >= 10
}

// Alias to other names
// I don't know if i will keep like this because this does not have type safety.

type LegalName = Name
type FantasyName = Name

func NewLegalName(value string) (name LegalName, err error) {
	return NewField[NameTag](value, name.Tag.Validate, ErrInvalidName)
}
func NewFantasyName(value string) (name LegalName, err error) {
	return NewField[NameTag](value, name.Tag.Validate, ErrInvalidName)
}
