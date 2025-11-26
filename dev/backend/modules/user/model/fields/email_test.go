package fields_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

// Basic tests
func TestValidEmail(t *testing.T) {
	_, err := fields.NewEmail("testemail@gmail.com")
	assert.NoError(t, err, "Expected no errors")
}

func TestInvalidEmail(t *testing.T) {
	_, err := fields.NewEmail("not-an-email-my-brother")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail)
}

func TestEmptyEmail(t *testing.T) {
	_, err := fields.NewEmail("")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail)
}

// Some edge cases
func TestSubdomain(t *testing.T) {
	_, err := fields.NewEmail("testemail@gmail.co.uk")
	assert.NoError(t, err, "Expected no errors with subdomains")
}

func TestWithPlus(t *testing.T) {
	_, err := fields.NewEmail("testemail+test@gmail.co.uk")
	assert.NoError(t, err, "Expected no error with subdomains")
}

func TestWithDashes(t *testing.T) {
	_, err := fields.NewEmail("test-email@gmail.com")
	assert.NoError(t, err, "Expected no error with dashes")
}

// Malformed emails
func TestAtSymbol(t *testing.T) {
	_, err := fields.NewEmail("usermail.com")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email error")

	_, err = fields.NewEmail("user@@example.com")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email error")
}

func TestWithInvalidCharacters(t *testing.T) {
	_, err := fields.NewEmail("user^name@example.com")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email error")

	_, err = fields.NewEmail("user;name@example.com")
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email error")
}

func TestWhitespace(t *testing.T) {
	emailRaw := "user@example.com"
	eml, err := fields.NewEmail("     " + emailRaw)
	assert.NoError(t, err, "No error expected")
	assert.EqualValues(t, emailRaw, eml.Get(), "Expected to the withespaces not be in the final email")
}
