package user_test

import (
	"testing"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseUserInvalidEmail(t *testing.T) {
	email := "invalid-email"
	phones := []string{"911911911"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"jajajajajajajajajajaja"}

	_, err := user.NewBaseUser(email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email")
}

func TestNewBaseUserInvalidPhone(t *testing.T) {
	email := "valid-email@gmail.com"
	phones := []string{"911911911", "988988988"}
	addresses := []string{"123 Main St"}
	contactNames := []string{"jajajajajajajajajajaja"}

	_, err := user.NewBaseUser(email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected invalid phone")
}

func TestNewBaseUserInvalidAddress(t *testing.T) {
	email := "valid-email@gmail.com"
	phones := []string{"911911911", "961911911"}
	addresses := []string{"123"}
	contactNames := []string{"jajajajajajajajajajaja"}

	_, err := user.NewBaseUser(email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidAddress, "Expected invalid address")
}

func TestNewBaseUserInvalidContactname(t *testing.T) {
	email := "valid-email@gmail.com"
	phones := []string{"911911911", "961911911"}
	addresses := []string{"123123123"}
	contactNames := []string{"ja"}

	_, err := user.NewBaseUser(email, phones, addresses, contactNames)
	assert.ErrorIs(t, err, user.ErrInvalidName, "Expected invalid contact name")
}
