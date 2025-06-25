package web

type UpdateProdukRequest struct {
	Id        int    `validate:"required" json:"id"`
	Nama      string `validate:"required,max=100" json:"nama"`
	Deskripsi string `validate:"required,max=255" json:"deskripsi"`
	Harga     string `validate:"required,numeric" json:"harga"`
	Kategori  string `validate:"required,max=200" json:"kategori"`
}
