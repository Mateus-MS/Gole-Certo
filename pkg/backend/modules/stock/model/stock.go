package stock_model

import generic_persistent "alves.com/backend/modules/common/model"

type StockEntity struct {
	generic_persistent.Persistent `bson:",inline"`

	Name string `json:"name" binding:"required"`
}
