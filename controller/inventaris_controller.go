package controller

import (

	"github.com/gin-gonic/gin"
)

type InventarisController interface{
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByName(ctx *gin.Context)
	AddStok(ctx *gin.Context)
	ReduceStok(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}