package user_service

import (
	"context"
	"time"

	user_model "alves.com/backend/modules/users/model"
)

func (s *service) SaveInCache(ctx context.Context, token string, userData user_model.UserCache, TTL time.Duration) error {
	return s.cache.Save(ctx, token, userData, TTL)
}

func (s *service) ReadFromCache(ctx context.Context, token string) (user_model.UserCache, error) {
	return s.cache.Read(ctx, token)
}
