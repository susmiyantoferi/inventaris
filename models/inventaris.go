package models

type Inventaris struct{
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ProdukID uint `gorm:"notnull" json:"produkId"`
	Produk Produk `gorm:"forignKey;references:Id" json:"produk"`
	Jumlah int `gorm:"notnull" json:"jumlah"`
	Lokasi string `gorm:"notnull" json:"lokasi"`
}