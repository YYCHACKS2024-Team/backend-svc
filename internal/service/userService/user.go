package userservice

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/helper/logger"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service"
	"go.uber.org/zap"
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

	resultList, err := srv.userRepo.GetAll()
	if err != nil {
		logger.Error("Can't get all users", zap.Error(err))
	}

	return resultList, nil
}

func (srv userSvc) GetUserById(userId string) (models.User, error) {

	result, err := srv.userRepo.GetUserById(userId)
	if err != nil {
		logger.Error("Can't get user by id", zap.Error(err))
	}

	return result, nil
}
