package product_service

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	product_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/repository"
	product_utils "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository product_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{repository: *product_repository.New(coll)}
}

func (s *service) Create(prod product.Product) (err error) {
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

func (s *service) Read(filter bson.M) (prod product.Product, err error) {
	// Perform the query
	if prod, err = s.repository.Read(filter); err != nil {
		return prod, err
	}

	return prod, nil
}
func (s *service) ReadByID(id string) (prod product.Product, err error) {
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
func (s *service) ReadByName(name string) (prod product.Product, err error) {
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

func (s *service) UpdateByID(prod product.Product) (err error) {
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
func (s *service) UpdateByName(prod product.Product) (err error) {
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

func (s *service) Delete(filter bson.M) (err error) {
	// Perform the query
	if err = s.repository.Delete(filter); err != nil {
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

func (s *service) ValidateList(prods []product.Product) bool {
	for _, prod := range prods {
		if _, err := s.ReadByName(prod.Name); err != nil {
			return false
		}
	}

	return true
}
