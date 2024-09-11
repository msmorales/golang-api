package main

import (
	"api/config"
	"api/routes"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDB()

	// Configurar las rutas
	router := routes.SetupRouter()

	// Iniciar el servidor
	router.Run(":8080")
}
