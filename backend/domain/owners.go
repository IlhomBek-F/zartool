package domain

import (
	"time"
)

type OwnerRepository interface {
	GetOwnerByLogin(login string) (Owner, error)
	CreateOwner(owner Owner) error
}

type LoginOwnerUsecase interface {
	GetOwnerByLogin(login string) (Owner, error)
	GeneretaAccessToken(user Owner, secret string, expiry int) (token string, err error)
}

type CreateOwnerUsecase interface {
	GetOwnerByLogin(login string) (Owner, error)
	CreateOwner(owner Owner) error
}

type (
	OwnerPayload struct {
		Login    string `json:"login" validate:"required,max=100"`
		Password string `json:"password" validate:"required,min=3"`
	}

	Owner struct {
		Base
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	Credential struct {
		ID          uint      `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		AccessToken string    `json:"access_token"`
	}

	OwnerResponse = SuccessResponseWithData[Credential]
)
