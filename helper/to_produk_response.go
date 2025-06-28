package helper

import (
	"inventaris/models"
	"inventaris/web"
)

func ToProdukResponse(produk models.Produk) web.ProdukResponse{
	return web.ProdukResponse{
		ID: produk.Id,
		Nama: produk.Nama,
		Deskripsi: produk.Deskripsi,
		Harga: produk.Harga,
		Kategori: produk.Kategori,
		Gambar: produk.Gambar,
	}
}