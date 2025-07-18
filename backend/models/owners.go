package models

import (
	"time"
)

type Owners struct {
	Base
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Credential struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	AccessToken string    `json:"access_token"`
}

type OwnerResponse = SuccessResponseWithData[Credential]
