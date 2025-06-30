package helper

import (
	"inventaris/models"
	"inventaris/web"
)

func ToPesananResponse(pesanan models.Pesanan, produk models.Produk) web.PesananResponse{
	return web.PesananResponse{
		ID: pesanan.ID,
		ProdukID: pesanan.ProdukID,
		Produk: web.ProdukInfo{
			Nama: produk.Nama,
			Deskripsi: produk.Deskripsi,
			Harga: produk.Harga,
			Kategori: produk.Kategori,
			Gambar: produk.Gambar,
		},
		Jumlah: pesanan.Jumlah,
		TanggalPesanan: pesanan.TanggalPesanan,
	}
}