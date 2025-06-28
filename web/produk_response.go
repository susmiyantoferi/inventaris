package web

import "github.com/shopspring/decimal"

type ProdukResponse struct {
	ID        uint            `json:"id"`
	Nama      string          `json:"nama"`
	Deskripsi string          `json:"deskripsi"`
	Harga     decimal.Decimal `json:"harga"`
	Kategori  string          `json:"kategori"`
	Gambar    string          `json:"gambar"`
}
