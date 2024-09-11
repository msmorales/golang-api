package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Rutas para Users
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Rutas para Videos
	router.GET("/videos", controllers.GetVideos)
	router.GET("/videos/:id", controllers.GetVideo)
	router.POST("/videos", controllers.CreateVideo)
	router.PUT("/videos/:id", controllers.UpdateVideo)
	router.DELETE("/videos/:id", controllers.DeleteVideo)

	// Rutas para Challenges
	router.GET("/challenges", controllers.GetChallenges)
	router.GET("/challenges/:id", controllers.GetChallenge)
	router.POST("/challenges", controllers.CreateChallenge)
	router.PUT("/challenges/:id", controllers.UpdateChallenge)
	router.DELETE("/challenges/:id", controllers.DeleteChallenge)

	// Ruta para llenar datos
	router.POST("/fill", controllers.FillData)

	return router
}
