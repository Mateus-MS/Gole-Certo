package stock_repository

import "context"

func (repo *Repository) Create(ctx context.Context, prod Product) (err error) {
	if _, err = repo.collection.InsertOne(ctx, prod); err != nil {
		return err
	}

	return nil
}
