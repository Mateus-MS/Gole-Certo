package stock_service

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	product_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/repository"
	product_utils "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Alias
type Stock = product.ProductStock

type service struct {
	repository product_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{repository: *product_repository.New(coll)}
}

func (s *service) Create(prod Stock) (err error) {
	if err = prod.IsValid(); err != nil {
		return err
	}

	// Search for any product with the same `name` into DB
	// NOTE: should not have two products with same name
	if _, err = s.repository.ReadByName(prod.Name); err == nil {
		return product.ErrDuplicated
	}

	return s.repository.Create(prod)
}

func (s *service) ReadByID(id string) (prod Stock, err error) {
	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetID(id).Build(); err != nil {
		return prod, err
	}

	// Perform the query
	if prod, err = s.repository.Read(filter); err != nil {
		return prod, err
	}

	return prod, nil
}
func (s *service) ReadByName(name string) (prod Stock, err error) {
	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetName(name).Build(); err != nil {
		return prod, err
	}

	// Perform the query
	if prod, err = s.repository.Read(filter); err != nil {
		return prod, err
	}

	return prod, nil
}

func (s *service) UpdateByID(prod Stock) (err error) {
	if err = prod.IsValid(); err != nil {
		return err
	}

	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetID(prod.ProductID).Build(); err != nil {
		return err
	}

	// Perform the query
	if err = s.repository.Update(prod, filter); err != nil {
		return err
	}

	return nil
}
func (s *service) UpdateByName(prod Stock) (err error) {
	if err = prod.IsValid(); err != nil {
		return err
	}

	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetName(prod.Name).Build(); err != nil {
		return err
	}

	// Perform the query
	if err = s.repository.Update(prod, filter); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteByID(id string) (err error) {
	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetID(id).Build(); err != nil {
		return err
	}

	// Perform the query
	if err = s.repository.Delete(filter); err != nil {
		return err
	}

	return nil
}
func (s *service) DeleteByName(name string) (err error) {
	// Build the query
	var filter bson.M
	if filter, err = product_utils.NewQueryFilter().SetName(name).Build(); err != nil {
		return err
	}

	// Perform the query
	if err = s.repository.Delete(filter); err != nil {
		return err
	}

	return nil
}

// Utils

func (s *service) ValidateProductByID(id string) bool {
	if _, err := s.ReadByID(id); err != nil {
		return false
	}
	return true
}

// Base functions

func (s *service) Read(filter bson.M) (prod Stock, err error) {
	// Perform the query
	if prod, err = s.repository.Read(filter); err != nil {
		return prod, err
	}

	return prod, nil
}

func (s *service) Delete(filter bson.M) (err error) {
	// Perform the query
	if err = s.repository.Delete(filter); err != nil {
		return err
	}

	return nil
}
