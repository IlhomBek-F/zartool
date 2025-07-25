package repositories

import (
	"context"
	"time"
	"zartool/domain"

	"gorm.io/gorm"
)

type ownerRepository struct {
	db gorm.DB
}

func NewOwnerRepository(db gorm.DB) domain.OwnerRepository {
	return &ownerRepository{db: db}
}

func (or *ownerRepository) CreateOwner(owner domain.Owner) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := or.db.WithContext(ctx).Create(&owner)

	return result.Error
}

func (or *ownerRepository) GetOwnerByLogin(login string) (domain.Owner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var owner domain.Owner

	result := or.db.WithContext(ctx).Where("login = ?", login).First(&owner)

	return owner, result.Error
}
