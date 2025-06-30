package repository

import "inventaris/models"

type PesananRepository interface {
	Create(pesanan models.Pesanan) (models.Pesanan, error)
	Update(produkName string, pesanan models.Pesanan) (models.Pesanan, error)
	Delete(pesananId int) error
	FindById(pesananId int) (models.Pesanan, error)
	FindAll() ([]models.Pesanan, error)
}
