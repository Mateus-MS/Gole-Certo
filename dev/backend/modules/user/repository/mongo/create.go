package user_repository

import "context"

func (repo *Repository) Create(ctx context.Context, usr User) (err error) {
	if _, err = repo.collection.InsertOne(ctx, usr); err != nil {
		return err
	}

	return nil
}
