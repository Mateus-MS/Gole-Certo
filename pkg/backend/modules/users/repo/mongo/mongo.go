package user_repository_mongo

import (
	generic_repository_mongo "alves.com/backend/modules/common/repo/mongo"
	user_model "alves.com/backend/modules/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*generic_repository_mongo.GenericRepository[*user_model.UserEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository_mongo.New[*user_model.UserEntity](coll),
	}
}
