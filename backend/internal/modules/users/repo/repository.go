package user_repository

import (
	"context"

	generic_repository "alves.com/modules/common/repo"
	user_model "alves.com/modules/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	ReadByName(context.Context, string) (user_model.UserEntity, error)

	generic_repository.IGenericRepository[*user_model.UserEntity]
}

type Repository struct {
	*generic_repository.GenericRepository[*user_model.UserEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository.New[*user_model.UserEntity](coll),
	}
}
