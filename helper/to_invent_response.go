package helper

import (
	"inventaris/models"
	"inventaris/web"
)

func ToInventResponse(inv models.Inventaris, produk models.Produk) web.InventarisResponse {
	return web.InventarisResponse{
		ID:       inv.ID,
		ProdukID: inv.ProdukID,
		Produk: web.ProdukInfo{
			Nama:      produk.Nama,
			Deskripsi: produk.Deskripsi,
			Harga:     produk.Harga,
			Kategori:  produk.Kategori,
			Gambar: produk.Gambar,
		},
		Jumlah: inv.Jumlah,
		Lokasi: inv.Lokasi,
	}
}
