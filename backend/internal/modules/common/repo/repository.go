package generic_repository

import (
	"context"

	generic_persistent "alves.com/modules/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IGenericRepository[T generic_persistent.IPersistent] interface {
	Create(context.Context, T) error

	Read(context.Context, bson.M) (generic_persistent.IPersistent, error)

	Update(context.Context, bson.M, bson.M) error

	Delete(context.Context, bson.M) error
	DeleteByID(context.Context, primitive.ObjectID) error
}
