package persistence

import (
	"context"
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	// DB *mongo.Client --- Removing, if needed add aggain later :P
	Collection *mongo.Collection
}

func (repo *ClientRepository) Save(client client.Client) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), client); err != nil {
		return err
	}

	return nil
}

var (
	// query errors

	// internal erros
	ErrDocumentTypeUnkown = errors.New("document doesn't match any client type")
	ErrMissingTypeField   = errors.New("the queryied document doesn't has the type field")
)

func (repo *ClientRepository) Search(identifier string) (c client.Client, err error) {
	// The query
	filter := bson.M{"_id": identifier}

	var raw bson.Raw

	// Query in DB
	if err = repo.Collection.FindOne(context.TODO(), filter).Decode(&raw); err != nil {
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

			return nil, ErrDocumentTypeUnkown

		}
	}

	return nil, ErrMissingTypeField

}
