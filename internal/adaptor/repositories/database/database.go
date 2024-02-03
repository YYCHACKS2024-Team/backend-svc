package database

import (
	"time"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
)

type (

	// carts table
	CartRepository interface {
		Insert(models.Cart) error
		Find(cartId string, no int) (models.Cart, error)
		FindById(cartId string) ([]models.Cart, error)
		FindByCustomerIdAndStatus(customerId string, status string) ([]models.Cart, error)
		UpdateStatus(id string, no int, status string) error
		UpdateStatusById(id string, status string) error
		GetCartByDateMonth(date time.Time) ([]models.Cart, error)
	}

	// orders table
	OrderRepository interface {
		Insert(models.Order) error
		FindByCustomerIdAndStatus(customerId string, status string) ([]models.Order, error)
		GetAll() ([]models.Order, error)
		GetOrderByDateMonth(date time.Time) ([]models.Order, error)
		GetOrderByStatus(status string) ([]models.Order, error)
		UpdateOrderStatus(orderId string, status string) error
	}

	// discounts table
	DiscountRepository interface {
		Find(code string) (models.Discount, error)
	}
)
