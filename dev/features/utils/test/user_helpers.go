package testutils

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
)

// Returns the same generic and valid user of type `individual`
func CreateValidUserIndividual(t *testing.T) user.Individual {
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
func CreateValidUserCompany(t *testing.T) user.Company {
	t.Helper()

	usr, _ := user.NewCompany(
		"79.379.491/0001-15",
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
func RegisterUser(t *testing.T, app *Application, usr user.User) (err error) {
	t.Helper()

	if err = app.Services.User.Create(usr); err != nil {
		return err
	}
	return nil
}
