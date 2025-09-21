package user_error

import (
	"errors"
)

var (
	ErrUserInexistent = errors.New("user not found")
	ErrCannotConvert  = errors.New("cannot convert")
)
