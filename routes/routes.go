package routes

import (
	controllers "escuelaApiREST/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, estudiante *controllers.EstudianteController) {
	router.GET("/estudiantes", estudiante.GetEstudiantes)
	router.POST("/estudiantes", estudiante.CreateEstudiante)
	router.PUT("/estudiantes/:id", estudiante.UpdateEstudiante)
	router.PATCH("/estudiantes/:id", estudiante.PatchEstudiante)
	router.DELETE("/estudiantes/:id", estudiante.DeleteEstudiante)
}
