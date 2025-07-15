package userservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Register(cli user.User) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = cli.IsValid(); err != nil {
		return err
	}

	return app.GetInstance().Repositories.User.Save(cli)
}
