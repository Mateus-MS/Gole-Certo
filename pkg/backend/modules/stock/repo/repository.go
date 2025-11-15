package stock_repository

import (
	"context"

	generic_repository "alves.com/backend/modules/common/repo"
	stock_model "alves.com/backend/modules/stock/model"
)

type IRepository interface {
	ReadByName(context.Context, string) (stock_model.StockEntity, error)

	generic_repository.IGenericRepository[*stock_model.StockEntity]
}
