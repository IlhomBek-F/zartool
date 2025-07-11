package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func CreateNewRental(db gorm.DB, rental models.User) error {
	result := db.Create(&rental)

	if result.Error != nil {
		return result.Error
	}

	result = db.Save(&rental)

	return result.Error
}

func UpdateRental(db gorm.DB, rental models.User) error {
	var existingTools []models.RentTools
	var updatedRentTools = rental.RentTools

	if err := db.Model(&rental).Association("RentTools").Find(&existingTools); err != nil {
		return err
	}

	toolsMap := make(map[uint]models.RentTools)

	for _, tool := range updatedRentTools {
		toolsMap[tool.ID] = tool
	}

	removedTools := []models.RentTools{}

	for _, currentTool := range existingTools {
		if _, ok := toolsMap[currentTool.ID]; !ok {
			removedTools = append(removedTools, currentTool)
		}
	}

	if err := db.Select("RentTools").Delete(removedTools).Error; err != nil {
		return err
	}

	if len(toolsMap) > 0 {
		for _, updatingTool := range toolsMap {
			if err := db.Model(&models.RentTools{}).Where("user_id = ? AND id = ?", rental.ID, updatingTool.ID).Updates(updatingTool); err.Error != nil {
				return err.Error
			}
		}
	}

	return db.Save(rental).Error
}

func GetRentals(db gorm.DB) ([]models.User, error) {
	var rentals []models.User

	result := db.Preload("RentTools").Find(&rentals)

	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return rentals, nil
}
