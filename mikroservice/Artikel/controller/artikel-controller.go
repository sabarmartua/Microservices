package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Artikel/dto"
	"github.com/sabarmartua/Artikel/helper"
	"github.com/sabarmartua/Artikel/service"
)

type ArtikelController interface {
	GetArtikels(context *gin.Context)
	GetArtikel(context *gin.Context)
	CreateArtikel(context *gin.Context)
	UpdateArtikel(context *gin.Context)
	DeleteArtikel(context *gin.Context)
}

type artikelController struct {
	ArtikelService service.ArtikelService
}

func NewArtikelController(artikelService service.ArtikelService) ArtikelController {
	return &artikelController{
		ArtikelService: artikelService,
	}
}

func (ac *artikelController) GetArtikels(context *gin.Context) {
	artikels, err := ac.ArtikelService.GetArtikels()
	if err != nil {
		response := helper.BuildErrorResponse("Failed to fetch articles", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helper.BuildResponse(true, "Articles retrieved successfully", artikels)
	context.JSON(http.StatusOK, response)
}

func (ac *artikelController) GetArtikel(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	artikel, err := ac.ArtikelService.GetArtikel(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to fetch article", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Article retrieved successfully", artikel)
	context.JSON(http.StatusOK, response)
}

func (ac *artikelController) CreateArtikel(context *gin.Context) {
	var artikelDTO dto.NewArtikelDTO
	if err := context.ShouldBindJSON(&artikelDTO); err != nil {
		response := helper.BuildErrorResponse("Invalid article data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	artikel, err := ac.ArtikelService.CreateArtikel(artikelDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to create article", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Article created successfully", artikel)
	context.JSON(http.StatusOK, response)
}

func (ac *artikelController) UpdateArtikel(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	var artikelDTO dto.UpdateArtikelDTO
	if err := context.ShouldBindJSON(&artikelDTO); err != nil {
		response := helper.BuildErrorResponse("Invalid article data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	artikel, err := ac.ArtikelService.UpdateArtikel(id, artikelDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to update article", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Article updated successfully", artikel)
	context.JSON(http.StatusOK, response)
}

func (ac *artikelController) DeleteArtikel(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err = ac.ArtikelService.DeleteArtikel(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to delete article", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Article deleted successfully", nil)
	context.JSON(http.StatusOK, response)
}
