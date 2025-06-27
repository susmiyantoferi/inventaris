package route

import (
	"inventaris/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	ProdukController controller.ProdukController,
	InventarisController controller.InventarisController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/produk", ProdukController.Create)
		api.PUT("/produk/:produkId", ProdukController.Update)
		api.DELETE("/produk/:produkId", ProdukController.Delete)
		api.GET("/produk/:produkId", ProdukController.FindById)
		api.GET("/produk", ProdukController.FindAll)

		api.POST("/inventaris", InventarisController.Create)
		api.GET("/inventaris/:produkName", InventarisController.FindByName)
		api.DELETE("/inventaris/:inventId", InventarisController.Delete)
		api.PUT("/inventaris/:produkName/add-stok", InventarisController.AddStok)
		api.PUT("/inventaris/:produkName/reduce-stok", InventarisController.ReduceStok)
		api.GET("/inventaris", InventarisController.FindAll)
	}

	return router
}
