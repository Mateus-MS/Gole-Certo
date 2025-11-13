package generic_repository

import "errors"

var (
	ErrItemInexistent = errors.New("this generic item does not exists on DB")
)
