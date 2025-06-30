package repository

import (
	"errors"
	"inventaris/models"

	"gorm.io/gorm"
)

type PesananRepositoryImpl struct {
	DB *gorm.DB
}

func NewPesananRepositoryImpl(db *gorm.DB) *PesananRepositoryImpl {
	return &PesananRepositoryImpl{DB: db}
}


var ErrorIdNotFound = errors.New("produk id not found")
var ErrorNameNotFound = errors.New("produk name not found")

func (p *PesananRepositoryImpl) Create(pesanan models.Pesanan) (models.Pesanan, error) {
	if err := p.DB.Create(&pesanan).Error; err != nil {
		return models.Pesanan{}, err
	}

	return pesanan, nil
}


func (p *PesananRepositoryImpl) Update(produkName string, pesanan models.Pesanan) (models.Pesanan, error) {
	var data models.Pesanan

	result := p.DB.Joins("JOIN produks ON produks.id = pesanans.produk_id").
		Where("produks.nama", produkName).Preload("Produk").First(&data)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return data, ErrorNameNotFound
		}
		return data, result.Error
	}

	update := map[string]interface{}{
		"jumlah":          pesanan.Jumlah,
		"tanggal_pesanan": pesanan.TanggalPesanan,
	}

	if err := p.DB.Model(&data).Updates(update).Error; err != nil {
		return models.Pesanan{}, err
	}

	data.Jumlah = pesanan.Jumlah
	data.TanggalPesanan = pesanan.TanggalPesanan

	return data, nil
}


func (p *PesananRepositoryImpl) Delete(pesananId int) error {
	var pesanan models.Pesanan

	result := p.DB.Delete(&pesanan, pesananId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrorIdNotFound
	}

	return nil
}


func (p *PesananRepositoryImpl) FindById(pesananId int) (models.Pesanan, error) {
	var pesanan models.Pesanan

	result := p.DB.Preload("Produk").First(&pesanan, pesananId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Pesanan{}, ErrorIdNotFound
		}

		return models.Pesanan{}, result.Error
	}

	return pesanan, nil
}


func (p *PesananRepositoryImpl) FindAll() ([]models.Pesanan, error) {
	var pesanan []models.Pesanan

	result := p.DB.Preload("Produk").Find(&pesanan)
	if result.Error != nil {
		return nil, result.Error
	}

	return pesanan, nil
}
