package client

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"
)

type Company struct {
	CNPJ fields.CNPJ `json:"CNPJ" bson:"_id"`

	BaseClient `bson:",inline"`

	FantasyName fields.FantasyName `json:"FantasyName" bson:"fantasyName"`
	LegalName   fields.LegalName   `json:"LegalName"   bson:"legalName"`
}

func NewCompany(cnpjRaw, emailRaw, fantasyNameRaw, legalNameRaw string, phoneRaw, address, contactNames []string) (Company, error) {
	var (
		base        BaseClient
		client      Company
		cnpj        fields.CNPJ
		err         error
		fantasyName fields.FantasyName
		legalName   fields.LegalName
	)

	if cnpj, err = fields.NewCNPJ(cnpjRaw); err != nil {
		return client, err
	}

	if fantasyName, err = fields.NewFantasyName(fantasyNameRaw); err != nil {
		return client, err
	}
	if legalName, err = fields.NewLegalName(legalNameRaw); err != nil {
		return client, err
	}

	if base, err = NewBaseClient(emailRaw, phoneRaw, address, contactNames); err != nil {
		return client, err
	}

	return Company{
		CNPJ:        cnpj,
		BaseClient:  base,
		FantasyName: fantasyName,
		LegalName:   legalName,
	}, nil
}

func (c *Company) IsValid() bool {
	if c.CNPJ.Get() == "" {
		return false
	}

	if c.Email.Get() == "" {
		return false
	}

	if len(c.Phone) < 1 {
		return false
	}

	if len(c.Address) < 1 {
		return false
	}

	if len(c.ContactNames) < 1 {
		return false
	}

	if c.FantasyName.Get() == "" {
		return false
	}

	if c.LegalName.Get() == "" {
		return false
	}

	return true
}

func (c *Company) GetIdentifier() string {
	return c.CNPJ.Get()
}
