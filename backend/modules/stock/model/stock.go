package stock_model

import generic_persistent "alves.com/backend/modules/common/model"

type StockEntity struct {
	generic_persistent.Persistent `json:"ID"`

	Name string `json:"name" binding:"required"`
}
