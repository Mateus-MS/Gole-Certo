package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type Company struct {
	CNPJ fields.CNPJ

	BaseClient

	FantasyName string
	LegalName   string
}
