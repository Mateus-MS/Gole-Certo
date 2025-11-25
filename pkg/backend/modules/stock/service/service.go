package stock_service

import (
	"context"

	user_model "alves.com/backend/modules/user/model"
	user_repository "alves.com/backend/modules/user/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	Create(context.Context, user_model.UserEntity) error
	ReadByName(context.Context, string) (*user_model.UserEntity, error)
	DeleteByID(context.Context, primitive.ObjectID) error
}

type service struct {
	repository *user_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: user_repository.New(coll),
	}
}
