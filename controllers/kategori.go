package controllers

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateKategori - Tambah kategori baru
func CreateKategori(c *gin.Context) {
	var kategori models.Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategori})
}

// GetKategori - Ambil semua kategori
func GetKategori(c *gin.Context) {
	var kategori []models.Kategori
	if err := config.DB.Find(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategori})
}

// UpdateKategori - Update data kategori berdasarkan ID
func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori

	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategori})
}

// DeleteKategori - Hapus kategori berdasarkan ID
func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori

	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Kategori berhasil dihapus"})
}
