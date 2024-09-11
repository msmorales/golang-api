package controllers

import (
	"net/http"
	"strconv"

	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
)

// GetChallenges obtiene todos los desafíos con paginación
func GetChallenges(c *gin.Context) {
	var challenges []models.Challenge
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	config.DB.Limit(limit).Offset(offset).Find(&challenges)
	c.JSON(http.StatusOK, challenges)
}

// GetChallenge obtiene un desafío por ID
func GetChallenge(c *gin.Context) {
	id := c.Param("id")
	var challenge models.Challenge
	if err := config.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}
	c.JSON(http.StatusOK, challenge)
}

// CreateChallenge crea un nuevo desafío
func CreateChallenge(c *gin.Context) {
	var challenge models.Challenge
	if err := c.ShouldBindJSON(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&challenge)
	c.JSON(http.StatusCreated, challenge)
}

// UpdateChallenge actualiza un desafío existente
func UpdateChallenge(c *gin.Context) {
	id := c.Param("id")
	var challenge models.Challenge
	if err := config.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}
	if err := c.ShouldBindJSON(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&challenge)
	c.JSON(http.StatusOK, challenge)
}

// DeleteChallenge elimina un desafío por ID
func DeleteChallenge(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Challenge{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
