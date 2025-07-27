package usecase

import "zartool/domain"

type rentalUsecase struct {
	rentalRepository domain.RentalRepository
}

func NewRentalUsecase(rr domain.RentalRepository) domain.RentalUsecase {
	return rentalUsecase{
		rentalRepository: rr,
	}
}

func (ru rentalUsecase) CreateNewRental(rentalPayload *domain.User) error {
	return ru.rentalRepository.CreateNewRental(rentalPayload)
}

func (ru rentalUsecase) UpdateRental(rental *domain.User) error {
	return ru.rentalRepository.UpdateRental(rental)
}

func (ru rentalUsecase) DeleteRental(rentalId uint) error {
	return ru.rentalRepository.DeleteRental(rentalId)
}

func (ru rentalUsecase) CompleteRental(rentalId uint) error {
	return ru.rentalRepository.CompleteRental(rentalId)
}

func (ru rentalUsecase) GetRentalReport(page int, pageSize int, queryTerm string) (domain.RentalReport, domain.MetaModel, error) {
	return ru.rentalRepository.GetRentalReport(page, pageSize, queryTerm)
}

func (ru rentalUsecase) GetRentals(page int, pageSize int, queryTerm string) ([]domain.User, domain.MetaModel, error) {
	return ru.rentalRepository.GetRentals(page, pageSize, queryTerm)
}
