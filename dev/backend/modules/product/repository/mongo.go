package product_repository

import (
	"context"
	"errors"
	"time"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Alias
type Product = product.ProductStock

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{Collection: coll}
}

func (repo *Repository) Create(prod Product) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), prod); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Read(queryFilter bson.M) (prod Product, err error) {
	if err = repo.Collection.FindOne(context.TODO(), queryFilter).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prod, product.ErrProductInexistent
		}
	}

	return prod, nil
}

// Used for pagination, can receive filters, the start point of the query and how much prods to return
func (repo *Repository) ReadManyAfterID(filter bson.M, lastID string, limit int64) (prods []Product, err error) {

	// 1 - Define the time limit for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find().SetLimit(limit).SetSort(bson.D{{Key: "_id", Value: 1}})

	// 2 - If received a non empty lastID, add it to the query
	if lastID != "" {

		var objID primitive.ObjectID
		// Convert the received ID into how mongoDB expectes it to be
		if objID, err = primitive.ObjectIDFromHex(lastID); err != nil {
			return prods, err
		}

		// Add the id in the filer, allowing only IDs greater than the one received
		filter["_id"] = bson.M{"$gt": objID}
	}

	// 3 - Perform the query
	var result *mongo.Cursor
	if result, err = repo.Collection.Find(ctx, filter, findOptions); err != nil {
		return prods, err
	}
	defer result.Close(ctx)

	// 4 - Decode the result
	if err := result.All(ctx, &prods); err != nil {
		return nil, err
	}
	return prods, nil
}

func (repo *Repository) ReadByID(id string) (prod Product, err error) {
	// Convert the received ID into how mongoDB expectes it to be
	var objID primitive.ObjectID
	if objID, err = primitive.ObjectIDFromHex(id); err != nil {
		return prod, err
	}

	// Perform the query
	if err = repo.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prod, product.ErrProductInexistent
		}
	}

	return prod, nil
}

func (repo *Repository) ReadByName(name string) (prod Product, err error) {
	if err = repo.Collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prod, product.ErrProductInexistent
		}
	}

	return prod, nil
}

func (repo *Repository) Delete(queryFilter bson.M) (err error) {
	var result *mongo.DeleteResult
	if result, err = repo.Collection.DeleteOne(context.TODO(), queryFilter); err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return product.ErrProductInexistent
	}

	return nil
}

func (repo *Repository) Update(prod Product, filter bson.M) (err error) {
	update := bson.M{
		"$set": bson.M{
			"name":     prod.Name,
			"brand":    prod.Brand,
			"price":    prod.Price,
			"quantity": prod.Stock,
		},
	}

	result, err := repo.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return product.ErrProductInexistent
		}
		return err
	}

	if result.MatchedCount == 0 {
		return product.ErrProductInexistent
	}

	return nil
}
