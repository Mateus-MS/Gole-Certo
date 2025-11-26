package supplierOrder_repository

import "context"

func (repo *Repository) Create(ctx context.Context, ord Order) (id string, err error) {
	if _, err = repo.collection.InsertOne(ctx, ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}
