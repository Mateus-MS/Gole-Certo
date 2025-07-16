package repository

import (
	"context"
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (repo *UserRepository) Save(usr user.User) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), usr); err != nil {
		return err
	}

	return nil
}

var (
	// query errors

	// internal erros
	ErrDocumentTypeUnkown = errors.New("document doesn't match any user type")
	ErrMissingTypeField   = errors.New("the queryied document doesn't has the type field")
)

func (repo *UserRepository) Search(identifier string) (c user.User, err error) {
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

			switch typeStr {
			case "individual":
				var ind user.Individual
				// Unmarshal the whole struct and return it
				if err = bson.Unmarshal(raw, &ind); err == nil {
					return &ind, nil
				}
			case "company":
				var comp user.Company
				// Unmarshal the whole struct and return it
				if err = bson.Unmarshal(raw, &comp); err == nil {
					return &comp, nil
				}
			default:
				return nil, ErrDocumentTypeUnkown
			}

		}
	}

	return nil, ErrMissingTypeField

}
