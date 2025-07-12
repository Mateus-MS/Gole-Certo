package DAO

import (
	"database/sql"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
)

type MockIndividualRepository struct {
	DB *sql.DB
}

func (db *MockIndividualRepository) Register(client client.Individual) (err error) {
	return nil
}
