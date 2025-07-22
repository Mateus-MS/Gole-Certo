package costumerOrder_service

import costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"

type Service interface {
	Register(costumerOrder.CostumerOrder) (string, error)

	// C R U D
	Create(costumerOrder.CostumerOrder) (string, error)
}
