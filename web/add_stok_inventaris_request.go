package web


type AddStokRequest struct{
	Jumlah   int        `validate:"required" json:"jumlah"`
	Lokasi   *string     `json:"lokasi,omitempty"`
}