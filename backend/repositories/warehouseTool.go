package repositories

import (
	"context"
	"time"
	"zartool/domain"

	"gorm.io/gorm"
)

func AddNewTool(db gorm.DB, tools *[]domain.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := db.WithContext(ctx).Create(&tools)

	return result.Error
}

func GetWareHouseTools(db gorm.DB, page int, pageSize int) ([]domain.WarehouseTools, domain.MetaModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var warehouseTools []domain.WarehouseTools
	var count int64
	var metaData domain.MetaModel

	if countResult := db.WithContext(ctx).Model(&domain.WarehouseTools{}).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Order("created_at DESC").Find(&warehouseTools)

	metaData.Page = page
	metaData.Total = count

	return warehouseTools, metaData, result.Error
}

func UpdateWareHouseTool(db gorm.DB, tool *domain.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := db.WithContext(ctx).Save(&tool)

	return result.Error
}

func DeleteWarehouseTool(db gorm.DB, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var deletedTool domain.WarehouseTools
	result := db.WithContext(ctx).Delete(&deletedTool, id)

	return result.Error
}
