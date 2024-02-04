package userdb

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}

func (repo repository) GetAll() ([]models.User, error) {

	var users []models.User
	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (repo repository) GetUserById(userId string) (models.User, error) {

	var user models.User
	tx := repo.db.Where(models.User{UserId: userId}).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return user, nil
}
