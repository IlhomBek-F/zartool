package models

import (
	"database/sql"
	"time"
)

type Base struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type MetaModel struct {
	Page    uint `json:"page"`
	Total   uint `json:"total"`
	PerPage uint `json:"per_page"`
}
