package user_service

import "errors"

var (
	ErrInvalidCredentials = errors.New("the provided credentials are invalids")
)
