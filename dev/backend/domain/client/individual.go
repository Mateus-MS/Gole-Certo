package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type Individual struct {
	CPF        fields.CPF `json:"CPF" bson:"_id"`
	BaseClient `bson:",inline"`
	Age        int `json:"Age" bson:"age"`
}

func NewIndividual(cpfRaw, emailRaw, phoneRaw, address string, age int) (Individual, error) {
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
	if base, err = NewBaseClient(emailRaw, phoneRaw, address); err != nil {
		return client, err
	}

	return Individual{
		CPF:        cpf,
		BaseClient: base,
		Age:        age,
	}, nil
}

func (c *Individual) GetIdentifier() string {
	return c.CPF.Get()
}
