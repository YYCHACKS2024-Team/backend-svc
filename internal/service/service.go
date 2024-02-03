package service

import "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"

type User interface {
	GetAll() ([]models.User, error)
}
