package generic_repository

import (
	generic_persistent "alves.com/backend/modules/common/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Alias
type iPersistent = generic_persistent.IPersistent

// I know that this approach is not much a "go" way of doing things but hey i'm just messing with concepts.

// This `GenericRepository` carry the basic C-R-U-D methods
type GenericRepository[T iPersistent] struct {
	collection *mongo.Collection
}

func New[T iPersistent](coll *mongo.Collection) *GenericRepository[T] {
	return &GenericRepository[T]{collection: coll}
}
