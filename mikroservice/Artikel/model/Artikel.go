package model

import (
	"time"
)

type Artikel struct {
	ID            uint64    `gorm:"primary_key:auto_increment" json:"id"`
    Nama          string    `gorm:"type:varchar(255)" json:"nama"`
    Konten        string    `gorm:"type:text" json:"konten"`
    KategoriID    uint64    `json:"kategori_id"`
    Gambar        string    `gorm:"type:varchar(255)" json:"gambar"`
    CreatedAt     time.Time `gorm:"type:datetime" json:"created_at"`
    UpdatedAt     time.Time `gorm:"type:datetime" json:"updated_at"`
}