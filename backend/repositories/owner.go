package repositories

import (
	"context"
	"time"
	"zartool/domain"

	"gorm.io/gorm"
)

func GetOwnerByLogin(db gorm.DB, login string) (domain.Owner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var owner domain.Owner

	result := db.WithContext(ctx).Where("login = ?", login).First(&owner)

	return owner, result.Error
}

func CreateOwner(db gorm.DB, owner domain.Owner) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := db.WithContext(ctx).Create(&owner)

	return result.Error
}
