package clientservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Register(cli client.Client) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = cli.IsValid(); err != nil {
		return err
	}

	return app.GetInstance().Repositories.Client.Save(cli)
}
