package user_error

import "errors"

var (
	ErrUserInexistent        = errors.New("this user does not exists on DB")
	ErrUserAlreadyExists     = errors.New("this user already exists on DB")
	ErrErrInvalidCredentials = errors.New("invalid credentials")
)
