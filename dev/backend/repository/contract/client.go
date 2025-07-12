package contract

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"

type CompanyRepository interface {
	Save(client *client.Company) error
	FindByCNPJ(cnpj string) (*client.Company, error)
}

type IndividualRepository interface {
	Save(client *client.Individual) error
	FindByCPF(cpf string) (*client.Individual, error)
}
