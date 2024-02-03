package cartservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type cartService struct {
	cartRepo database.CartRepository
}

func New(cartRepo database.CartRepository) service.CartService {
	return cartService{
		cartRepo: cartRepo,
	}
}

const ()

func (srv cartService) AddToCart() error {

	return nil
}
