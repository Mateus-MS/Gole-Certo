package fields_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user/fields"
	"github.com/stretchr/testify/assert"
)

func TestValidName(t *testing.T) {
	_, err := fields.NewName("jorge da silva teles")
	assert.NoError(t, err, "Expected no errors")
}

func TestTooShortValidName(t *testing.T) {
	_, err := fields.NewName("jorge")
	assert.ErrorIs(t, err, fields.ErrInvalidName, "Excpeted error invalid name")
}

func TestWithInvalidCharactersName(t *testing.T) {
	name, err := fields.NewName("jorge ; ++ francisco")
	assert.NoError(t, err, "Expected no errors")
	assert.EqualValues(t, "jorge francisco", name.Get(), "Expected: 'jorge francisco' got: "+name.Get())
}
