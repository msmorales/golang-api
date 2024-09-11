package controllers

import (
	"net/http"
	"strconv"

	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
)

// GetVideos obtiene todos los videos con paginación
func GetVideos(c *gin.Context) {
	var videos []models.Video
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	config.DB.Limit(limit).Offset(offset).Find(&videos)
	c.JSON(http.StatusOK, videos)
}

// GetVideo obtiene un video por ID
func GetVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := config.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	c.JSON(http.StatusOK, video)
}

// CreateVideo crea un nuevo video
func CreateVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&video)
	c.JSON(http.StatusCreated, video)
}

// UpdateVideo actualiza un video existente
func UpdateVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := config.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&video)
	c.JSON(http.StatusOK, video)
}

// DeleteVideo elimina un video por ID
func DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Video{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
