package controller

import (
	"fmt"
	"inventaris/helper"
	"inventaris/service"
	"inventaris/web"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProdukControllerImpl struct {
	ProdukService service.ProdukService
}

func NewProdukControllerImpl(produkService service.ProdukService) *ProdukControllerImpl {
	return &ProdukControllerImpl{
		ProdukService: produkService,
	}
}

func (p ProdukControllerImpl) Create(ctx *gin.Context) {
	produkReq := web.CreateProdukRequest{}
	if err := ctx.ShouldBindJSON(&produkReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	result, err := p.ProdukService.Create(produkReq)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Create Failed", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusCreated, "Created", result)
}

func (p *ProdukControllerImpl) Update(ctx *gin.Context) {
	produkReq := web.UpdateProdukRequest{}
	if err := ctx.ShouldBindJSON(&produkReq); err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk ID", err.Error())
		return
	}
	produkReq.Id = id

	result, err := p.ProdukService.Update(produkReq)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Update Failed", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Updated", result)
}

func (p *ProdukControllerImpl) Delete(ctx *gin.Context) {
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk ID", err.Error())
		return
	}

	if err := p.ProdukService.Delete(id); err != nil {
		helper.ResponseJSON(ctx, http.StatusNotFound, "Produk Not Found", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Deleted", nil)
}

func (p *ProdukControllerImpl) FindById(ctx *gin.Context) {
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk ID", err.Error())
		return
	}

	result, err := p.ProdukService.FindById(id)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusNotFound, "Produk ID Not Found", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Ok", result)
}

func (p *ProdukControllerImpl) FindAll(ctx *gin.Context) {
	produk, err := p.ProdukService.FindAll()
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Ok", produk)

}

func (p *ProdukControllerImpl) UpdateImage(ctx *gin.Context) {

	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk Id", err.Error())
		return
	}

	file, err := ctx.FormFile("gambar")
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid File", err.Error())
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	err = ctx.SaveUploadedFile(file, "uploads/"+filename)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Failed Upload File", err.Error())
		return
	}

	result, err := p.ProdukService.UpdateImage(id, filename)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusInternalServerError, "Failed Save File", err.Error())
		return
	}

	helper.ResponseJSON(ctx, http.StatusOK, "Updated Gambar", result)
}

func (p *ProdukControllerImpl) DownloadGambar(ctx *gin.Context) {
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk Id", err.Error())
		return
	}

	produk, err := p.ProdukService.FindById(id)
	if err != nil {
		helper.ResponseJSON(ctx, http.StatusNotFound, "Produk Id Not Found", nil)
		return
	}

	if produk.Gambar == "" {
		helper.ResponseJSON(ctx, http.StatusNotFound, "Gambar Produk Not Found", nil)
		return
	}

	file := "uploads/" + produk.Gambar
	download := ctx.DefaultQuery("download", "false")

	if download == "true" {
		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", produk.Gambar))
	}
	//ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", produk.Gambar))

	ctx.File(file)
}
