package models

type Transaksi struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id_transaksi"`
	IDProduk  uint   `json:"id_produk"`
	IDPenjual uint   `json:"id_penjual"`
	IDPembeli uint   `json:"id_pembeli"`
	Status    string `json:"status"` // "ditunda", "selesai", atau "batal"
}
