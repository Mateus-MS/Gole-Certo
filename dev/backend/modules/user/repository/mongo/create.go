package user_repository

import "context"

func (repo *Repository) Create(usr User) (err error) {
	if _, err = repo.collection.InsertOne(context.TODO(), usr); err != nil {
		return err
	}

	return nil
}
