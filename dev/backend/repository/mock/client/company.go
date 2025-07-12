package DAO

import (
	"database/sql"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
)

type MockCompanyRepository struct {
	Db *sql.DB
}

func (db *MockCompanyRepository) Save(client client.Company) (err error) {
	return nil
}
