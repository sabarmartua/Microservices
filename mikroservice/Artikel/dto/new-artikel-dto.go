package dto

type NewArtikelDTO struct {
	Nama       string `json:"nama" binding:"required"`
    Konten     string `json:"konten" binding:"required"`
    KategoriID uint64 `json:"kategori_id" binding:"required"`
    Gambar     string `json:"gambar" binding:"required"`
}