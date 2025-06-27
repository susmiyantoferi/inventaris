package controller

import (
	"errors"
	"inventaris/helper"
	"inventaris/service"
	"inventaris/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventControllerImpl struct {
	InventService service.InventarisService
}

func NewInventControllerImpl(inventService service.InventarisService) *InventControllerImpl {
	return &InventControllerImpl{InventService: inventService}
}

func (i *InventControllerImpl) Create(ctx *gin.Context) {
	invReq := web.CreateInventarisRequest{}

	if err := ctx.ShouldBindBodyWithJSON(&invReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	result, err := i.InventService.Create(invReq)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Ok", result)
}

func (i *InventControllerImpl) Delete(ctx *gin.Context) {
	inventId := ctx.Param("inventId")
	id, err := strconv.ParseUint(inventId, 10, 64)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Inventaris ID", err.Error())
		return
	}

	if err := i.InventService.Delete(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Inventaris Id Not Found", nil)
			return
		}

		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Deleted", nil)
}

func (i *InventControllerImpl) FindByName(ctx *gin.Context) {
	produkName := ctx.Param("produkName")
	inv, err := i.InventService.FindByName(produkName)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusNotFound, "Record Not Found", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Ok", inv)
}

func (i *InventControllerImpl) AddStok(ctx *gin.Context) {
	invReq := web.AddStokRequest{}
	if err := ctx.ShouldBindJSON(&invReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	namaProduk := ctx.Param("produkName")

	result, err := i.InventService.AddStok(namaProduk, invReq)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Inventaris Name Not found", nil)
			return
		}
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Stock Added", result)
}

func (i *InventControllerImpl) ReduceStok(ctx *gin.Context) {
	invReq := web.AddStokRequest{}
	if err := ctx.ShouldBindJSON(&invReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	namaProduk := ctx.Param("produkName")
	result, err := i.InventService.ReduceStok(namaProduk, invReq)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Inventaris Name Not found", nil)
			return
		}

		if err.Error() == "not enough stock" {
			helper.ResponseJSON(ctx, http.StatusBadRequest, "Data Jumlah Over Stock", nil)
			return
		}

		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Stock Reduce", result)
}

func (i *InventControllerImpl) FindAll(ctx *gin.Context) {
	result, err := i.InventService.FindAll()
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Ok", result)
}
