package DAO

import (
	"database/sql"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
)

type CompanyRepository struct {
	Db *sql.DB
}

func (db *CompanyRepository) Save(client client.Company) (err error) {
	return nil
}
