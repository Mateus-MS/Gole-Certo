package generic_repository_mongo

import (
	generic_persistent "alves.com/modules/common/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenericRepository[T generic_persistent.IPersistent] struct {
	Collection *mongo.Collection
}

func New[T generic_persistent.IPersistent](coll *mongo.Collection) *GenericRepository[T] {
	return &GenericRepository[T]{Collection: coll}
}
