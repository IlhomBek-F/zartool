package repositories

import (
	"context"
	"time"
	"zartool/models"

	"gorm.io/gorm"
)

func GetOwnerByLogin(db gorm.DB, login string) (models.Owner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var owner models.Owner

	result := db.WithContext(ctx).Where("login = ?", login).First(&owner)

	return owner, result.Error
}

func CreateOwner(db gorm.DB, owner models.Owner) error {
	result := db.Create(&owner)

	return result.Error
}
