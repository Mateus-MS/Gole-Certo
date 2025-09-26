package generic_repository

import (
	"context"

	generic_persistent "alves.com/modules/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Alias
type iPersistent = generic_persistent.IPersistent

type IGenericRepository[T iPersistent] interface {
	Create(context.Context, T) error

	Read(context.Context, bson.M) (generic_persistent.IPersistent, error)

	Update(context.Context, bson.M, bson.M) error

	Delete(context.Context, bson.M) error
	DeleteByID(context.Context, primitive.ObjectID) error
}

type GenericRepository[T iPersistent] struct {
	Collection *mongo.Collection
}

func New[T iPersistent](coll *mongo.Collection) *GenericRepository[T] {
	return &GenericRepository[T]{Collection: coll}
}
