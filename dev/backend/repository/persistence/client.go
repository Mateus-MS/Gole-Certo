package persistence

import (
	"context"
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	DB *mongo.Client
}

func (db *ClientRepository) Save(client client.Client) (err error) {
	collection := db.DB.Database("goleCertoDB").Collection("clients")

	if _, err = collection.InsertOne(context.TODO(), client); err != nil {
		return err
	}

	return nil
}

var (
	// query errors

	// internal erros
	ErrorDocumentTypeUnkown = errors.New("document doesn't match any client type")
	ErrorMissingTypeField   = errors.New("the queryied document doesn't has the type field")
)

func (db *ClientRepository) Search(identifier string) (c client.Client, err error) {
	// TODO: Store the collection inside `db`
	collection := db.DB.Database("goleCertoDB").Collection("clients")

	// The query
	filter := bson.M{"_id": identifier}

	var raw bson.Raw

	// Query in DB
	if err = collection.FindOne(context.TODO(), filter).Decode(&raw); err != nil {
		return c, err
	}

	// Get the data for the field type
	typeVal := raw.Lookup("type")
	// Check if is of type string
	if typeVal.Type == bson.TypeString {
		var typeStr string

		if err = typeVal.Unmarshal(&typeStr); err == nil {

			if typeStr == "individual" {
				var ind client.Individual
				// Unmarshal the whole struct and return it
				if err = bson.Unmarshal(raw, &ind); err == nil {
					return &ind, nil
				}
			}

			if typeStr == "company" {
				var comp client.Company
				// Unmarshal the whole struct and return it
				if err = bson.Unmarshal(raw, &comp); err == nil {
					return &comp, nil
				}
			}

			return nil, ErrorDocumentTypeUnkown

		}
	}

	return nil, ErrorMissingTypeField

}
