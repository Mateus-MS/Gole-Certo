package fields_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

func TestValidCNPJ(t *testing.T) {
	cnpj, err := fields.NewCNPJ("12.345.678/0001-95")
	assert.NoError(t, err, "Expected no errors")
	assert.EqualValues(t, "12345678000195", cnpj.Get())

	_, err = fields.NewCNPJ("12.345.????6!!!78     /00;01-95")
	assert.NoError(t, err, "Expected no errors")
	assert.EqualValues(t, "12345678000195", cnpj.Get())

	cnpj, err = fields.NewCNPJ("12345678000195")
	assert.NoError(t, err, "Expected no errors")
	assert.EqualValues(t, "12345678000195", cnpj.Get())
}

func TestInvalidCNPJ(t *testing.T) {
	_, err := fields.NewCNPJ("not-an-CNPJ-my-brother")
	assert.ErrorIs(t, err, fields.ErrInvalidCNPJ)

	_, err = fields.NewCNPJ("12.345.678/0001-94")
	assert.ErrorIs(t, err, fields.ErrInvalidCNPJ)

	_, err = fields.NewCNPJ("12.345!!.?678; /0   001-94")
	assert.ErrorIs(t, err, fields.ErrInvalidCNPJ)
}
