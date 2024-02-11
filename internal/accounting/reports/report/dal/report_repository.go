package dal

import "gorm.io/gorm"


type GormReportRepository struct {
	db *gorm.DB
}

// func (repo *GormReportRepository) FindByID(id uuid.UUID) (*entities.Product, error) {
// 	var dbProduct Product
// 	if err := repo.db.First(&dbProduct, id).Error; err != nil {
// 		return nil, err
// 	}


// 	return FromDBProduct(&dbProduct), nil
// }
