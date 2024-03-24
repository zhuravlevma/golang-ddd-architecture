package dal

import (
	"github.com/google/uuid"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/dal/orm"
	"github.com/zhuravlevma/golang-ddd-architecture/internal/warehouse/order-managment/warehouse/domain/entities"
	"gorm.io/gorm"
)

type WarehouseRepository struct {
	Db *gorm.DB
}

func (repo *WarehouseRepository) GetWarehouseWithOrder(warehouseId uuid.UUID, orderId uuid.UUID) (*entities.WarehouseEntity, error) {
	var dbWarehouse orm.WarehouseOrm
	if err := repo.Db.First(&dbWarehouse, warehouseId).Where("OrderId = ?", orderId).Error; err != nil {
		return nil, err
	}

	return WarehouseToDomain(&dbWarehouse), nil
}

func (repo *WarehouseRepository) GetWarehouseWithOrders(warehouseId uuid.UUID) (*entities.WarehouseEntity, error) {
	var dbWarehouse orm.WarehouseOrm
	if err := repo.Db.First(&dbWarehouse, warehouseId).Error; err != nil {
		return nil, err
	}

	return WarehouseToDomain(&dbWarehouse), nil
}

func (repo *WarehouseRepository) UpdateWarehouse(warehouse *entities.WarehouseEntity) error {
	dbWarehouse := WarehouseToOrm(warehouse)
	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&dbWarehouse).Where("id = ?", dbWarehouse.ID).Updates(dbWarehouse).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedWarehouse, err := repo.GetWarehouseWithOrders(dbWarehouse.ID)
	if err != nil {
		return err
	}
	*warehouse = *storedWarehouse
	return nil
}

func (repo *WarehouseRepository) CreateWarehouse(warehouse *entities.WarehouseEntity) error {
	dbWarehouse := WarehouseToOrm(warehouse)

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		if err := repo.Db.Create(dbWarehouse).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}

	storedWarehouse, err := repo.GetWarehouseWithOrders(dbWarehouse.ID)
	if err != nil {
		return err
	}
	*warehouse = *storedWarehouse
	return nil
}
