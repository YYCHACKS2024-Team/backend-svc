package cartservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type cartService struct {
}

func New() service.CartService {
	return cartService{}
}

const ()

func (srv cartService) AddToCart() error {

	return nil
}
