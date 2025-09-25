package stock_service

import (
	stock_repository "alves.com/modules/stock/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	Repo() *stock_repository.Repository // TODO: REMOVE THIS
}

type service struct {
	repository *stock_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: stock_repository.New(coll),
	}
}

func (s *service) Repo() *stock_repository.Repository {
	return s.repository
}
