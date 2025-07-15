package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type Individual struct {
	CPF        fields.CPF `json:"CPF" bson:"_id"`
	BaseClient `bson:",inline"`
	Age        int `json:"Age" bson:"age"`
}

func NewIndividual(cpfRaw, emailRaw string, age int, phoneRaw, address, contactNames []string) (Individual, error) {
	var (
		base   BaseClient
		client Individual
		cpf    fields.CPF
		err    error
	)

	if cpf, err = fields.NewCPF(cpfRaw); err != nil {
		return client, err
	}

	// Create the base client
	if base, err = NewBaseClient(emailRaw, phoneRaw, address, contactNames); err != nil {
		return client, err
	}

	return Individual{
		CPF:        cpf,
		BaseClient: base,
		Age:        age,
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
