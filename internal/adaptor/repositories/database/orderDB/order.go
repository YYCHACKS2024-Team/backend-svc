package orderdb

import (
	"time"

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

func (repo repository) Insert(order models.Order) error {

	if tx := repo.db.Create(&order); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo repository) FindByCustomerIdAndStatus(customerId string, status string) ([]models.Order, error) {

	var orders []models.Order
	tx := repo.db.Where("customer_id = ? AND status = ?", customerId, status).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orders, nil
}

func (repo repository) GetAll() ([]models.Order, error) {

	var orders []models.Order
	tx := repo.db.Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orders, nil
}

func (repo repository) GetOrderByDateMonth(date time.Time) ([]models.Order, error) {

	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	var orders []models.Order
	if err := repo.db.Where("YEAR(date_time) = ? AND MONTH(date_time) = ? AND date_time BETWEEN ? AND ?", date.Year(), date.Month(), startOfMonth, endOfMonth).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo repository) GetOrderByStatus(status string) ([]models.Order, error) {

	var orders []models.Order
	tx := repo.db.Where(models.Order{Status: status}).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orders, nil
}

func (repo repository) UpdateOrderStatus(orderId string, status string) error {

	tx := repo.db.Model(&models.Order{}).Where(models.Order{OrderId: orderId}).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
