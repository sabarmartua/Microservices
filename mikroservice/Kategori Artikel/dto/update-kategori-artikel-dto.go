package dto

type UpdateKategoriArtikelDTO struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
}
