package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func AddNewTool(db gorm.DB, tools []models.WarehouseTools) error {
	result := db.Create(&tools)

	return result.Error
}

func GetWareHouseTools(db gorm.DB, page int, pageSize int) ([]models.WarehouseTools, models.MetaModel, error) {
	var warehouseTools []models.WarehouseTools
	var count int64
	var metaData models.MetaModel

	if countResult := db.Model(&models.WarehouseTools{}).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.Scopes(Paginate(page, pageSize)).Order("created_at DESC").Find(&warehouseTools)

	metaData.Page = page
	metaData.Total = count

	return warehouseTools, metaData, result.Error
}

func UpdateWareHouseTool(db gorm.DB, tool models.WarehouseTools) error {
	result := db.Save(&tool)

	return result.Error
}

func DeleteWarehouseTool(db gorm.DB, id int) error {
	var deletedTool models.WarehouseTools
	result := db.Delete(&deletedTool, id)

	return result.Error
}
