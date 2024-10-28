package main

import "escuelaApiREST/app"

func main() {
	// Inicializar la aplicación
	r := app.Initialize()

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
