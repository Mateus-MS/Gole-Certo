package stock_error

import (
	"errors"
)

var (
	ErrProductInexistent = errors.New("product not found")
	ErrCannotConvert     = errors.New("cannot convert")
)
