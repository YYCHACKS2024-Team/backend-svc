package constant

import "errors"

var (
	ErrorCardNotFound  = errors.New("not found cart_id")
	ErrorCartProceeded = errors.New("cart is already proceeded or cancel")
)
