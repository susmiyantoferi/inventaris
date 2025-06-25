package service

import "inventaris/web"

type ProdukService interface {
	Create(r web.CreateProdukRequest) (web.ProdukResponse, error)
	Update(r web.UpdateProdukRequest) (web.ProdukResponse, error)
	Delete(produkId int) (error)
	FindById(produkId int) (web.ProdukResponse, error)
	FindAll() ([]web.ProdukResponse, error)
}
