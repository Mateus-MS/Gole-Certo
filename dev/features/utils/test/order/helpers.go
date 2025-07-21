package ordertestutils

import (
	order "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/order/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
)

func GetMock(prods []product.Product) order.Order {
	ord, _ := order.New(
		"batching",
		prods,
	)
	return ord
}
