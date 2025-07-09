package models

type RentTools struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Quantity uint   `json:"quantity"`
	UserId   uint
}
