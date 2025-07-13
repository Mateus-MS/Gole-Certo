package client

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"
)

type Company struct {
	CNPJ fields.CNPJ `json:"cnpj"`

	BaseClient

	FantasyName string
	LegalName   string
}

func NewCompany(cnpjRaw, emailRaw, phoneRaw, address, fantasyName, legalName string) (Company, error) {
	var (
		base   BaseClient
		client Company
		cnpj   fields.CNPJ
		err    error
	)

	if cnpj, err = fields.NewCNPJ(cnpjRaw); err != nil {
		return client, err
	}

	if base, err = NewBaseClient(emailRaw, phoneRaw, address); err != nil {
		return client, err
	}

	return Company{
		CNPJ:        cnpj,
		BaseClient:  base,
		FantasyName: fantasyName,
		LegalName:   legalName,
	}, nil
}

func (c *Company) GetIdentifier() string {
	return c.CNPJ.Get()
}
