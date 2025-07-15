package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/google/uuid"
)

func Register(userID string, products []product.Product) (_ string, err error) {
	// 1 - Check if the received user exists
	if _, err = userservice.Search(userID); err != nil {
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
		"batching",          // State
		products,            // Products
	)

	// 4 - Save in DB
	if err = app.GetInstance().Repositories.Order.Save(ord); err != nil {
		return "", err
	}

	// 5 - Check if the batching is full
	// NOTE: Whenever a new order is registered, it is registered as `batching` which means that we waiting till we have 1000 products
	// waiting to be ordered from `Duff Beer`. The point is, whenever we register a new order, we need to check if now, we have enough products
	// in the waiting list

	return ord.OrderID, nil
}
