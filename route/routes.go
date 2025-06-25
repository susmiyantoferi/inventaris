package route

import (
	"inventaris/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(ProdukController controller.ProdukController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/produk", ProdukController.Create)
		api.PUT("/produk/:produkId", ProdukController.Update)
		api.DELETE("/produk/:produkId", ProdukController.Delete)
		api.GET("/produk/:produkId", ProdukController.FindById)
		api.GET("/produk", ProdukController.FindAll)
	}

	return router
}
