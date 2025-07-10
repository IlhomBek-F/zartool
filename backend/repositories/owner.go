package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func GetOwnerByLogin(db gorm.DB, login string) (models.Owners, error) {
	var owner models.Owners

	result := db.Where("login = ?", login).First(&owner)

	return owner, result.Error
}

func CreateOwner(db gorm.DB, owner models.Owners) error {
	result := db.Create(&owner)

	return result.Error
}
