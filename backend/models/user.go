package models

type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Address     string
	Pre_payment uint
	RentTools   []RentTools
}
