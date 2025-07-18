package userservice_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user/fields"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	usertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/user"
	"github.com/stretchr/testify/assert"
)

func TestRegisterValidIndividual(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the user
	usr := usertestutils.CreateValidIndividual(t)

	// Register the user on MockDB
	err := usertestutils.Register(t, app, &usr)
	assert.NoError(t, err, "Expected none errors")
}

func TestRegisterValidCompany(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the user
	usr := usertestutils.CreateValidCompany(t)

	// Register the user on MockDB
	err := usertestutils.Register(t, app, &usr)
	assert.NoError(t, err, "Expected none errors")
}

// Normally, all users in the backend are created using the New() constructor functions,
// which perform full validation on all fields and only return a user instance if everything is valid.
// This ensures invalid users cannot be created through normal means.
//
// However, these tests are designed to verify that even if a corrupted or invalid user somehow
// bypasses the constructors—such as by manually instantiating the struct (e.g., user.Individual{...})
// with raw fields, which is discouraged and practically impossible for validated fields since
// they cannot be created without the constructor
// the backend service will still detect and reject such invalid users during registration.
func TestRegisterInvalidEmail(t *testing.T) {
	// app := testutils.SetupTest(t)

	_, err := user.NewBaseUser(
		"not-an-email-mai-brodi",
		[]string{"911911911"},
		[]string{"kakakakakakakaka"},
		[]string{"jose luis da silva"},
	)
	assert.ErrorIs(t, err, fields.ErrInvalidEmail)

}
