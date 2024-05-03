package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mashingan/smapping"
	"github.com/sabarmartua/Artikel/dto"
	"github.com/sabarmartua/Artikel/model"
	"github.com/sabarmartua/Artikel/repository"
)

type ArtikelService interface {
	GetArtikels() ([]model.Artikel, error)
	GetArtikel(id uint64) (model.Artikel, error)
	CreateArtikel(dto dto.NewArtikelDTO) (model.Artikel, error)
	UpdateArtikel(id uint64, dto dto.UpdateArtikelDTO) (model.Artikel, error)
	DeleteArtikel(id uint64) error
	GetKategoriID(id uint64) (uint64, error)
}

type artikelService struct {
	ArtikelRepository repository.ArtikelRepository
	KategoriAPIURL    string
}

func NewArtikelService(repository repository.ArtikelRepository, kategoriAPIURL string) ArtikelService {
	return &artikelService{
		ArtikelRepository: repository,
		KategoriAPIURL:    kategoriAPIURL,
	}
}

func (as *artikelService) GetArtikels() ([]model.Artikel, error) {
	return as.ArtikelRepository.GetArtikels()
}

func (as *artikelService) GetArtikel(id uint64) (model.Artikel, error) {
	return as.ArtikelRepository.GetArtikel(id)
}

func (as *artikelService) CreateArtikel(dto dto.NewArtikelDTO) (model.Artikel, error) {
	ArtikelDTOToModel := model.Artikel{}

	err := smapping.FillStruct(&ArtikelDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.Artikel{}, err
	}

	// Mengisi waktu pembuatan dan pembaruan
	now := time.Now()
	ArtikelDTOToModel.CreatedAt = now
	ArtikelDTOToModel.UpdatedAt = now

	return as.ArtikelRepository.CreateArtikel(ArtikelDTOToModel)
}

func (as *artikelService) UpdateArtikel(id uint64, dto dto.UpdateArtikelDTO) (model.Artikel, error) {
	ArtikelDTOToModel := model.Artikel{}
	err := smapping.FillStruct(&ArtikelDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.Artikel{}, err
	}

	// Mendapatkan ID kategori berdasarkan ID kategorinya
	idKategori, err := as.GetKategoriID(dto.KategoriID) // Menggunakan ID kategori
	if err != nil {
		return model.Artikel{}, fmt.Errorf("failed to get kategori ID: %v", err)
	}

	ArtikelDTOToModel.KategoriID = idKategori

	// Mengisi waktu pembaruan
	ArtikelDTOToModel.UpdatedAt = time.Now()

	return as.ArtikelRepository.UpdateArtikel(id, ArtikelDTOToModel)
}

func (as *artikelService) DeleteArtikel(id uint64) error {
	return as.ArtikelRepository.DeleteArtikel(id)
}

func (as *artikelService) GetKategoriID(id uint64) (uint64, error) {
	// Panggil API kategori untuk mendapatkan informasi kategori berdasarkan ID
	url := fmt.Sprintf("http://localhost:8080/api/kategori-artikel/kategoriArtikel/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch kategori ID: %s", resp.Status)
	}

	var kategori struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&kategori); err != nil {
		return 0, err
	}

	return kategori.ID, nil
}
