package service

import "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"

type (
	User interface {
		GetAll() ([]models.User, error)
	}

	Roommate interface {
		Accept() error
		Decline() error
	}

	Chat interface {
		Store() error
		GetCoversationById() error
		GetMessageById() error
	}

	House interface {
		GetHouseById() error
	}
)
