package controller

import (
	"errors"
	"inventaris/helper"
	"inventaris/repository"
	"inventaris/service"
	"inventaris/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PesananControllerImpl struct {
	PesananService service.PesananService
}

func NewPesananControllerImpl(pesananService service.PesananService) *PesananControllerImpl {
	return &PesananControllerImpl{PesananService: pesananService}
}

func (p *PesananControllerImpl) Create(ctx *gin.Context) {
	pesananReq := web.CreatePesananRequest{}

	if err := ctx.ShouldBindJSON(&pesananReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	result, err := p.PesananService.Create(pesananReq)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusCreated, "Created", result)
}

func (p *PesananControllerImpl) Update(ctx *gin.Context) {
	pesananReq := web.UpdatePesananRequest{}

	if err := ctx.ShouldBindJSON(&pesananReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	produkName := ctx.Param("produkName")
	if _, err := strconv.Atoi(produkName); err == nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid produkName, must be string not purely numeric", nil)
		return
	}

	result, err := p.PesananService.Update(produkName, pesananReq)
	if err != nil {
		if errors.Is(err, repository.ErrorNameNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Produk Name Not Found", err.Error())
			return
		}

		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Updated", result)

}

func (p *PesananControllerImpl) Delete(ctx *gin.Context) {
	idStr := ctx.Param("pesananId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Pesanan Id Format", err.Error())
		return
	}

	err = p.PesananService.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Pesanan Id Not Found", err.Error())
			return
		}

		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Deleted", nil)

}

func (p *PesananControllerImpl) FindById(ctx *gin.Context) {
	idStr := ctx.Param("pesananId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Pesanan Id Format", err.Error())
		return
	}

	result, err := p.PesananService.FindById(id)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			helper.ResponseJSON(ctx, http.StatusNotFound, "Pesanan Id Not Found", err.Error())
			return
		}

		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Success", result)
}

func (p *PesananControllerImpl) FindAll(ctx *gin.Context) {
	result, err := p.PesananService.FindAll()
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Success", result)
}
