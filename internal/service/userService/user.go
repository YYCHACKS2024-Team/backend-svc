package userservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
)

type userSvc struct {
	userRepo database.UserRepository
	roleRepo database.RoleRepository
}

func New(userRepo database.UserRepository, roleRepo database.RoleRepository) service.User {
	return userSvc{userRepo: userRepo, roleRepo: roleRepo}
}

const ()

func (srv userSvc) GetAll() ([]models.User, error) {

	return nil, nil
}
