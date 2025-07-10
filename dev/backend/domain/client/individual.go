package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type Individual struct {
	CPF fields.CPF

	BaseClient

	Age int
}
