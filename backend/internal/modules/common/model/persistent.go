package generic_persistent

import "go.mongodb.org/mongo-driver/bson/primitive"

type Persistent struct {
	ID primitive.ObjectID `json:"ID" bson:"_id"`
}

// I know that this maybe considered a bad habit from JAVA, but i'm from JAVA :P
type IPersistent interface {
	GetID() primitive.ObjectID
}

func (p Persistent) GetID() primitive.ObjectID {
	return p.ID
}
