package utils

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidIDType   = errors.New("only allows a string or a primitive.ObjectID")
	ErrInvalidIDFormat = errors.New("the received string id is invalid formated")
)

// Can receive either a `string` or a `primitive.ObjectID`
func ParseObjectID(objectID any) (ordID_obj primitive.ObjectID, err error) {
	var objID_obj primitive.ObjectID
	switch val := objectID.(type) {
	case string:
		objID_obj, err = primitive.ObjectIDFromHex(val)
		if err != nil {
			return objID_obj, ErrInvalidIDFormat
		}
	case primitive.ObjectID:
		objID_obj = val
	default:
		return ordID_obj, ErrInvalidIDType
	}

	return objID_obj, nil
}
