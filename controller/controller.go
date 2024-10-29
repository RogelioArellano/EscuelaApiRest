package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"escuelaApiREST/models"
	"escuelaApiREST/repositories"
)

// EstudianteController se utiliza para gestionar las operaciones CRUD de estudiantes usando un repositorio
type EstudianteController struct {
	repo repositories.EstudianteRepository
}

// NewEstudianteController crea una instancia de EstudianteController
func NewEstudianteController(repo repositories.EstudianteRepository) *EstudianteController {
	return &EstudianteController{repo: repo}
}

// GetEstudiantes maneja la solicitud para obtener todos los estudiantes
func (ctrl *EstudianteController) GetEstudiantes(c *gin.Context) {
	estudiantes, err := ctrl.repo.GetEstudiantes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, estudiantes)
}

// CreateEstudiante maneja la solicitud para crear un nuevo estudiante
func (ctrl *EstudianteController) CreateEstudiante(c *gin.Context) {
	// Vincular el JSON de la solicitud al modelo, manejando el error
	var input models.Estudiante
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Llamar al repositorio para insertar el nuevo estudiande en BD
	err := ctrl.repo.CreateEstudiante(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//Mensaje de exito
	c.JSON(http.StatusOK, gin.H{"status": "Estudiante creado"})
}

// UpdateEstudiante maneja la solicitud para actualizar todos los datos del estudiante
func (ctrl *EstudianteController) UpdateEstudiante(c *gin.Context) {
	id := c.Param("id") // Obtener el ID de la URL
	var input models.Estudiante

	// Validar y deserializar el JSON de la solicitud en el modelo Estudiante
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de JSON inválido"})
		return
	}

	// Llamar al repositorio para actualizar el estudiante con el ID proporcionado
	err := ctrl.repo.UpdateEstudiante(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Estudiante actualizado"})
}

// PatchEstudiante maneja la solicitud para actualizar campos específicos de un estudiante
func (ctrl *EstudianteController) PatchEstudiante(c *gin.Context) {
	id := c.Param("id") // Obtener el ID de la URL

	// Leer los campos a actualizar desde el cuerpo de la solicitud
	var fields map[string]interface{}
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de JSON inválido"})
		return
	}

	// Llamar al repositorio para actualizar los campos específicos del estudiante
	err := ctrl.repo.PatchEstudiante(id, fields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Campos del estudiante actualizados"})
}

func (ctrl *EstudianteController) DeleteEstudiante(c *gin.Context) {
	id := c.Param("id") // Obtener el ID de la URL

	// Llamar al repositorio para eliminar el estudiante con el ID proporcionado
	err := ctrl.repo.DeleteEstudiante(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Estudiante eliminado"})
}
