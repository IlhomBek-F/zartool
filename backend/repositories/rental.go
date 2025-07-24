package repositories

import (
	"context"
	"time"
	"zartool/domain"
	"zartool/internal"

	"gorm.io/gorm"
)

func CreateNewRental(db gorm.DB, rentalPayload *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := db.WithContext(ctx).Create(&rentalPayload); err != nil {
		return err.Error
	}

	return db.Save(&rentalPayload).Error
}

func UpdateRental(db gorm.DB, rental *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var existingTools []domain.RentTools
	var updatedRentTools = rental.RentTools

	if err := db.WithContext(ctx).Model(&rental).Association("RentTools").Find(&existingTools); err != nil {
		return err
	}

	toolsMap := make(map[uint]domain.RentTools)

	for _, tool := range updatedRentTools {
		toolsMap[tool.ID] = tool
	}

	removedTools := []domain.RentTools{}

	for _, currentTool := range existingTools {
		if _, ok := toolsMap[currentTool.ID]; !ok {
			removedTools = append(removedTools, currentTool)
		}
	}

	transactionError := internal.WithTransaction(ctx, &db, func(tx *gorm.DB) error {
		if len(removedTools) > 0 {
			if err := tx.Select("RentTools").Delete(removedTools).Error; err != nil {
				return err
			}
		}

		if len(toolsMap) > 0 {
			for _, updatingTool := range toolsMap {
				if err := tx.WithContext(ctx).
					Model(&domain.RentTools{}).
					Where("user_id = ? AND id = ?", rental.ID, updatingTool.ID).
					Updates(updatingTool); err.Error != nil {
					return err.Error
				}
			}
		}

		return nil
	})

	if transactionError != nil {
		return transactionError
	}

	return db.WithContext(ctx).Save(rental).Error
}

func DeleteRental(db gorm.DB, rentalId uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var user domain.User
	user.ID = rentalId

	return db.WithContext(ctx).Select("RentTools").Delete(&user).Error
}

func CompleteRental(db gorm.DB, rentalId uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var user domain.User
	user.ID = rentalId

	return db.WithContext(ctx).Model(&user).Update("active", false).Error
}

func GetRentalReport(db gorm.DB, page int, pageSize int, queryTerm string) (domain.RentalReport, domain.MetaModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var totalCompletedRent int64
	var totalCreatedRent int64
	var todayRents []domain.User
	var meta = domain.MetaModel{Page: page}

	totalCompletedRentResult := db.WithContext(ctx).Model(&domain.User{}).Where("active = ?", false).Count(&totalCompletedRent)

	if totalCompletedRentResult.Error != nil {
		return domain.RentalReport{}, meta, totalCompletedRentResult.Error
	}

	totalCreatedRentResult := db.WithContext(ctx).Model(&domain.User{}).Where("active = ?", true).Count(&totalCreatedRent)

	if totalCreatedRentResult.Error != nil {
		return domain.RentalReport{}, meta, totalCreatedRentResult.Error
	}

	rentsTotal, err := getTodayRents(*db.WithContext(ctx), &todayRents, page, pageSize, queryTerm)

	if err != nil {
		return domain.RentalReport{}, meta, err
	}

	meta.Total = rentsTotal

	report := domain.RentalReport{
		Total_created_rent:   totalCreatedRent,
		Total_completed_rent: totalCompletedRent,
		Rents:                todayRents,
	}

	return report, meta, nil
}

func GetRentals(db gorm.DB, page int, pageSize int, queryTerm string) ([]domain.User, domain.MetaModel, error) {
	var rentals []domain.User
	var count int64
	var metaData domain.MetaModel

	if countResult := db.Model(&rentals).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := db.Scopes(Paginate(page, pageSize)).
		Preload("RentTools").
		Where("full_name ILIKE ? OR phones LIKE ?", "%"+queryTerm+"%", "%"+queryTerm+"%").
		Order("created_at DESC").
		Find(&rentals)

	if result.Error != nil {
		return nil, metaData, result.Error
	}

	metaData.Page = page
	metaData.Total = count

	return rentals, metaData, nil
}

func getTodayRents(db gorm.DB, todayRents *[]domain.User, page int, pageSize int, queryTerm string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var total int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour).Format("02-01-2006 15:04")
	formatStartOfdDay := startOfDay.Format("02-01-2006 15:04")

	todayRentsTotalResult := db.WithContext(ctx).Model(domain.User{}).Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).Count(&total)

	if todayRentsTotalResult.Error != nil {
		return 0, todayRentsTotalResult.Error
	}

	todayRentsResult := db.WithContext(ctx).Scopes(Paginate(page, pageSize)).
		Preload("RentTools").
		Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).
		Where("full_name ILIKE ? OR phones LIKE ?", "%"+queryTerm+"%", "%"+queryTerm+"%").
		Find(&todayRents)

	if todayRentsResult.Error != nil {
		return 0, todayRentsResult.Error
	}

	return total, nil
}
