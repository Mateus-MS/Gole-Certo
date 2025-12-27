package stock_service

import (
	"context"

	stock_model "alves.com/backend/modules/stock/model"
	stock_repository "alves.com/backend/modules/stock/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	Create(context.Context, stock_model.StockEntity) error
	ReadByName(context.Context, string) (*stock_model.StockEntity, error)
	ReadByID(context.Context, primitive.ObjectID) (*stock_model.StockEntity, error)
	DeleteByID(context.Context, primitive.ObjectID) error
	UpdateByID(context.Context, stock_model.StockEntity) error
	AtomicDecreaseStockByID(context.Context, primitive.ObjectID, int) error
}

type service struct {
	repository *stock_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: stock_repository.New(coll),
	}
}
