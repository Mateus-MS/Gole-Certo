package generic_repository

import "context"

func (repo *GenericRepository[T]) Create(ctx context.Context, item T) error {
	_, err := repo.collection.InsertOne(ctx, item)

	if err != nil {
		return err
	}

	return nil
}
