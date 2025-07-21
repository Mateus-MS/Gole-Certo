package usertestutils

import (
	"testing"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
)

func GetMock(t *testing.T) user.Individual {
	t.Helper()

	usr, _ := user.NewIndividual(
		"033.355.662-38",
		"teste@gmail.com",
		[]string{"911911911"},
		[]string{"bem ali meu cupadi"},
		[]string{"mateus alves de sousa"},
	)
	return usr
}

func GetMockRegistered(t *testing.T, app *testutils.Application) user.Individual {
	t.Helper()

	usr, _ := user.NewIndividual(
		"033.355.662-38",
		"teste@gmail.com",
		[]string{"911911911"},
		[]string{"bem ali meu cupadi"},
		[]string{"mateus alves de sousa"},
	)

	app.Services.User.Create(&usr)

	return usr
}

func CreateValidBase(t *testing.T) user.BaseUser {
	t.Helper()

	usr, _ := user.NewBaseUser(
		"teste@gmail.com",
		[]string{"911911911"},
		[]string{"bem ali meu cupadi"},
		[]string{"mateus alves de sousa"},
	)
	return usr
}

// Returns the same generic and valid user of type `individual`
func CreateValidIndividual(t *testing.T) user.Individual {
	t.Helper()

	usr, _ := user.NewIndividual(
		"033.355.662-38",
		"teste@gmail.com",
		[]string{"911911911"},
		[]string{"bem ali meu cupadi"},
		[]string{"mateus alves de sousa"},
	)
	return usr
}

// Returns the same generic and valid user of type `company`
func CreateValidCompany(t *testing.T) user.Company {
	t.Helper()

	usr, _ := user.NewCompany(
		"12.345.678/0001-95",
		"teste@gmail.com",
		"empresaaleatoriadobabado",
		"nossonomeéumafantasia",
		[]string{"911911911"},
		[]string{"bem ali meu cupadi"},
		[]string{"mateus alves de sousa"},
	)
	return usr
}

// Try to register the received user into received application service
func Register(t *testing.T, app *testutils.Application, usr user.User) (err error) {
	t.Helper()

	if err = app.Services.User.Create(usr); err != nil {
		return err
	}
	return nil
}
