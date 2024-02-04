package database

import "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"

type (
	// carts          table
	UserRepository interface {
		GetAll() ([]models.User, error)
	}

	RooomRepository interface {
	}
	RoleRepository interface {
	}
	ConversationRepository interface {
	}
	MessageRepository interface {
	}
)
