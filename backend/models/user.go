package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint        `gorm:"primaryKey" json:"id"`
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	Pre_payment uint        `json:"pre_payment"`
	Active      bool        `json:"active"`
	RentTools   []RentTools `json:"rent_tools"`
}
