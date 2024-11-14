package env

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// LoadEnv carga el archivo .env si aún no se ha cargado
func LoadEnv() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No se encontró archivo .env, Procediendo con variables de entorno alternas")
		}
	})
}

// GetEnv devuelve el valor de una variable de entorno
func GetEnv(key, fallback string) string {
	LoadEnv() // Asegura que el entorno esté cargado
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
