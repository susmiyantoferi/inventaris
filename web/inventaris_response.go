package web

import (
	"github.com/shopspring/decimal"
)

type ProdukInfo struct {
	Nama      string          `json:"nama"`
	Deskripsi string          `json:"deskripsi"`
	Harga     decimal.Decimal `json:"harga"`
	Kategori  string          `json:"kategori"`
}

type InventarisResponse struct {
	ID       uint       `json:"id"`
	ProdukID uint       `json:"produk_id"`
	Produk   ProdukInfo `json:"produk"`
	Jumlah   int        `json:"jumlah"`
	Lokasi   string     `json:"lokasi"`
}
