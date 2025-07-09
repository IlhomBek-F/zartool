package models

type WarehouseTools struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
}
