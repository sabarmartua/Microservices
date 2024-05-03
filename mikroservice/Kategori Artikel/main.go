package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Kategori-Artikel/conn"
	"github.com/sabarmartua/Kategori-Artikel/controller"
	"github.com/sabarmartua/Kategori-Artikel/repository"
	"github.com/sabarmartua/Kategori-Artikel/service"
	"gorm.io/gorm"
)

var (
	db                        *gorm.DB                             = conn.SetupDatabaseConnection()
	kategoriArtikelRepository repository.KategoriArtikelRepository = repository.NewKategoriArtikelRepository(db)
	kategoriArtikelService    service.KategoriArtikelService       = service.NewKategoriArtikelService(kategoriArtikelRepository)
	KategoriArtikelController controller.KategoriArtikelController = controller.NewKategoriArtikelController(kategoriArtikelService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	router := gin.Default()

	routes := router.Group("/api/kategori-artikel")
	{
		routes.GET("/kategoriArtikels", KategoriArtikelController.GetKategoriArtikels)
		routes.GET("/kategoriArtikel/:id", KategoriArtikelController.GetKategoriArtikel)
		routes.POST("/create/kategoriArtikel", KategoriArtikelController.CreateKategoriArtikel)
		routes.PUT("/update/kategoriArtikel/:id", KategoriArtikelController.UpdateKategoriArtikel)
		routes.DELETE("/delete/kategoriArtikel/:id", KategoriArtikelController.DeleteKategoriArtikel)
	}

	err := router.Run(os.Getenv("BASE_URL_KATEGORI_ARTIKEL"))
	if err != nil {
		panic(err)
	}
}
