package stock_service

import (
	"context"

	stock_model "alves.com/backend/modules/stock/model"
	stock_repository "alves.com/backend/modules/stock/repo"
	stock_repository_mongo "alves.com/backend/modules/stock/repo/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// DeductFromStock(context.Context, product.ProductStock, int64) error
	Register(context.Context, stock_model.StockEntity) error
	ReadByName(context.Context, string) (stock_model.StockEntity, error)
}

type service struct {
	repository stock_repository.IRepository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: stock_repository_mongo.New(coll),
	}
}
