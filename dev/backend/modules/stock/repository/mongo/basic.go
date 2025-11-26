package stock_repository

import (
	"context"
	"errors"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) read(ctx context.Context, queryFilter bson.M) (prod Product, err error) {
	if err = repo.collection.FindOne(ctx, queryFilter).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prod, product.ErrProductInexistent
		}
	}

	return prod, nil
}

func (repo *Repository) delete(ctx context.Context, queryFilter bson.M) (err error) {
	var result *mongo.DeleteResult
	if result, err = repo.collection.DeleteOne(ctx, queryFilter); err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return product.ErrProductInexistent
	}

	return nil
}

func (repo *Repository) update(ctx context.Context, prod Product, filter bson.M) (err error) {
	update := bson.M{
		"$set": bson.M{
			"name":         prod.Name,
			"brand":        prod.Brand,
			"price":        prod.Price,
			"stock":        prod.Stock,
			"minthreshold": prod.MinThreshold,
			"maxstock":     prod.MaxStock,
		},
	}

	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return product.ErrProductInexistent
	}

	return nil
}
