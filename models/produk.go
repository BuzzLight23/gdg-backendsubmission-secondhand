package models

type Produk struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id_produk"`
	IDPengguna uint    `json:"id_pengguna"`
	IDKategori uint    `json:"id_kategori"`
	Nama       string  `json:"nama"`
	Deskripsi  string  `json:"deskripsi"`
	Harga      float64 `json:"harga"`
	Kondisi    string  `json:"kondisi"` // "baru" atau "bekas"
}
