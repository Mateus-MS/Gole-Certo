package user

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
)

type Company struct {
	CNPJ fields.CNPJ `json:"CNPJ" bson:"_id"`

	BaseUser `bson:",inline"`

	FantasyName string `json:"FantasyName" bson:"fantasyName"`
	LegalName   string `json:"LegalName"   bson:"legalName"`
}

func NewCompany(cnpjRaw, emailRaw, fantasyName, legalName string, phoneRaw, address, contactNames []string) (comp Company, err error) {
	var (
		base BaseUser
		cnpj fields.CNPJ
	)

	if cnpj, err = fields.NewCNPJ(cnpjRaw); err != nil {
		return comp, fields.ErrInvalidCNPJ
	}

	if len(fantasyName) < 10 {
		return comp, ErrInvalidName
	}

	if len(legalName) < 10 {
		return comp, ErrInvalidName
	}

	if base, err = NewBaseUser(emailRaw, phoneRaw, address, contactNames); err != nil {
		return comp, err
	}
	base.Type = "company"

	return Company{
		CNPJ:        cnpj,
		BaseUser:    base,
		FantasyName: fantasyName,
		LegalName:   legalName,
	}, nil
}

func (c *Company) IsValid() error {
	if c.CNPJ.Get() == "" {
		return fields.ErrInvalidCNPJ
	}

	if c.Email.Get() == "" {
		return fields.ErrInvalidEmail
	}

	if len(c.Phone) < 1 {
		return fields.ErrInvalidPhone
	}

	if len(c.Address) < 1 {
		return fields.ErrInvalidAddress
	}

	if len(c.ContactNames) < 1 {
		return ErrInvalidName
	}

	if c.FantasyName == "" {
		return ErrInvalidName
	}

	if c.LegalName == "" {
		return ErrInvalidName
	}

	return nil
}

func (c *Company) GetIdentifier() string {
	return c.CNPJ.Get()
}
