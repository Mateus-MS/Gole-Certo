package user

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user/fields"

type Company struct {
	CNPJ fields.CNPJ `json:"CNPJ" bson:"_id"`

	BaseUser `bson:",inline"`

	FantasyName fields.FantasyName `json:"FantasyName" bson:"fantasyName"`
	LegalName   fields.LegalName   `json:"LegalName"   bson:"legalName"`
}

func NewCompany(cnpjRaw, emailRaw, fantasyNameRaw, legalNameRaw string, phoneRaw, address, contactNames []string) (comp Company, err error) {
	var (
		base        BaseUser
		cnpj        fields.CNPJ
		fantasyName fields.FantasyName
		legalName   fields.LegalName
	)

	if cnpj, err = fields.NewCNPJ(cnpjRaw); err != nil {
		return comp, fields.ErrInvalidCNPJ
	}

	if fantasyName, err = fields.NewFantasyName(fantasyNameRaw); err != nil {
		return comp, fields.ErrInvalidName
	}
	if legalName, err = fields.NewLegalName(legalNameRaw); err != nil {
		return comp, fields.ErrInvalidName
	}

	if base, err = NewBaseUser(emailRaw, phoneRaw, address, contactNames); err != nil {
		return comp, err
	}

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
		return fields.ErrInvalidName
	}

	if c.FantasyName.Get() == "" {
		return fields.ErrInvalidName
	}

	if c.LegalName.Get() == "" {
		return fields.ErrInvalidName
	}

	return nil
}

func (c *Company) GetIdentifier() string {
	return c.CNPJ.Get()
}
