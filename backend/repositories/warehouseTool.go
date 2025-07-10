package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func AddNewTool(db gorm.DB, tools []models.WarehouseTools) error {
	result := db.Create(&tools)

	return result.Error
}

func GetWareHouseTools(db gorm.DB) ([]models.WarehouseTools, error) {
	var warehouseTools []models.WarehouseTools

	result := db.Find(&warehouseTools)

	return warehouseTools, result.Error
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
