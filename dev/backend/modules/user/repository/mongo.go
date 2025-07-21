package user_repository

import (
	"context"
	"errors"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{Collection: coll}
}

// Alias
type User = user.User

func (repo *Repository) Create(usr User) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), usr); err != nil {
		return err
	}

	return nil
}

var (
	// query errors
	ErrUserNotFound = errors.New("user not found")

	// internal erros
	ErrDocumentTypeUnkown = errors.New("document doesn't match any user type")
	ErrMissingTypeField   = errors.New("the queryied document doesn't has the type field")
)

func (repo *Repository) Read(identifier string) (c User, err error) {
	// The query
	filter := bson.M{"_id": identifier}

	var raw bson.Raw

	// Query in DB
	if err = repo.Collection.FindOne(context.TODO(), filter).Decode(&raw); err != nil {
		// If the cause of the error is, document not found
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Means the user don't exists
			return c, ErrUserNotFound
		}

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
