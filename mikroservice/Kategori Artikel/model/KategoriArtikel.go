package model

import (
	"time"
)

type KategoriArtikel struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string    `gorm:"type:varchar(255)" json:"nama"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
