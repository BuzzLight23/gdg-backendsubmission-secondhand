package controllers

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaksi - Tambah transaksi baru
func CreateTransaksi(c *gin.Context) {
	var transaksi models.Transaksi
	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// GetTransaksi - Ambil semua transaksi
func GetTransaksi(c *gin.Context) {
	var transaksi []models.Transaksi
	if err := config.DB.Find(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// UpdateTransaksi - Update data transaksi berdasarkan ID
func UpdateTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi

	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// DeleteTransaksi - Hapus transaksi berdasarkan ID
func DeleteTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi

	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Transaksi berhasil dihapus"})
}
