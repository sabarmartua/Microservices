package service

import (
	"time"

	"github.com/sabarmartua/Kategori-Artikel/dto"
	"github.com/sabarmartua/Kategori-Artikel/model"
	"github.com/sabarmartua/Kategori-Artikel/repository"

	"github.com/mashingan/smapping"
)

type KategoriArtikelService interface {
	GetKategoriArtikels() []model.KategoriArtikel
	GetKategoriArtikel(id int) model.KategoriArtikel
	CreateKategoriArtikel(dto dto.NewKategoriArtikelDTO) model.KategoriArtikel
	UpdateKategoriArtikel(id int, dto dto.UpdateKategoriArtikelDTO) model.KategoriArtikel
	DeleteKategoriArtikel(id int) model.KategoriArtikel
}

type kategoriArtikelService struct {
	kategoriArtikelRepository repository.KategoriArtikelRepository
}

func NewKategoriArtikelService(repository repository.KategoriArtikelRepository) KategoriArtikelService {
	return &kategoriArtikelService{
		kategoriArtikelRepository: repository,
	}
}

func (kas *kategoriArtikelService) GetKategoriArtikels() []model.KategoriArtikel {
	return kas.kategoriArtikelRepository.GetKategoriArtikels()
}

func (kas *kategoriArtikelService) GetKategoriArtikel(id int) model.KategoriArtikel {
	return kas.kategoriArtikelRepository.GetKategoriArtikel(id)
}

func (kas *kategoriArtikelService) CreateKategoriArtikel(dto dto.NewKategoriArtikelDTO) model.KategoriArtikel {
	kategoriArtikelDTOToModel := model.KategoriArtikel{}

	err := smapping.FillStruct(&kategoriArtikelDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.KategoriArtikel{}
	}

	kategoriArtikelDTOToModel.Nama = dto.Nama

	// Mengisi waktu pembuatan dan pembaruan
	now := time.Now()
	kategoriArtikelDTOToModel.CreatedAt = now
	kategoriArtikelDTOToModel.UpdatedAt = now

	return kas.kategoriArtikelRepository.CreateKategoriArtikel(kategoriArtikelDTOToModel)
}

func (kas *kategoriArtikelService) UpdateKategoriArtikel(id int, dto dto.UpdateKategoriArtikelDTO) model.KategoriArtikel {
	kategoriArtikelDTOToModel := model.KategoriArtikel{}
	err := smapping.FillStruct(&kategoriArtikelDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.KategoriArtikel{}
	}

	// Mengisi waktu pembaruan
	kategoriArtikelDTOToModel.UpdatedAt = time.Now()

	return kas.kategoriArtikelRepository.UpdateKategoriArtikel(id, kategoriArtikelDTOToModel)
}

func (kas *kategoriArtikelService) DeleteKategoriArtikel(id int) model.KategoriArtikel {
	return kas.kategoriArtikelRepository.DeleteKategoriArtikel(id)
}
