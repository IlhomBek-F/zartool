package repositories

import (
	"context"
	"time"
	"zartool/domain"
	"zartool/internal"

	"gorm.io/gorm"
)

type rentalRepository struct {
	db gorm.DB
}

func NewRentalRepository(db gorm.DB) domain.RentalRepository {
	return &rentalRepository{db: db}
}

func (rentalRepo rentalRepository) CreateNewRental(rentalPayload *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := rentalRepo.db.WithContext(ctx).Create(&rentalPayload); err != nil {
		return err.Error
	}

	return rentalRepo.db.Save(&rentalPayload).Error
}

func (rentalRepo rentalRepository) UpdateRental(rental *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var existingTools []domain.RentTools
	var updatedRentTools = rental.RentTools

	if err := rentalRepo.db.WithContext(ctx).Model(&rental).Association("RentTools").Find(&existingTools); err != nil {
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

	transactionError := internal.WithTransaction(ctx, &rentalRepo.db, func(tx *gorm.DB) error {
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

	return rentalRepo.db.WithContext(ctx).Save(rental).Error
}

func (rentalRepo rentalRepository) DeleteRental(rentalId uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var user domain.User
	user.ID = rentalId

	return rentalRepo.db.WithContext(ctx).Select("RentTools").Delete(&user).Error
}

func (rentalRepo rentalRepository) CompleteRental(rentalId uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var user domain.User
	user.ID = rentalId

	return rentalRepo.db.WithContext(ctx).Model(&user).Update("active", false).Error
}

func (rentalRepo rentalRepository) GetRentalReport(page int, pageSize int, queryTerm string) (domain.RentalReport, domain.MetaModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var totalCompletedRent int64
	var totalCreatedRent int64
	var todayRents []domain.User
	var meta = domain.MetaModel{Page: page}

	totalCompletedRentResult := rentalRepo.db.WithContext(ctx).Model(&domain.User{}).Where("active = ?", false).Count(&totalCompletedRent)

	if totalCompletedRentResult.Error != nil {
		return domain.RentalReport{}, meta, totalCompletedRentResult.Error
	}

	totalCreatedRentResult := rentalRepo.db.WithContext(ctx).Model(&domain.User{}).Where("active = ?", true).Count(&totalCreatedRent)

	if totalCreatedRentResult.Error != nil {
		return domain.RentalReport{}, meta, totalCreatedRentResult.Error
	}

	rentsTotal, err := rentalRepo.getTodayRents(&todayRents, page, pageSize, queryTerm)

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

func (rentalRepo rentalRepository) GetRentals(page int, pageSize int, queryTerm string) ([]domain.User, domain.MetaModel, error) {
	var rentals []domain.User
	var count int64
	var metaData domain.MetaModel

	if countResult := rentalRepo.db.Model(&rentals).Count(&count); countResult.Error != nil {
		return nil, metaData, countResult.Error
	}

	result := rentalRepo.db.Scopes(Paginate(page, pageSize)).
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

func (rentalRepo rentalRepository) getTodayRents(todayRents *[]domain.User, page int, pageSize int, queryTerm string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var total int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour).Format("02-01-2006 15:04")
	formatStartOfdDay := startOfDay.Format("02-01-2006 15:04")

	todayRentsTotalResult := rentalRepo.db.WithContext(ctx).Model(domain.User{}).Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).Count(&total)

	if todayRentsTotalResult.Error != nil {
		return 0, todayRentsTotalResult.Error
	}

	todayRentsResult := rentalRepo.db.WithContext(ctx).Scopes(Paginate(page, pageSize)).
		Preload("RentTools").
		Where("date >= ? AND date < ?", formatStartOfdDay, endOfDay).
		Where("full_name ILIKE ? OR phones LIKE ?", "%"+queryTerm+"%", "%"+queryTerm+"%").
		Find(&todayRents)

	if todayRentsResult.Error != nil {
		return 0, todayRentsResult.Error
	}

	return total, nil
}
