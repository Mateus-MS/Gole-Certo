package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	clientservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/client"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/google/uuid"
)

func Register(userID string, products []product.Product) (_ string, err error) {
	// 1 - Check if the received user exists
	if _, err = clientservice.Search(userID); err != nil {
		return "", err
	}

	// 2 - Check if the received product list match existing products
	// NOTE: currently, it's not checking, it's using a mock, always returning true :P
	for _, product := range products {
		if _, err = productservice.Search(product.ProductID); err != nil {
			return "", err
		}
	}

	// 3 - Create the structure to save in DB
	ord := order.New(
		userID,              // UserID
		uuid.New().String(), // OrderID
		"processing",        // State
		products,            // Products
	)

	// 4 - Save in DB
	if err = app.GetInstance().Repositories.Order.Save(ord); err != nil {
		return "", err
	}

	return ord.OrderID, nil
}
