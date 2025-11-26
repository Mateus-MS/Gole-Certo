package user_service

import (
	"context"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
)

func (s *service) Register(ctx context.Context, usr user.User) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = usr.IsValid(); err != nil {
		return err
	}

	return s.Repo().Create(ctx, usr)
}
