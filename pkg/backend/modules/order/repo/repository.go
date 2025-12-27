package order_repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		Collection: coll,
	}
}
