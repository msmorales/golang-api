package services

import (
	"fmt"
	"math/rand"
	"time"

	"api/config"
	"api/models"
)

// FillUsers llena la tabla de usuarios con datos generados
func FillUsers(count int) {
	for i := 0; i < count; i++ {
		user := models.User{
			Name:  fmt.Sprintf("User%d", i+1),
			Email: fmt.Sprintf("user%d@example.com", i+1),
		}
		config.DB.Create(&user)
	}
}

// FillVideos llena la tabla de videos con datos generados
func FillVideos(count int) {
	for i := 0; i < count; i++ {
		video := models.Video{
			Title:       fmt.Sprintf("Video Title %d", i+1),
			Content:     fmt.Sprintf("Video Content %d", i+1),
			Description: fmt.Sprintf("Description for video %d", i+1),
			UserID:      uint(rand.Intn(10) + 1), // Asumiendo que hay al menos 10 usuarios
		}
		config.DB.Create(&video)
	}
}

// FillChallenges llena la tabla de desafíos con datos generados
func FillChallenges(count int) {
	for i := 0; i < count; i++ {
		challenge := models.Challenge{
			Title:       fmt.Sprintf("Challenge Title %d", i+1),
			Description: fmt.Sprintf("Description for challenge %d", i+1),
			UserID:      uint(rand.Intn(10) + 1), // Asumiendo que hay al menos 10 usuarios
		}
		config.DB.Create(&challenge)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
