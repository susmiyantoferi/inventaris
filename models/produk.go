package models

import "github.com/shopspring/decimal"

type Produk struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama      string          `gorm:"notnull" json:"nama"`
	Deskripsi string          `gorm:"default:null" json:"deskripsi"`
	Harga     decimal.Decimal `gorm:"type:DECIMAL(10,2);notnull" json:"harga"`
	Kategori  string          `gorm:"default:null" json:"kategori"`
}
