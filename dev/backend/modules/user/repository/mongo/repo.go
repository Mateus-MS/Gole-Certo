package user_repository

import (
	"errors"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{collection: coll}
}

// Alias
type User = user.User

var (
	// query errors
	ErrUserNotFound = errors.New("user not found")

	// internal erros
	ErrDocumentTypeUnkown = errors.New("document doesn't match any user type")
	ErrMissingTypeField   = errors.New("the queryied document doesn't has the type field")
)
