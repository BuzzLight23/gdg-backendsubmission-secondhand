package models

type Review struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id_review"`
	IDProduk   uint   `json:"id_produk"`
	IDPengguna uint   `json:"id_pengguna"`
	Rating     int    `json:"rating"`   // Nilai rating, misalnya dari 1 hingga 5
	Komentar   string `json:"komentar"` // Komentar pengguna
}
