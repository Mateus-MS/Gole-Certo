package user_repository

import (
	generic_repository "alves.com/backend/modules/common/repo"
	user_model "alves.com/backend/modules/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*generic_repository.GenericRepository[*user_model.UserEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository.New[*user_model.UserEntity](coll),
	}
}
