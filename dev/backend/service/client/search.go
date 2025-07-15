package clientservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Search(identifier string) (cli client.Client, err error) {
	return app.GetInstance().Repositories.Client.Search(identifier)
}
