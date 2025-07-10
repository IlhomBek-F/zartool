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
	err := db.Model(&rental).Association("RentTools").Replace(rental.RentTools)

	if err != nil {
		return err
	}

	result := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&rental)

	return result.Error
}

func GetRentals(db gorm.DB) ([]models.User, error) {
	var rentals []models.User

	result := db.Preload("RentTools").Find(&rentals)

	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return rentals, nil
}
