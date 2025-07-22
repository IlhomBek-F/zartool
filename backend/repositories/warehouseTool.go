package repositories

import (
	"context"
	"time"
	"zartool/models"

	"gorm.io/gorm"
)

func AddNewTool(db gorm.DB, tools *[]models.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := db.WithContext(ctx).Create(&tools)

	return result.Error
}

func GetWareHouseTools(db gorm.DB, page int, pageSize int) ([]models.WarehouseTools, models.MetaModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var warehouseTools []models.WarehouseTools
	var count int64
	var metaData models.MetaModel

	if countResult := db.WithContext(ctx).Model(&models.WarehouseTools{}).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Order("created_at DESC").Find(&warehouseTools)

	metaData.Page = page
	metaData.Total = count

	return warehouseTools, metaData, result.Error
}

func UpdateWareHouseTool(db gorm.DB, tool *models.WarehouseTools) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := db.WithContext(ctx).Save(&tool)

	return result.Error
}

func DeleteWarehouseTool(db gorm.DB, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var deletedTool models.WarehouseTools
	result := db.WithContext(ctx).Delete(&deletedTool, id)

	return result.Error
}
