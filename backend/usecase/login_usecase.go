package usecase

import (
	"zartool/domain"
	"zartool/internal"
)

type loginUsecase struct {
	ownerRepository domain.OwnerRepository
}

func NewLoginUsecase(or domain.OwnerRepository) domain.LoginOwnerUsecase {
	return loginUsecase{
		ownerRepository: or,
	}
}

func (lu loginUsecase) GetOwnerByLogin(login string) (domain.Owner, error) {
	return lu.ownerRepository.GetOwnerByLogin(login)
}

func (lu loginUsecase) GeneretaAccessToken(user domain.Owner, secret string, expiry int) (token string, err error) {
	return internal.GeneretaAccessToken(user, secret, expiry)
}
