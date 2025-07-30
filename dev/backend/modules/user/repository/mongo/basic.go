package user_repository

import (
	"context"
	"errors"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Simple search for documents on `DataBase` and return the raw bson.
func (repo *Repository) read(ctx context.Context, filter bson.M) (raw bson.Raw, err error) {
	if err = repo.collection.FindOne(ctx, filter).Decode(&raw); err != nil {
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
