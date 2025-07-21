package product_utils

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: REMOVE ALL OF IT

type queryFilter struct {
	Name string
	ID   primitive.ObjectID
	errs []error
}

func NewQueryFilter() *queryFilter {
	return &queryFilter{}
}

func (qf *queryFilter) SetID(id any) *queryFilter {
	switch idTyped := id.(type) {
	case string:
		idOBJ, err := primitive.ObjectIDFromHex(idTyped)
		if err != nil {
			qf.errs = append(qf.errs, product.ErrInvalidID)
		}
		qf.ID = idOBJ
	case primitive.ObjectID:
		qf.ID = idTyped
	default:
		qf.errs = append(qf.errs, product.ErrInvalidID)
	}

	return qf
}

func (qf *queryFilter) SetName(name string) *queryFilter {
	qf.Name = name
	return qf
}

func (qf *queryFilter) Build() (bson.M, error) {
	query := bson.M{}

	if len(qf.errs) > 0 {
		return query, qf.errs[0]
	}

	if qf.ID != primitive.NilObjectID {
		query["_id"] = qf.ID
	}

	if qf.Name != "" {
		query["name"] = qf.Name
	}

	return query, nil
}
