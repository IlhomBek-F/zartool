package models

import "gorm.io/gorm"

type WarehouseTools struct {
	gorm.Model
	Name string `json:"name"`
	Size string `json:"size"`
}
