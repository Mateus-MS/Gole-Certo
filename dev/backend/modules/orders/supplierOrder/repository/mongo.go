package supplierOrder_repository

import (
	"context"
	"errors"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Errors
var (
	ErrOrderNotFound = errors.New("supplier order does not exists in db")
)

// Alias
type Order = supplierOrder.SupplierOrder

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{Collection: coll}
}

func (repo *Repository) Create(ord Order) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Read(filter bson.M) (ord Order, err error) {
	if err = repo.Collection.FindOne(context.TODO(), filter).Decode(&ord); err != nil {
		return ord, ErrOrderNotFound
	}

	return ord, nil
}
func (repo *Repository) ReadMany(filter bson.M, limit int) ([]Order, error) {
	ctx := context.TODO()

	cur, err := repo.Collection.Find(ctx, filter, options.Find().SetLimit(int64(limit)))
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

func (repo *Repository) Update(ord Order, filter bson.M) (err error) {
	update := bson.M{
		"$set": bson.M{
			"products":      ord.Products,
			"state":         ord.State,
			"totalQuantity": ord.TotalQuantity,
		},
	}

	result, err := repo.Collection.UpdateOne(context.TODO(), filter, update)
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
