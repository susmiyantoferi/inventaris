package service

import (
	"errors"
	"fmt"
	"inventaris/helper"
	"inventaris/models"
	"inventaris/repository"
	"inventaris/web"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type InventarisServImpl struct {
	InventarisRepo repository.InventarisRepository
	Validate       *validator.Validate
	DB             *gorm.DB
}

func NewInventarisServImpl(inventarisRepo repository.InventarisRepository, validate *validator.Validate, db *gorm.DB) *InventarisServImpl {
	return &InventarisServImpl{
		InventarisRepo: inventarisRepo,
		Validate:       validate,
		DB:             db,
	}
}

func (i *InventarisServImpl) Create(r web.CreateInventarisRequest) (web.InventarisResponse, error) {
	err := i.Validate.Struct(r)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	var produk models.Produk
	err = i.DB.First(&produk, r.ProdukID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.InventarisResponse{}, fmt.Errorf("id %d not found in produk", r.ProdukID)
		}
	}

	inv := models.Inventaris{
		ProdukID: produk.Id,
		Jumlah:   r.Jumlah,
		Lokasi:   r.Lokasi,
	}

	result, er := i.InventarisRepo.Create(inv)
	if er != nil {
		return web.InventarisResponse{}, er
	}

	return helper.ToInventResponse(result, produk), nil

}

func (i *InventarisServImpl) Delete(inventId uint) error {
	result, err := i.InventarisRepo.FindById(inventId)
	if err != nil {
		return err
	}

	err = i.InventarisRepo.Delete(result.ID)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventarisServImpl) FindByName(produkName string) (web.InventarisResponse, error) {
	inv, err := i.InventarisRepo.FindByName(produkName)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	produk := inv.Produk
	return helper.ToInventResponse(inv, produk), nil
}

func (i *InventarisServImpl) AddStok(produkName string, r web.AddStokRequest) (web.InventarisResponse, error) {
	err := i.Validate.Struct(r)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	inv := models.Inventaris{
		Jumlah: r.Jumlah,
	}

	if r.Lokasi != nil {
		inv.Lokasi = *r.Lokasi
	}

	result, err := i.InventarisRepo.AddStok(produkName, inv)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	return helper.ToInventResponse(result, result.Produk), nil
}

func (i *InventarisServImpl) ReduceStok(produkName string, r web.AddStokRequest) (web.InventarisResponse, error) {
	err := i.Validate.Struct(r)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	inv := models.Inventaris{
		Jumlah: r.Jumlah,
	}

	if r.Lokasi != nil {
		inv.Lokasi = *r.Lokasi
	}

	result, err := i.InventarisRepo.ReduceStok(produkName, inv)
	if err != nil {
		return web.InventarisResponse{}, err
	}

	return helper.ToInventResponse(result, result.Produk), nil
}

func (i *InventarisServImpl) FindAll() ([]web.InventarisResponse, error) {
	inv, err := i.InventarisRepo.FindAll()
	if err != nil {
		return []web.InventarisResponse{}, err
	}

	var responses []web.InventarisResponse
	for _, v := range inv {
		response := web.InventarisResponse{
			ID: v.ID,
			ProdukID: v.ProdukID,
			Produk: web.ProdukInfo{
				Nama: v.Produk.Nama,
				Deskripsi: v.Produk.Kategori,
				Harga: v.Produk.Harga,
				Kategori: v.Produk.Kategori,
				Gambar: v.Produk.Gambar,
			},
			Jumlah: v.Jumlah,
			Lokasi: v.Lokasi,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
