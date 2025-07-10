package DAO

import (
	"database/sql"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
)

type IndividualRepository struct {
	DB *sql.DB
}

func (db *IndividualRepository) Register(client client.Individual) (err error) {
	return nil
}
