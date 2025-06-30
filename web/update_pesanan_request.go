package web

type UpdatePesananRequest struct {
	Jumlah         int    `validate:"required,gt=0" json:"jumlah"`
	TanggalPesanan string `validate:"required,datetime=02-01-2006" json:"tanggal_pesanan"`
}
