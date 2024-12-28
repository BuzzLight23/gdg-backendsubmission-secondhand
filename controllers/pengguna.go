package controllers

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePengguna - Tambah pengguna baru
func CreatePengguna(c *gin.Context) {
	var pengguna models.Pengguna
	if err := c.ShouldBindJSON(&pengguna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pengguna})
}

// GetPengguna - Ambil semua pengguna
func GetPengguna(c *gin.Context) {
	var pengguna []models.Pengguna
	if err := config.DB.Find(&pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pengguna})
}

// UpdatePengguna - Update data pengguna berdasarkan ID
func UpdatePengguna(c *gin.Context) {
	id := c.Param("id")
	var pengguna models.Pengguna

	if err := config.DB.First(&pengguna, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&pengguna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pengguna})
}

// DeletePengguna - Hapus pengguna berdasarkan ID
func DeletePengguna(c *gin.Context) {
	id := c.Param("id")
	var pengguna models.Pengguna

	if err := config.DB.First(&pengguna, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Pengguna berhasil dihapus"})
}
