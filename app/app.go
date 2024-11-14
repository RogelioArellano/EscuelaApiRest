package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"escuelaApiREST/config"
	controllers "escuelaApiREST/controller"
	"escuelaApiREST/repositories"
	"escuelaApiREST/routes"
)

// Initialize configura la base de datos, repositorios, controladores y rutas
func Initialize() *gin.Engine {
	// Establecer conexión a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	// Crear instancia del repositorio de estudiantes
	estudianteRepo := repositories.NewEstudianteRepo(db)

	// Crear instancia del controlador de estudiantes, pasando el repositorio
	estudianteController := controllers.NewEstudianteController(estudianteRepo)

	// Configurar el router
	r := gin.Default()

	// Registrar las rutas, pasando el controlador de estudiantes
	routes.RegisterRoutes(r, estudianteController)

	return r
}
