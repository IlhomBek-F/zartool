package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Address     string
	Pre_payment uint
	RentTools   []RentTools
}
