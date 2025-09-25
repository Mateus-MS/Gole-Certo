package user_service

import (
	"context"

	user_model "alves.com/modules/users/model"
)

// Wraps all crud methods from repo
func (s *service) ReadByName(ctx context.Context, name string) (user_model.UserEntity, error) {
	return s.repository.ReadByName(ctx, name)
}
