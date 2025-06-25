package web


type CreateProdukRequest struct {
	Nama      string `validate:"required,min=1,max=100" json:"nama"`
	Deskripsi string `validate:"required,min=1,max=255" json:"deskripsi"`
	Harga     string `validate:"required,numeric" json:"harga"`
	Kategori  string `validate:"required,min=1,max=200" json:"kategori"`
}
