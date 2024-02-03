package discountdb

// type repository struct {
// 	db *gorm.DB
// }

// func New(db *gorm.DB) repository {
// 	return repository{
// 		db: db,
// 	}
// }

// func (repo repository) Find(code string) (models.Discount, error) {

// 	var discount models.Discount
// 	if tx := repo.db.Where(models.Discount{DiscountCode: code}).First(&discount); tx.Error != nil {
// 		return models.Discount{}, tx.Error
// 	}

// 	return discount, nil
// }
