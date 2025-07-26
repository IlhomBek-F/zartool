package usecase

import "zartool/domain"

type createOwnerUsecase struct {
	ownerRepository domain.OwnerRepository
}

func NewCreateOwnerusecase(or domain.OwnerRepository) domain.CreateOwnerUsecase {
	return createOwnerUsecase{ownerRepository: or}
}

func (cu createOwnerUsecase) GetOwnerByLogin(login string) (domain.Owner, error) {
	return cu.ownerRepository.GetOwnerByLogin(login)
}

func (cu createOwnerUsecase) CreateOwner(owner domain.Owner) error {
	return cu.ownerRepository.CreateOwner(owner)
}
