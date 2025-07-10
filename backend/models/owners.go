package models

import "gorm.io/gorm"

type Owners struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
}
