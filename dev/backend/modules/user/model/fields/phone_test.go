package fields_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

func TestValidPhone(t *testing.T) {
	phone, err := fields.NewPhone("+351911911911")
	assert.NoError(t, err, "Expected no errors")
	assert.EqualValues(t, "911911911", phone.Get(), "Expected to remove the +351")

	_, err = fields.NewPhone("911911911")
	assert.NoError(t, err, "Expected no errors")

	_, err = fields.NewPhone("00351911911911")
	assert.NoError(t, err, "Expected no errors")
}

func TestInvalidPhones(t *testing.T) {
	_, err := fields.NewPhone("351911911911")
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected no errors")

	_, err = fields.NewPhone("981911911")
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected no errors")

	_, err = fields.NewPhone("9119119112")
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected no errors")
}
