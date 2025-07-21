package repositories

import (
	"zartool/models"

	"gorm.io/gorm"
)

func GetOwnerByLogin(db gorm.DB, login string) (models.Owner, error) {
	var owner models.Owner

	result := db.Where("login = ?", login).First(&owner)

	return owner, result.Error
}

func CreateOwner(db gorm.DB, owner models.Owner) error {
	result := db.Create(&owner)

	return result.Error
}
