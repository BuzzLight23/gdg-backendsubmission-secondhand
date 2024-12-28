package routes

import (
	"gdg_backendsubmission_secondhand/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Endpoint untuk Pengguna
	r.POST("/pengguna", controllers.CreatePengguna)
	r.GET("/pengguna", controllers.GetPengguna)
	r.PUT("/pengguna/:id", controllers.UpdatePengguna)
	r.DELETE("/pengguna/:id", controllers.DeletePengguna)

	// Endpoint untuk Produk
	r.POST("/produk", controllers.CreateProduk)
	r.GET("/produk", controllers.GetProduk)
	r.PUT("/produk/:id", controllers.UpdateProduk)
	r.DELETE("/produk/:id", controllers.DeleteProduk)

	// Endpoint untuk Transaksi
	r.POST("/transaksi", controllers.CreateTransaksi)
	r.GET("/transaksi", controllers.GetTransaksi)
	r.PUT("/transaksi/:id", controllers.UpdateTransaksi)
	r.DELETE("/transaksi/:id", controllers.DeleteTransaksi)

	// Endpoint untuk Kategori
	r.POST("/kategori", controllers.CreateKategori)
	r.GET("/kategori", controllers.GetKategori)
	r.PUT("/kategori/:id", controllers.UpdateKategori)
	r.DELETE("/kategori/:id", controllers.DeleteKategori)

	// Endpoint untuk Review
	r.POST("/review", controllers.CreateReview)
	r.GET("/review", controllers.GetReview)
	r.PUT("/review/:id", controllers.UpdateReview)
	r.DELETE("/review/:id", controllers.DeleteReview)

	return r
}
