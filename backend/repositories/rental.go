package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func CreateNewRental(db gorm.DB, rental models.User) error {
	if err := db.Create(&rental); err != nil {
		return err.Error
	}

	return db.Save(&rental).Error
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

	if len(removedTools) > 0 {
		if err := db.Select("RentTools").Delete(removedTools).Error; err != nil {
			return err
		}
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

func DeleteRental(db gorm.DB, rentalId uint) error {
	var user models.User
	user.ID = rentalId

	return db.Select("RentTools").Delete(&user).Error
}

func CompleteRental(db gorm.DB, rentalId uint) error {
	var user models.User
	user.ID = rentalId

	return db.Model(&user).Update("active", false).Error
}

func GetRentals(db gorm.DB, page int, pageSize int) ([]models.User, models.MetaModel, error) {
	var rentals []models.User
	var count int64
	var metaData models.MetaModel

	if countResult := db.Model(&rentals).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.Scopes(Paginate(page, pageSize)).Preload("RentTools").Find(&rentals)

	if result.Error != nil {
		return nil, metaData, result.Error
	}

	metaData.Page = page
	metaData.Total = count

	return rentals, metaData, nil
}
