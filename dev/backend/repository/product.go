package repository

import (
	"context"
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrProductNotFound = errors.New("product does not exists in DB")
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (repo *ProductRepository) Create(prod product.Product) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), prod); err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Read(queryFilter bson.M) (prod product.Product, err error) {
	if err = repo.Collection.FindOne(context.TODO(), queryFilter).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prod, ErrProductNotFound
		}
	}

	return prod, nil
}

func (repo *ProductRepository) Delete(queryFilter bson.M) (err error) {
	if _, err = repo.Collection.DeleteOne(context.TODO(), queryFilter); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrProductNotFound
		}
	}

	return nil
}

func (repo *ProductRepository) Update(prod product.Product, filter bson.M) (err error) {
	if _, err = repo.Collection.UpdateOne(context.TODO(), filter, prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrProductNotFound
		}
	}

	return nil
}
