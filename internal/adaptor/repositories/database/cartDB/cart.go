package cartdb

// type repository struct {
// 	db *gorm.DB
// }

// func New(db *gorm.DB) repository {
// 	return repository{
// 		db: db,
// 	}
// }

// func (repo repository) Insert(cart models.Cart) error {

// 	if tx := repo.db.Create(&cart); tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }

// func (repo repository) Find(cartId string, no int) (models.Cart, error) {

// 	var cart models.Cart
// 	if tx := repo.db.Where(models.Cart{CartId: cartId, No: no}).First(&cart); tx.Error != nil {
// 		return models.Cart{}, tx.Error
// 	}

// 	return cart, nil
// }

// func (repo repository) FindById(cartId string) ([]models.Cart, error) {

// 	var carts []models.Cart

// 	if tx := repo.db.Where(models.Cart{CartId: cartId}).Find(&carts); tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return carts, nil
// }

// func (repo repository) FindByCustomerIdAndStatus(customerId string, status string) ([]models.Cart, error) {

// 	var carts []models.Cart
// 	tx := repo.db.Where("customer_id = ? AND status = ?", customerId, status).Find(&carts)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return carts, nil
// }

// func (repo repository) UpdateStatus(id string, no int, status string) error {

// 	if tx := repo.db.Model(&models.Cart{}).Where(models.Cart{CartId: id, No: no}).Update("status", status); tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }

// func (repo repository) UpdateStatusById(id string, status string) error {

// 	if tx := repo.db.Model(&models.Cart{}).Where(models.Cart{CartId: id}).Updates(models.Cart{Status: status}); tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }

// func (repo repository) GetCartByDateMonth(date time.Time) ([]models.Cart, error) {

// 	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
// 	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

// 	var carts []models.Cart
// 	if err := repo.db.Where("date BETWEEN ? AND ?", startOfMonth, endOfMonth).Find(&carts).Error; err != nil {
// 		return nil, err
// 	}

// 	return carts, nil
// }
