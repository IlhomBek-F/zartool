package models

import "gorm.io/gorm"

type RentTools struct {
	gorm.Model
	Name     string `json:"name"`
	Size     string `json:"size"`
	Quantity uint   `json:"quantity"`
	UserId   uint
}
