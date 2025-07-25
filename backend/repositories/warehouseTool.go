package repositories

import (
	"context"
	"time"
	"zartool/domain"

	"gorm.io/gorm"
)

type warehouseRepository struct {
	db gorm.DB
}

func NewWarehouseRepository(db gorm.DB) domain.WarehouseRepository {
	return warehouseRepository{db: db}
}

func (wr warehouseRepository) AddNewTool(tools *[]domain.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := wr.db.WithContext(ctx).Create(&tools)

	return result.Error
}

func (wr warehouseRepository) GetWareHouseTools(page int, pageSize int) ([]domain.WarehouseTools, domain.MetaModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var warehouseTools []domain.WarehouseTools
	var count int64
	var metaData domain.MetaModel

	if countResult := wr.db.WithContext(ctx).Model(&domain.WarehouseTools{}).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := wr.db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Order("created_at DESC").Find(&warehouseTools)

	metaData.Page = page
	metaData.Total = count

	return warehouseTools, metaData, result.Error
}

func (wr warehouseRepository) UpdateWareHouseTool(tool *domain.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := wr.db.WithContext(ctx).Save(&tool)

	return result.Error
}

func (wr warehouseRepository) DeleteWarehouseTool(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var deletedTool domain.WarehouseTools
	result := wr.db.WithContext(ctx).Delete(&deletedTool, id)

	return result.Error
}
