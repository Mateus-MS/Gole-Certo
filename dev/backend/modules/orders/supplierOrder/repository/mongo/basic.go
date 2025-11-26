package supplierOrder_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) read(ctx context.Context, filter bson.M) (ord Order, err error) {
	if err = repo.collection.FindOne(ctx, filter).Decode(&ord); err != nil {
		return ord, ErrOrderNotFound
	}

	return ord, nil
}

func (repo *Repository) readMany(ctx context.Context, filter bson.M, limit int) ([]Order, error) {
	cur, err := repo.collection.Find(ctx, filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []Order
	for cur.Next(ctx) {
		var ord Order
		if err := cur.Decode(&ord); err != nil {
			return nil, err
		}
		results = append(results, ord)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *Repository) update(ctx context.Context, ord Order, filter bson.M) (err error) {
	update := bson.M{
		"$set": bson.M{
			"products":      ord.Products,
			"state":         ord.State,
			"totalQuantity": ord.TotalQuantity,
		},
	}

	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrOrderNotFound
		}
		return err
	}

	if result.MatchedCount == 0 {
		return ErrOrderNotFound
	}

	return nil
}
