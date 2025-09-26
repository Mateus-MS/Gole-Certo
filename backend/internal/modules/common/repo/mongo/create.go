package generic_repository_mongo

import "context"

func (repo *GenericRepository[T]) Create(ctx context.Context, item T) error {
	_, err := repo.Collection.InsertOne(ctx, item)

	if err != nil {
		return err
	}

	return nil
}
