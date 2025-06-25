package controller

import (
	"inventaris/service"
	"inventaris/web"
	"net/http"
	"strconv"

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

	if errs := p.ProdukService.Delete(id); errs != nil{
		ctx.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Produk Not Found",
			Data:   errs.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code: http.StatusOK,
		Status: "Deleted",
		Data: nil,
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
		Code: http.StatusOK,
		Status: "Ok",
		Data: result,
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
		Code: http.StatusOK,
		Status: "Ok",
		Data: produk,
	})

}
