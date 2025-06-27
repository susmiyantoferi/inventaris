package web

type CreateInventarisRequest struct{
	ProdukID uint `validate:"required" json:"produk_id"`
	Jumlah int `validate:"required,gt=0" json:"jumlah"`
	Lokasi string `validate:"required,max=100" json:"lokasi"`
}