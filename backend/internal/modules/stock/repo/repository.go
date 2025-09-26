package stock_repository

import (
	"context"

	generic_repository "alves.com/modules/common/repo"
	stock_model "alves.com/modules/stock/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	ReadByName(context.Context, string) (stock_model.StockEntity, error)

	generic_repository.IGenericRepository[*stock_model.StockEntity]
}

type Repository struct {
	*generic_repository.GenericRepository[*stock_model.StockEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository.New[*stock_model.StockEntity](coll),
	}
}
