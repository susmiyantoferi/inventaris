package service

import (
	"errors"
	"fmt"
	"inventaris/helper"
	"inventaris/models"
	"inventaris/repository"
	"inventaris/web"
	"time"

	"github.com/go-playground/validator/v10"
)

type PesananServiceImpl struct {
	PesananRepository repository.PesananRepository
	ProdukRepository  repository.ProdukRepository
	Validate          *validator.Validate
}

func NewPesananServiceImpl (
	pesananRepoimpl repository.PesananRepository,
	validate *validator.Validate,
	produkRepository repository.ProdukRepository,
) *PesananServiceImpl {

	return &PesananServiceImpl{
		PesananRepository: pesananRepoimpl,
		Validate:          validate,
		ProdukRepository:  produkRepository,
	}
}

func (p *PesananServiceImpl) Create(req web.CreatePesananRequest) (web.PesananResponse, error) {
	err := p.Validate.Struct(req)
	if err != nil {
		return web.PesananResponse{}, err
	}

	produk, err := p.ProdukRepository.FindById(req.ProdukID)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			return web.PesananResponse{}, repository.ErrorIdNotFound
		}

		return web.PesananResponse{}, err
	}

	t, err := time.Parse("02-01-2006", req.TanggalPesanan)
	if err != nil {
		return web.PesananResponse{}, fmt.Errorf("invalid date format %w", err)
	}

	pesanan := models.Pesanan{
		ProdukID:       produk.Id,
		Jumlah:         req.Jumlah,
		TanggalPesanan: t,
	}

	result, err := p.PesananRepository.Create(pesanan)
	if err != nil {
		return web.PesananResponse{}, err
	}

	return helper.ToPesananResponse(result, produk), nil
}


func (p *PesananServiceImpl) Update(produkName string, req web.UpdatePesananRequest) (web.PesananResponse, error) {
	err := p.Validate.Struct(req)
	if err != nil {
		return web.PesananResponse{}, err
	}

	t, err := time.Parse("02-01-2006", req.TanggalPesanan)
	if err != nil {
		return web.PesananResponse{}, fmt.Errorf("invalid date format %w", err)
	}

	pesanan := models.Pesanan{
		Jumlah:         req.Jumlah,
		TanggalPesanan: t,
	}

	result, err := p.PesananRepository.Update(produkName, pesanan)
	if err != nil {
		if errors.Is(err, repository.ErrorNameNotFound) {
			return web.PesananResponse{}, repository.ErrorNameNotFound
		}

		return web.PesananResponse{}, err
	}

	return helper.ToPesananResponse(result, result.Produk), nil
}


func (p *PesananServiceImpl) Delete(pesananId int) error {
	_, err := p.PesananRepository.FindById(pesananId)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			return repository.ErrorIdNotFound
		}

		return err
	}

	err = p.PesananRepository.Delete(pesananId)
	if err != nil {
		return err
	}

	return nil
}


func (p *PesananServiceImpl) FindById(pesananId int) (web.PesananResponse, error) {
	result, err := p.PesananRepository.FindById(pesananId)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			return web.PesananResponse{}, repository.ErrorIdNotFound
		}

		return web.PesananResponse{}, err
	}

	return helper.ToPesananResponse(result, result.Produk), nil
}


func (p *PesananServiceImpl) FindAll() ([]web.PesananResponse, error) {
	result, err := p.PesananRepository.FindAll()
	if err != nil {
		return []web.PesananResponse{}, err
	}

	var responses []web.PesananResponse
	for _, v := range result {
		response := web.PesananResponse{
			ID:       v.ID,
			ProdukID: v.ProdukID,
			Produk: web.ProdukInfo{
				Nama:      v.Produk.Nama,
				Deskripsi: v.Produk.Deskripsi,
				Harga:     v.Produk.Harga,
				Kategori:  v.Produk.Kategori,
				Gambar:    v.Produk.Gambar,
			},
			Jumlah:         v.Jumlah,
			TanggalPesanan: v.TanggalPesanan,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
