package userservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Search(identifier string) (cli user.User, err error) {
	return app.GetInstance().Repositories.User.Search(identifier)
}
