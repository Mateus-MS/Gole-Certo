package user_test

import (
	"testing"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

func TestNewIndividualValid(t *testing.T) {
	cpf := "12345678909"
	email := "valid-email@gmail.com"
	phones := []string{"911911911"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"Valid Name"}

	ind, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.NoError(t, err, "Expected no error on valid individual")
	assert.Equal(t, cpf, ind.CPF.Get(), "CPF mismatch")
	assert.Equal(t, "individual", ind.Type, "Expected type to be 'individual'")
	assert.Equal(t, email, ind.Email.Get(), "Email mismatch")
}

func TestNewIndividualInvalidCPF(t *testing.T) {
	cpf := "invalid-cpf"
	email := "valid-email@gmail.com"
	phones := []string{"911911911"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"Valid Name"}

	_, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidCPF, "Expected invalid CPF")
}

func TestNewIndividualInvalidEmail(t *testing.T) {
	cpf := "12345678909"
	email := "invalid-email"
	phones := []string{"911911911"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"Valid Name"}

	_, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email")
}

func TestNewIndividualInvalidPhone(t *testing.T) {
	cpf := "12345678909"
	email := "valid-email@gmail.com"
	phones := []string{"invalid-phone"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"Valid Name"}

	_, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected invalid phone")
}

func TestNewIndividualInvalidAddress(t *testing.T) {
	cpf := "12345678909"
	email := "valid-email@gmail.com"
	phones := []string{"911911911"}
	addresses := []string{"1"}
	contactNames := []string{"Valid Name"}

	_, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidAddress, "Expected invalid address")
}

func TestNewIndividualInvalidContactName(t *testing.T) {
	cpf := "12345678909"
	email := "valid-email@gmail.com"
	phones := []string{"911911911"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"ja"}

	_, err := user.NewIndividual(cpf, email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, user.ErrInvalidName, "Expected invalid contact name")
}
