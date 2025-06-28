package service

import (
	"inventaris/models"
	"inventaris/repository"
	"inventaris/web"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

type ProdukServiceImpl struct {
	ProdukRepository repository.ProdukRepository
	Validate         *validator.Validate
}

func NewProdukServiceImpl(produkRepository repository.ProdukRepository, validate *validator.Validate) *ProdukServiceImpl {
	return &ProdukServiceImpl{
		ProdukRepository: produkRepository,
		Validate:         validate,
	}
}

func (p *ProdukServiceImpl) Create(r web.CreateProdukRequest) (web.ProdukResponse, error) {
	err := p.Validate.Struct(r)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	hargacvt, errs := decimal.NewFromString(r.Harga)
	if errs != nil {
		return web.ProdukResponse{}, errs
	}

	produk := models.Produk{
		Nama:      r.Nama,
		Deskripsi: r.Deskripsi,
		Harga:     hargacvt,
		Kategori:  r.Kategori,
	}

	result := p.ProdukRepository.Create(produk)
	response := web.ProdukResponse{
		ID:        result.Id,
		Nama:      result.Nama,
		Deskripsi: result.Deskripsi,
		Harga:     result.Harga,
		Kategori:  result.Kategori,
	}

	return response, nil
}

func (p *ProdukServiceImpl) Update(r web.UpdateProdukRequest) (web.ProdukResponse, error) {

	err := p.Validate.Struct(r)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	produk, errs := p.ProdukRepository.FindById(r.Id)
	if errs != nil {
		return web.ProdukResponse{}, errs
	}

	hargacvt, _ := decimal.NewFromString(r.Harga)

	produk.Nama = r.Nama
	produk.Deskripsi = r.Deskripsi
	produk.Harga = hargacvt
	produk.Kategori = r.Kategori

	result := p.ProdukRepository.Update(produk)
	response := web.ProdukResponse{
		ID:        result.Id,
		Nama:      result.Nama,
		Deskripsi: result.Deskripsi,
		Harga:     result.Harga,
		Kategori:  result.Kategori,
	}

	return response, nil
}

func (p *ProdukServiceImpl) Delete(produkId int) error {
	_, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		return err
	}

	p.ProdukRepository.Delete(produkId)

	return nil
}

func (p *ProdukServiceImpl) FindById(produkId int) (web.ProdukResponse, error) {
	produk, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	response := web.ProdukResponse{
		ID:        produk.Id,
		Nama:      produk.Nama,
		Deskripsi: produk.Deskripsi,
		Harga:     produk.Harga,
		Kategori:  produk.Kategori,
		Gambar:    produk.Gambar,
	}

	return response, nil
}

func (p *ProdukServiceImpl) FindAll() ([]web.ProdukResponse, error) {
	produk, err := p.ProdukRepository.FindAll()
	if err != nil {
		return []web.ProdukResponse{}, err
	}

	var responses []web.ProdukResponse
	for _, value := range produk {
		response := web.ProdukResponse{
			ID:        value.Id,
			Nama:      value.Nama,
			Deskripsi: value.Deskripsi,
			Harga:     value.Harga,
			Kategori:  value.Kategori,
			Gambar:    value.Gambar,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *ProdukServiceImpl) UpdateImage(produkId int, gambar string) (web.ProdukResponse, error) {
	produk, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	result, err := p.ProdukRepository.UpdateImage(produkId, gambar)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	response := web.ProdukResponse{
		ID:        produk.Id,
		Nama:      produk.Nama,
		Deskripsi: produk.Deskripsi,
		Harga:     produk.Harga,
		Kategori:  produk.Kategori,
		Gambar:    result.Gambar,
	}

	return response, nil
}
