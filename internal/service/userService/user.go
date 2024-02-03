package cartservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type userService struct {
}

func New() service.User {
	return userService{}
}

const ()

func (srv userService) GetAll() ([]models.User, error) {

	return nil, nil
}
