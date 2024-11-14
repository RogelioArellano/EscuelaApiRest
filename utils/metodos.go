package utils

import (
	"fmt"
	"sync"
	"time"
)

// GenerarClaveEstudiante genera una clave única basada en nombre, apellido y fecha
func GenerarClaveEstudiante(nombre string, folio int) string {
	// Obtener los primeros 2 caracteres del nombre y apellido, si existen
	nombreParte := ""

	if len(nombre) >= 2 {
		nombreParte = nombre[:2]
	}

	// Formatear la fecha actual en "DDMMYY"
	fechaParte := time.Now().Format("020106")

	// Combinar para formar la clave
	return fmt.Sprintf("%s%s%s%05d", ClaveEscuela, nombreParte, fechaParte, folio)
}

var (
	currentFolio int
	folioMutex   sync.Mutex
)
