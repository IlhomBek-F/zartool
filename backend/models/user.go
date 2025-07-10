package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	Pre_payment uint        `json:"pre_payment"`
	RentTools   []RentTools `json:"rent_tools"`
}
