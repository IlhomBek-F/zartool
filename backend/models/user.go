package models

type User struct {
	Base
	Full_name   string      `json:"full_name"`
	Address     string      `json:"address"`
	Pre_payment uint        `json:"pre_payment"`
	Active      bool        `gorm:"default:true" json:"active"`
	Phones      []string    `gorm:"serializer:json" json:"phones"`
	Date        string      `json:"date"`
	RentTools   []RentTools `json:"rent_tools"`
}

type RentalsResponse = SuccessResponse[User]
