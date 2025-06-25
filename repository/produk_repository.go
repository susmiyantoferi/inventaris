package repository

import "inventaris/models"

type ProdukRepository interface {
	Create(produk models.Produk) models.Produk
	Update(produk models.Produk) models.Produk
	Delete(produkId int)
	FindById(produkId int) (models.Produk, error)
	FindAll() ([]models.Produk, error)
}
