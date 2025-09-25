package stock_service

import (
	"context"

	stock_model "alves.com/modules/stock/model"
	stock_repository "alves.com/modules/stock/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	Create(context.Context, stock_model.StockEntity) error
	ReadByName(context.Context, string) (stock_model.StockEntity, error)
}

type service struct {
	repository *stock_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: stock_repository.New(coll),
	}
}
