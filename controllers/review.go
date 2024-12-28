package controllers

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview - Tambah review baru
func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReview - Ambil semua review
func GetReview(c *gin.Context) {
	var review []models.Review
	if err := config.DB.Find(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview - Update data review berdasarkan ID
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := config.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview - Hapus review berdasarkan ID
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := config.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Review berhasil dihapus"})
}
