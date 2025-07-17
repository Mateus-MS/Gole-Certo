package user

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user/fields"

type Individual struct {
	CPF      fields.CPF `json:"CPF" bson:"_id"`
	BaseUser `bson:",inline"`
	Age      int `json:"Age" bson:"age"`
}

func NewIndividual(cpfRaw, emailRaw string, age int, phoneRaw, address, contactNames []string) (usr Individual, err error) {
	var (
		base BaseUser
		cpf  fields.CPF
	)

	if cpf, err = fields.NewCPF(cpfRaw); err != nil {
		return usr, err
	}

	// Create the base user
	if base, err = NewBaseUser(emailRaw, phoneRaw, address, contactNames); err != nil {
		return usr, err
	}
	base.Type = "individual"

	return Individual{
		CPF:      cpf,
		BaseUser: base,
		Age:      age,
	}, nil
}

func (c *Individual) IsValid() error {
	if c.CPF.Get() == "" {
		return fields.ErrInvalidCPF
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

	return nil
}

func (c *Individual) GetIdentifier() string {
	return c.CPF.Get()
}
