package repository

import (
	"github.com/sabarmartua/Kategori-Artikel/model"

	"gorm.io/gorm"
)

type KategoriArtikelRepository interface {
	GetKategoriArtikels() []model.KategoriArtikel
	GetKategoriArtikel(id int) model.KategoriArtikel
	CreateKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel
	UpdateKategoriArtikel(id int, kategoriArtikel model.KategoriArtikel) model.KategoriArtikel
	DeleteKategoriArtikel(id int) model.KategoriArtikel
}

type kategoriArtikelConnection struct {
	connection *gorm.DB
}

func NewKategoriArtikelRepository(conn *gorm.DB) KategoriArtikelRepository {
	return &kategoriArtikelConnection{
		connection: conn,
	}

}

func (db *kategoriArtikelConnection) GetKategoriArtikels() []model.KategoriArtikel {
	var kategoriArtikels []model.KategoriArtikel
	db.connection.Find(&kategoriArtikels)
	return kategoriArtikels
}

func (db *kategoriArtikelConnection) GetKategoriArtikel(id int) model.KategoriArtikel {
	var kategoriArtikel model.KategoriArtikel
	db.connection.First(&kategoriArtikel, id)
	return kategoriArtikel
}

func (db *kategoriArtikelConnection) CreateKategoriArtikel(kategoriArtikel model.KategoriArtikel) model.KategoriArtikel {
	db.connection.Create(&kategoriArtikel)
	return kategoriArtikel
}

func (db *kategoriArtikelConnection) UpdateKategoriArtikel(id int, kategoriArtikel model.KategoriArtikel) model.KategoriArtikel {
	var kategoriArtikelUpdate model.KategoriArtikel
	db.connection.First(&kategoriArtikelUpdate, id)

	if kategoriArtikelUpdate.Nama != kategoriArtikel.Nama {
		kategoriArtikelUpdate.Nama = kategoriArtikel.Nama
	}

	db.connection.Save(&kategoriArtikelUpdate)
	return kategoriArtikelUpdate
}

func (db *kategoriArtikelConnection) DeleteKategoriArtikel(id int) model.KategoriArtikel {
	var kategoriArtikel model.KategoriArtikel
	db.connection.First(&kategoriArtikel, id).Delete(&kategoriArtikel)
	return kategoriArtikel
}
