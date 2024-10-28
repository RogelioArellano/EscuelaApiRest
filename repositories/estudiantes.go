package repositories

import (
	"database/sql"

	"escuelaApiREST/models"
	"fmt"
)

// Definir la interfaz para el repositorio de estudiantes
type EstudianteRepository interface {
	GetEstudiantes() ([]models.Estudiante, error)
	CreateEstudiante(estudiante models.Estudiante) error
	UpdateEstudiante(id string, estudiante models.Estudiante) error
	PatchEstudiante(id string, fields map[string]interface{}) error
	DeleteEstudiante(id string) error
}

// Estructura que implementa la interfaz EstudianteRepository
type estudianteRepo struct {
	db *sql.DB
}

// Constructor para el repositorio de estudiantes
func NewEstudianteRepo(db *sql.DB) EstudianteRepository {
	return &estudianteRepo{db: db}
}

// GetEstudiantes devuelve todos los registros de estudiantes que se encuentren en la base de datos
func (r *estudianteRepo) GetEstudiantes() ([]models.Estudiante, error) {
	rows, err := r.db.Query("SELECT idEstudiante, nombre, direccion, email, telefono, altaLocal, altaSep FROM estudiantes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estudiantes []models.Estudiante
	for rows.Next() {
		var estudiante models.Estudiante
		err := rows.Scan(&estudiante.IDEstudiante, &estudiante.Nombre, &estudiante.Direccion, &estudiante.Email, &estudiante.Telefono, &estudiante.AltaLocal, &estudiante.AltaSep)
		if err != nil {
			return nil, err
		}
		estudiantes = append(estudiantes, estudiante)
	}
	return estudiantes, nil
}

func (r *estudianteRepo) CreateEstudiante(estudiante models.Estudiante) error {
	_, err := r.db.Exec("INSERT INTO estudiantes (nombre, direccion, email, telefono, altaLocal, altaSep) VALUES ($1, $2, $3, $4, $5, $6)",
		estudiante.Nombre, estudiante.Direccion, estudiante.Email, estudiante.Telefono, estudiante.AltaLocal, estudiante.AltaSep)
	return err
}

func (r *estudianteRepo) UpdateEstudiante(id string, estudiante models.Estudiante) error {
	_, err := r.db.Exec(
		"UPDATE estudiantes SET nombre=$1, direccion=$2, email=$3, telefono=$4, altaLocal=$5, altaSep=$6 WHERE idEstudiante=$7",
		estudiante.Nombre, estudiante.Direccion, estudiante.Email, estudiante.Telefono, estudiante.AltaLocal, estudiante.AltaSep, id,
	)
	return err
}

// Función para actualizar campos específicos de un estudiante
func (r *estudianteRepo) PatchEstudiante(id string, fields map[string]interface{}) error {
	// Inicializa la consulta SQL para actualizar el estudiante
	query := "UPDATE estudiantes SET "

	// Crea una lista para almacenar los valores que se van a reemplazar en la consulta SQL
	values := []interface{}{}

	// Contador para los parámetros de la consulta
	i := 1

	// Recorre el mapa "fields", que contiene los campos y valores a actualizar
	for field, value := range fields {
		// Agrega cada campo y un marcador de posición ($i) a la consulta
		// Ejemplo: "nombre = $1, direccion = $2, "
		query += fmt.Sprintf("%s = $%d, ", field, i)

		// Agrega el valor al slice "values" en el mismo orden de los campos
		values = append(values, value)

		// Incrementa el contador para el siguiente marcador de posición
		i++
	}

	// Elimina la última coma y espacio de la consulta para que no genere un error de SQL
	// Esto transforma "UPDATE estudiantes SET nombre = $1, direccion = $2, " a "UPDATE estudiantes SET nombre = $1, direccion = $2"
	query = query[:len(query)-2]

	// Agrega la cláusula WHERE para especificar el estudiante por "idEstudiante"
	query += fmt.Sprintf(" WHERE idEstudiante = $%d", i)

	// Agrega el valor de "id" como último parámetro en "values"
	values = append(values, id)

	// Ejecuta la consulta SQL con los valores reemplazados en cada marcador de posición
	_, err := r.db.Exec(query, values...)

	// Devuelve cualquier error que haya ocurrido durante la ejecución de la consulta
	return err
}
func (r *estudianteRepo) DeleteEstudiante(id string) error {
	_, err := r.db.Exec("DELETE FROM estudiantes WHERE idEstudiante=$1", id)
	return err
}
