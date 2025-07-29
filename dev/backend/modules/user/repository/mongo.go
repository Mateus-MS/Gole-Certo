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

func (repo *Repository) ReadByID(ctx context.Context, identifier string) (c User, err error) {
	// Build the filter
	filter := bson.M{"_id": identifier}

	// Query
	raw, err := repo.read(ctx, filter)
	if err != nil {
		return c, err
	}

	// Decode
	return decodeUser(raw)

}

// Works almost identically as `ReadById` but instead of returning the user or some error, it only returns a bool
func (repo *Repository) HasUser(ctx context.Context, identifier string) bool {
	filter := bson.M{"_id": identifier}

	err := repo.Collection.FindOne(ctx, filter).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}

// Basic CRUD

// Simple search for documents on `DataBase` and return the raw bson.
func (repo *Repository) read(ctx context.Context, filter bson.M) (raw bson.Raw, err error) {
	if err = repo.Collection.FindOne(ctx, filter).Decode(&raw); err != nil {
		// If the cause of the error is, document not found
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Means the user don't exists
			return raw, ErrUserNotFound
		}

		// Another error
		return raw, err
	}

	// Success
	return raw, err
}

// Utils

func decodeUser(raw bson.Raw) (usr User, err error) {
	typeVal := raw.Lookup("type")
	if typeVal.Type != bson.TypeString {
		return nil, ErrMissingTypeField
	}

	var typeStr string
	if err := typeVal.Unmarshal(&typeStr); err != nil {
		return nil, err
	}

	switch typeStr {
	case "individual":
		var ind user.Individual
		if err := bson.Unmarshal(raw, &ind); err != nil {
			return nil, err
		}
		return &ind, nil
	case "company":
		var comp user.Company
		if err := bson.Unmarshal(raw, &comp); err != nil {
			return nil, err
		}
		return &comp, nil
	default:
		return nil, ErrDocumentTypeUnkown
	}
}
