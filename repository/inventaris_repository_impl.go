package repository

import (
	"errors"
	"fmt"
	"inventaris/models"

	"gorm.io/gorm"
)

type InventarisRepoImpl struct{
	DB *gorm.DB
}

func NewInventarisRepositoryImpl(db *gorm.DB) *InventarisRepoImpl{
	return &InventarisRepoImpl{DB: db}
}

func(i *InventarisRepoImpl) Create(inventaris models.Inventaris) (models.Inventaris, error){
	if err := i.DB.Create(&inventaris).Error; err != nil{
		return models.Inventaris{}, err
	}
	
	return inventaris, nil
}

func(i *InventarisRepoImpl) Delete(inventId uint) error{
	var inv  models.Inventaris
	err := i.DB.Delete(&inv, inventId).Error
	if err != nil {
		return err
	}

	return nil
}

func(i *InventarisRepoImpl) FindByName(produkName string) (models.Inventaris, error){
	inv := models.Inventaris{}

	result := i.DB.Joins("JOIN produks ON produks.id = inventaris.produk_id").
	Where("produks.nama", produkName).
	Preload("Produk").First(&inv)

	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return inv, errors.New("name not found")
		}
		return inv, result.Error
	}

	return inv, nil
}

func(i *InventarisRepoImpl) AddStok(produkName string, invent models.Inventaris) (models.Inventaris, error){
	inv := models.Inventaris{}

	result := i.DB.Joins("JOIN produks ON produks.id = inventaris.produk_id").
	Where("produks.nama", produkName).
	Preload("Produk").First(&inv)

	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return inv, errors.New("name produk not found")
		}
		return inv, result.Error
	}

	inv.Jumlah += invent.Jumlah
	
	update := map[string]interface{}{
		"jumlah" : inv.Jumlah,
	}

	if invent.Lokasi != "" {
		update["lokasi"] = invent.Lokasi
	}

	if err := i.DB.Model(&inv).Updates(update).Error; err != nil{
		return models.Inventaris{}, err
	}

	return inv, nil
}

func(i *InventarisRepoImpl) ReduceStok(produkName string, invent models.Inventaris) (models.Inventaris, error){
	inv := models.Inventaris{}

	result := i.DB.Joins("JOIN produks ON produks.id = inventaris.produk_id").
	Where("produks.nama", produkName).
	Preload("Produk").First(&inv)

	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return inv, errors.New("name produk not found")
		}
		return inv, result.Error
	}

	if invent.Jumlah > inv.Jumlah {
		return models.Inventaris{}, errors.New("not enough stock ")
	}

	inv.Jumlah -= invent.Jumlah
	
	update := map[string]interface{}{
		"jumlah" : inv.Jumlah,
	}

	if invent.Lokasi != "" {
		update["lokasi"] = invent.Lokasi
	}

	if err := i.DB.Model(&inv).Updates(update).Error; err != nil{
		return models.Inventaris{}, err
	}

	return inv, nil
}

func(i *InventarisRepoImpl) FindAll() ([]models.Inventaris, error){
	var inv []models.Inventaris

	result := i.DB.Preload("Produk").Find(&inv)

	if result.Error != nil {
		return nil, result.Error
	}

	return inv, nil
}

func(i *InventarisRepoImpl) FindById(inventId uint) (models.Inventaris, error){
	var inv models.Inventaris
	result := i.DB.First(&inv, inventId).Error

	if result != nil{
		if errors.Is(result, gorm.ErrRecordNotFound) {
			return models.Inventaris{}, fmt.Errorf("id %d not found", inventId)
		}
		return models.Inventaris{}, result
	}

	return inv, nil
}