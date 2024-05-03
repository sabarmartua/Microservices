package repository

import (
    "github.com/sabarmartua/Artikel/model"
    "gorm.io/gorm"
)

type ArtikelRepository interface {
    GetArtikels() ([]model.Artikel, error)
    GetArtikel(id uint64) (model.Artikel, error)
    CreateArtikel(artikel model.Artikel) (model.Artikel, error)
    UpdateArtikel(id uint64, artikel model.Artikel) (model.Artikel, error)
    DeleteArtikel(id uint64) error
}

type ArtikelRepo struct {
    connection *gorm.DB
}

func NewArtikelRepository(conn *gorm.DB) ArtikelRepository {
    return &ArtikelRepo{
        connection: conn,
    }
}

func (repo *ArtikelRepo) GetArtikels() ([]model.Artikel, error) {
    var artikels []model.Artikel
    if err := repo.connection.Find(&artikels).Error; err != nil {
        return nil, err
    }
    return artikels, nil
}

func (repo *ArtikelRepo) GetArtikel(id uint64) (model.Artikel, error) {
    var artikel model.Artikel
    if err := repo.connection.First(&artikel, id).Error; err != nil {
        return model.Artikel{}, err
    }
    return artikel, nil
}

func (repo *ArtikelRepo) CreateArtikel(artikel model.Artikel) (model.Artikel, error) {
    if err := repo.connection.Create(&artikel).Error; err != nil {
        return model.Artikel{}, err
    }
    return artikel, nil
}

func (repo *ArtikelRepo) UpdateArtikel(id uint64, artikel model.Artikel) (model.Artikel, error) {
    if err := repo.connection.Model(&model.Artikel{}).Where("id = ?", id).Updates(artikel).Error; err != nil {
        return model.Artikel{}, err
    }
    return artikel, nil
}

func (repo *ArtikelRepo) DeleteArtikel(id uint64) error {
    if err := repo.connection.Delete(&model.Artikel{}, id).Error; err != nil {
        return err
    }
    return nil
}
