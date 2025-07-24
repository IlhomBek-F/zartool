package domain

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
	Page    int   `json:"page"`
	Total   int64 `json:"total"`
	PerPage uint  `json:"per_page"`
}
