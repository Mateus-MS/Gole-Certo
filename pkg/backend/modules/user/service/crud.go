package user_service

import (
	"context"

	user_model "alves.com/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) Create(ctx context.Context, user user_model.UserEntity) error {
	return s.repository.Create(ctx, user)
}
func (s *service) ReadByName(ctx context.Context, name string) (*user_model.UserEntity, error) {
	return s.repository.ReadByName(ctx, name)
}
func (s *service) ReadByID(ctx context.Context, id primitive.ObjectID) (*user_model.UserEntity, error) {
	return s.repository.ReadByID(ctx, id)
}
func (s *service) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return s.repository.DeleteByID(ctx, id)
}
