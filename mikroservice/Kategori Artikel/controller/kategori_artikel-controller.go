package controller

import (
	"github.com/sabarmartua/Kategori-Artikel/dto"
	"github.com/sabarmartua/Kategori-Artikel/helper"
	"github.com/sabarmartua/Kategori-Artikel/model"
	"github.com/sabarmartua/Kategori-Artikel/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoriArtikelController interface {
	GetKategoriArtikels(context *gin.Context)
	GetKategoriArtikel(context *gin.Context)
	CreateKategoriArtikel(context *gin.Context)
	UpdateKategoriArtikel(context *gin.Context)
	DeleteKategoriArtikel(context *gin.Context)
}

type kategoriArtikelController struct {
	kategoriArtikelService service.KategoriArtikelService
}

func NewKategoriArtikelController(kategoriArtikelService service.KategoriArtikelService) KategoriArtikelController {
	return &kategoriArtikelController{
		kategoriArtikelService: kategoriArtikelService,
	}
}

func (kac *kategoriArtikelController) GetKategoriArtikels(context *gin.Context) {
	kategoriArtikels := kac.kategoriArtikelService.GetKategoriArtikels()
	res := helper.BuildResponse(true, "Kategori artikels retrieved successfully", kategoriArtikels)
	context.JSON(http.StatusOK, res)
}

func (kac *kategoriArtikelController) GetKategoriArtikel(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid kategori artikel id", "Error", model.KategoriArtikel{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	kategoriArtikel := kac.kategoriArtikelService.GetKategoriArtikel(idInt)
	res := helper.BuildResponse(true, "Kategori artikel retrieved successfully", kategoriArtikel)
	context.JSON(http.StatusOK, res)
}

func (kac *kategoriArtikelController) CreateKategoriArtikel(context *gin.Context) {
	var newKategoriArtikelDTO dto.NewKategoriArtikelDTO
	err := context.ShouldBind(&newKategoriArtikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid kategori artikel data", err.Error(), model.KategoriArtikel{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	newKategoriArtikel := kac.kategoriArtikelService.CreateKategoriArtikel(newKategoriArtikelDTO)
	res := helper.BuildResponse(true, "Kategori artikel created successfully", newKategoriArtikel)
	context.JSON(http.StatusOK, res)
}

func (kac *kategoriArtikelController) UpdateKategoriArtikel(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	var updateKategoriArtikelDTO dto.UpdateKategoriArtikelDTO
	err = context.ShouldBind(&updateKategoriArtikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid kategori artikel data", err.Error(), model.KategoriArtikel{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	updatedKategoriArtikel := kac.kategoriArtikelService.UpdateKategoriArtikel(idInt, updateKategoriArtikelDTO)
	res := helper.BuildResponse(true, "Kategori artikel updated successfully", updatedKategoriArtikel)
	context.JSON(http.StatusOK, res)
}

func (kac *kategoriArtikelController) DeleteKategoriArtikel(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		res := helper.BuildErrorResponse("Invalid kategori artikel id", err.Error(), model.KategoriArtikel{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	kategoriArtikelDeleted := kac.kategoriArtikelService.DeleteKategoriArtikel(idInt)
	res := helper.BuildResponse(true, "Kategori artikel deleted successfully", kategoriArtikelDeleted)
	context.JSON(http.StatusOK, res)
}
