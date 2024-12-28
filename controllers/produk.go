package controllers

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduk - Tambah produk baru
func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Simpan produk ke database
	if err := config.DB.Create(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": produk})
}

// GetProduk - Ambil semua produk
func GetProduk(c *gin.Context) {
	var produk []models.Produk

	// Ambil semua produk dari database
	if err := config.DB.Find(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": produk})
}

// UpdateProduk - Update data produk berdasarkan ID
func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk

	// Cari produk berdasarkan ID
	if err := config.DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Bind data dari request body
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Simpan perubahan ke database
	if err := config.DB.Save(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": produk})
}

// DeleteProduk - Hapus produk berdasarkan ID
func DeleteProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk

	// Cari produk berdasarkan ID
	if err := config.DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Hapus produk dari database
	if err := config.DB.Delete(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Produk berhasil dihapus"})
}
