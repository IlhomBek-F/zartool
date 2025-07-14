package repositories

import (
	"time"
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

func GetRentalReport(db gorm.DB, page int, pageSize int) (models.RentalReport, models.MetaModel, error) {
	var totalCompletedRent int64
	var totalCreatedRent int64
	var todayRents []models.User
	var meta = models.MetaModel{Page: page}

	totalCompletedRentResult := db.Model(&models.User{}).Where("active = ?", false).Count(&totalCompletedRent)

	if totalCompletedRentResult.Error != nil {
		return models.RentalReport{}, meta, totalCompletedRentResult.Error
	}

	totalCreatedRentResult := db.Model(&models.User{}).Where("active = ?", true).Count(&totalCreatedRent)

	if totalCreatedRentResult.Error != nil {
		return models.RentalReport{}, meta, totalCreatedRentResult.Error
	}

	rentsTotal, err := getTodayRents(db, &todayRents, page, pageSize)

	if err != nil {
		return models.RentalReport{}, meta, err
	}

	meta.Total = rentsTotal

	report := models.RentalReport{
		Total_created_rent:   totalCreatedRent,
		Total_completed_rent: totalCompletedRent,
		Rents:                todayRents,
	}

	return report, meta, nil
}

func GetRentals(db gorm.DB, page int, pageSize int) ([]models.User, models.MetaModel, error) {
	var rentals []models.User
	var count int64
	var metaData models.MetaModel

	if countResult := db.Model(&rentals).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.Scopes(Paginate(page, pageSize)).Preload("RentTools").Order("created_at DESC").Find(&rentals)

	if result.Error != nil {
		return nil, metaData, result.Error
	}

	metaData.Page = page
	metaData.Total = count

	return rentals, metaData, nil
}

func getTodayRents(db gorm.DB, todayRents *[]models.User, page int, pageSize int) (int64, error) {
	var total int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour).Format("02-01-2006 15:04")
	formatStartOfdDay := startOfDay.Format("02-01-2006 15:04")

	todayRentsTotalResult := db.Model(models.User{}).Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).Count(&total)

	if todayRentsTotalResult.Error != nil {
		return 0, todayRentsTotalResult.Error
	}

	todayRentsResult := db.Scopes(Paginate(page, pageSize)).
		Preload("RentTools").
		Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).
		Find(&todayRents)

	if todayRentsResult.Error != nil {
		return 0, todayRentsResult.Error
	}

	return total, nil
}
