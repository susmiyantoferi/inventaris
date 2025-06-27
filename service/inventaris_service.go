package service

import "inventaris/web"

type InventarisService interface {
	Create(r web.CreateInventarisRequest) (web.InventarisResponse, error)
	Delete(inventId uint) error
	FindByName(produkName string) (web.InventarisResponse, error)
	AddStok(produkName string,r web.AddStokRequest) (web.InventarisResponse, error)
	ReduceStok(produkName string, r web.AddStokRequest) (web.InventarisResponse, error)
	FindAll() ([]web.InventarisResponse, error)
}
