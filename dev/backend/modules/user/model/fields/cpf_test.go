package fields_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

// Basic tests
func TestValidCPF(t *testing.T) {
	_, err := fields.NewCPF("033.355.662-38")
	assert.NoError(t, err, "Expected no error")
}

func TestInvalidCPF(t *testing.T) {
	_, err := fields.NewCPF("033.355.662-39")
	assert.ErrorIs(t, err, fields.ErrInvalidCPF)

	_, err = fields.NewCPF("abc.355.662-39")
	assert.ErrorIs(t, err, fields.ErrInvalidCPF)

	_, err = fields.NewCPF("abc.355;662-39")
	assert.ErrorIs(t, err, fields.ErrInvalidCPF)
}

func TestEmptyCPF(t *testing.T) {
	_, err := fields.NewCPF("")
	assert.ErrorIs(t, err, fields.ErrInvalidCPF)
}

func TestOnlyNumberCPF(t *testing.T) {
	_, err := fields.NewCPF("03335566238")
	assert.NoError(t, err, "Expected no errors")
}

func TestWithWhitespacesCPF(t *testing.T) {
	_, err := fields.NewCPF(" 03 3 35 566 238")
	assert.NoError(t, err, "Expected no errors")

	_, err = fields.NewCPF(" 03 3. 35 5.66 2-38")
	assert.NoError(t, err, "Expected no errors")
}
