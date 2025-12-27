package order_error

import "errors"

var (
	ErrUnavaiableQuantity = errors.New("there is not enough items in stock to supply this order")
)
