package models

import (
	"time"
)

type Pesanan struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProdukID       uint      `gorm:"notnull" json:"produk_id"`
	Produk         Produk    `gorm:"forignKey;references:Id" json:"produk"`
	Jumlah         int       `gorm:"notnull" json:"jumlah"`
	TanggalPesanan time.Time `gorm:"notnull" json:"tanggal_pesanan"`
}
