package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Artikel/conn"
	"github.com/sabarmartua/Artikel/controller"
	"github.com/sabarmartua/Artikel/repository"
	"github.com/sabarmartua/Artikel/service"
	"gorm.io/gorm"
)

var (
    db               *gorm.DB                             = conn.SetupDatabaseConnection()
    ArtikelRepository repository.ArtikelRepository = repository.NewArtikelRepository(db)
    ArtikelService    service.ArtikelService       = service.NewArtikelService(ArtikelRepository, os.Getenv("BASE_URL_KATEGORI_ARTIKEL"))
    ArtikelController controller.ArtikelController = controller.NewArtikelController(ArtikelService)
)


func main() {
	defer conn.CloseDatabaseConnection(db)
	router := gin.Default()

	routes := router.Group("/api/artikel")
	{
		routes.GET("/Artikels", ArtikelController.GetArtikels)
		routes.GET("/Artikel/:id", ArtikelController.GetArtikel)
		routes.POST("/create/Artikel", ArtikelController.CreateArtikel)
		routes.PUT("/update/Artikel/:id", ArtikelController.UpdateArtikel)
		routes.DELETE("/delete/Artikel/:id", ArtikelController.DeleteArtikel)
	}

	err := router.Run(os.Getenv("BASE_URL__ARTIKEL"))
	if err != nil {
		panic(err)
	}
}
