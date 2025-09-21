package user_repository

import "errors"

var (
	ErrUserInexistent = errors.New("user not found")
	ErrCannotConvert  = errors.New("cannot convert")
)
