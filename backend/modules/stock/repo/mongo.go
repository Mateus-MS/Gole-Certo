package stock_repository

import (
	generic_repository "alves.com/backend/modules/common/repo"
	stock_model "alves.com/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*generic_repository.GenericRepository[*stock_model.StockEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository.New[*stock_model.StockEntity](coll),
	}
}
