package dto

type NewKategoriArtikelDTO struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
}
