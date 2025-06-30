package route

import (
	"inventaris/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	ProdukController controller.ProdukController,
	InventarisController controller.InventarisController,
	PesananController controller.PesananController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/produk", ProdukController.Create)
		api.PUT("/produk/:produkId", ProdukController.Update)
		api.DELETE("/produk/:produkId", ProdukController.Delete)
		api.GET("/produk/:produkId", ProdukController.FindById)
		api.GET("/produk", ProdukController.FindAll)
		api.PUT("/produk/:produkId/gambar", ProdukController.UpdateImage)
		api.GET("/produk/:produkId/gambar", ProdukController.DownloadGambar)

		api.POST("/inventaris", InventarisController.Create)
		api.GET("/inventaris/:produkName", InventarisController.FindByName)
		api.DELETE("/inventaris/:inventId", InventarisController.Delete)
		api.PUT("/inventaris/:produkName/add-stok", InventarisController.AddStok)
		api.PUT("/inventaris/:produkName/reduce-stok", InventarisController.ReduceStok)
		api.GET("/inventaris", InventarisController.FindAll)

		api.POST("/pesanan", PesananController.Create)
		api.PUT("/pesanan/:produkName", PesananController.Update)
		api.DELETE("/pesanan/:pesananId", PesananController.Delete)
		api.GET("/pesanan/:pesananId", PesananController.FindById)
		api.GET("/pesanan", PesananController.FindAll)
	}

	return router
}
