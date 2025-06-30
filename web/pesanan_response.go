package web

import "time"

type PesananResponse struct {
	ID             uint       `json:"id"`
	ProdukID       uint       `json:"produk_id"`
	Produk         ProdukInfo `json:"produk"`
	Jumlah         int        `json:"jumlah"`
	TanggalPesanan time.Time  `json:"tanggal_pesanan"`
}
