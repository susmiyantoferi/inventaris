package web

type CreatePesananRequest struct {
	ProdukID       int    `validate:"required" json:"produk_id"`
	Jumlah         int    `validate:"required,gt=0" json:"jumlah"`
	TanggalPesanan string `validate:"required,datetime=02-01-2006" json:"tanggal_pesanan"`
}
