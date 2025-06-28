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
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Request",
			Data:   err.Error(),
		})
		return
	}

	result, err := p.ProdukService.Create(produkReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Create Failed",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   result,
	})
}

func (p *ProdukControllerImpl) Update(ctx *gin.Context) {
	produkReq := web.UpdateProdukRequest{}
	if err := ctx.ShouldBindJSON(&produkReq); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Request",
			Data:   err.Error(),
		})
		return
	}

	produkId := ctx.Param("produkId")
	id, errs := strconv.Atoi(produkId)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Produk ID",
			Data:   errs.Error(),
		})
		return
	}
	produkReq.Id = id

	result, err := p.ProdukService.Update(produkReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Update Failed",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   result,
	})
}

func (p *ProdukControllerImpl) Delete(ctx *gin.Context) {
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Produk ID",
			Data:   err.Error(),
		})
		return
	}

	if errs := p.ProdukService.Delete(id); errs != nil {
		ctx.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Produk Not Found",
			Data:   errs.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Deleted",
		Data:   nil,
	})

}

func (p *ProdukControllerImpl) FindById(ctx *gin.Context) {
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid Produk ID",
			Data:   err.Error(),
		})
		return
	}

	result, errs := p.ProdukService.FindById(id)
	if errs != nil {
		ctx.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Produk ID Not Found",
			Data:   errs.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	})
}

func (p *ProdukControllerImpl) FindAll(ctx *gin.Context) {
	produk, err := p.ProdukService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   produk,
	})

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

func(p *ProdukControllerImpl) DownloadGambar(ctx *gin.Context){
	produkId := ctx.Param("produkId")
	id, err := strconv.Atoi(produkId)
	if err != nil{
		helper.ResponseJSON(ctx, http.StatusBadRequest, "Invalid Produk Id", err.Error())
		return
	}

	produk, err := p.ProdukService.FindById(id)
	if err != nil{
		helper.ResponseJSON(ctx, http.StatusNotFound, "Produk Id Not Found", nil)
		return
	}

	if produk.Gambar == ""{
		helper.ResponseJSON(ctx, http.StatusNotFound, "Gambar Produk Not Found", nil)
		return
	}

	file := "uploads/" + produk.Gambar
	download := ctx.DefaultQuery("download", "false")
	
	if download == "true"{
		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", produk.Gambar))
	}
	//ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", produk.Gambar))

	ctx.File(file)
}