package service

import "inventaris/web"

type PesananService interface {
	Create(req web.CreatePesananRequest) (web.PesananResponse, error)
	Update(produkName string, req web.UpdatePesananRequest) (web.PesananResponse, error)
	Delete(pesananId int) error
	FindById(pesananId int) (web.PesananResponse, error)
	FindAll() ([]web.PesananResponse, error)
}
