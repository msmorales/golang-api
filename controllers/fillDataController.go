package controllers

import (
	"net/http"

	"api/services"

	"github.com/gin-gonic/gin"
)

// FillDataRequest representa la estructura del JSON de la solicitud
type FillDataRequest struct {
	Entity string `json:"entity" binding:"required"`
	Count  int    `json:"count" binding:"required"`
}

// FillData llena la tabla especificada con datos generados
func FillData(c *gin.Context) {
	var request FillDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch request.Entity {
	case "users":
		services.FillUsers(request.Count)
	case "videos":
		services.FillVideos(request.Count)
	case "challenges":
		services.FillChallenges(request.Count)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data filled successfully"})
}
