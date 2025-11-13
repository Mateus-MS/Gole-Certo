package user_repository

import "errors"

var (
	ErrUserInexistent = errors.New("user not found")
	ErrCannotConvert  = errors.New("cannot convert")
	ErrDuplicatedUser = errors.New("trying to register a user with an already existent username")
)
