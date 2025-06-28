package controller

import "github.com/gin-gonic/gin"

type ProdukController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	UpdateImage(ctx *gin.Context)
	DownloadGambar(ctx *gin.Context)
}
