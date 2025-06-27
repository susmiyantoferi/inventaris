package repository

import "inventaris/models"

type InventarisRepository interface{
	Create(inventaris models.Inventaris) (models.Inventaris, error)
	Delete(inventId uint) error
	FindByName(produkName string) (models.Inventaris, error)
	AddStok(produkName string, invent models.Inventaris) (models.Inventaris, error)
	ReduceStok(produkName string, invent models.Inventaris) (models.Inventaris, error)
	FindAll() ([]models.Inventaris, error)
	FindById(inventId uint) (models.Inventaris, error)
}