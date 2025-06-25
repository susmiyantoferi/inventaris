package repository

import (
	"errors"
	"inventaris/helper"
	"inventaris/models"

	"gorm.io/gorm"
)

type ProdukRepositoryImpl struct {
	DB *gorm.DB
}

func NewProdukRepositoryImpl(db *gorm.DB) *ProdukRepositoryImpl {
	return &ProdukRepositoryImpl{DB: db}
}

func (p *ProdukRepositoryImpl) Create(produk models.Produk) models.Produk {
	result := p.DB.Create(&produk)
	helper.PanicErr(result.Error)
	return produk
}

func (p *ProdukRepositoryImpl) Update(produk models.Produk) models.Produk {
	var data = models.Produk{
		Id:        produk.Id,
		Nama:      produk.Nama,
		Deskripsi: produk.Deskripsi,
		Harga:     produk.Harga,
		Kategori:  produk.Kategori,
	}

	result := p.DB.First(&produk, produk.Id).Updates(data)
	helper.PanicErr(result.Error)
	return produk
}

func (p *ProdukRepositoryImpl) Delete(produkId int) {
	produk := models.Produk{}
	result := p.DB.Delete(&produk, produkId)
	helper.PanicErr(result.Error)
}

func (p *ProdukRepositoryImpl) FindById(produkId int) (models.Produk, error) {
	data := models.Produk{}
	result := p.DB.First(&data, produkId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return data, errors.New("ID Not Found")
		}
		return data, result.Error
	}

	return data, nil
}

func (p *ProdukRepositoryImpl) FindAll() ([]models.Produk, error) {
	var produk []models.Produk
	result := p.DB.Find(&produk)

	if result.Error != nil {
		return nil, result.Error
	}

	return produk, nil
}
