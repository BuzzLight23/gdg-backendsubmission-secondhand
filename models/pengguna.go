package models

type Pengguna struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id_pengguna"`
	Nama     string `json:"nama_pengguna"`
	Email    string `json:"email"`
	Sandi    string `json:"sandi"`
	Sebagai  string `json:"sebagai"` // "pengguna" atau "pembeli"
}
