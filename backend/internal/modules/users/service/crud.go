package user_service

import (
	"context"

	user_model "alves.com/modules/users/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) Create(ctx context.Context, user user_model.UserEntity) error {
	return s.repository.Create(ctx, &user)
}
func (s *service) ReadByName(ctx context.Context, name string) (user_model.UserEntity, error) {
	return s.repository.ReadByName(ctx, name)
}
func (s *service) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return s.repository.DeleteByID(ctx, id)
}
