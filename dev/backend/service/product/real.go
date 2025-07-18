package productservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type service struct {
	repository repository.ProductRepository
}

func New(repo repository.ProductRepository) *service {
	return &service{repository: repo}
}

func (s *service) Create(prod product.Product) (err error) {
	return s.repository.Create(prod)
}

type QueryFilter struct {
	Name string
	ID   string
}

func (s *service) Read(filter QueryFilter) (prod product.Product, err error) {
	queryFilter := bson.M{}

	// Dinamically build the filter
	if filter.Name != "" {
		queryFilter["name"] = filter.Name
	}

	if filter.ID != "" {
		queryFilter["_id"] = filter.ID
	}

	// Perform the query
	if prod, err = s.repository.Read(queryFilter); err != nil {
		return prod, err
	}

	return prod, nil
}

func (s *service) Delete(filter QueryFilter) (err error) {
	queryFilter := bson.M{}

	// Dinamically build the filter
	if filter.Name != "" {
		queryFilter["name"] = filter.Name
	}

	if filter.ID != "" {
		queryFilter["_id"] = filter.ID
	}

	// Perform the query
	if err = s.repository.Delete(queryFilter); err != nil {
		return err
	}

	return nil
}

func (s *service) Update(prod product.Product, filter QueryFilter) (err error) {
	queryFilter := bson.M{}

	// Dinamically build the filter
	if filter.Name != "" {
		queryFilter["name"] = filter.Name
	}

	if filter.ID != "" {
		queryFilter["_id"] = filter.ID
	}

	// Perform the query
	if err = s.repository.Update(prod, queryFilter); err != nil {
		return err
	}

	return nil
}
