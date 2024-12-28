package models

type Kategori struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id_kategori"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}
