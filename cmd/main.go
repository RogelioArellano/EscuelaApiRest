package main

import (
	"escuelaApiREST/app"
	"escuelaApiREST/env"
	"log"
)

func main() {
	// Inicializar la aplicación
	r := app.Initialize()

	// Iniciar el servidor en el puerto 8080
	port := env.GetEnv("PORT", "8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
