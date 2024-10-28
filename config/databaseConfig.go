package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"escuelaApiREST/env"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := env.GetEnv("DBSERVER", "localhost")
	port := env.GetEnv("DBPORT", "5432")
	user := env.GetEnv("DBUSER", "postgres")
	password := env.GetEnv("DBPASSWORD", "password")

	// Conectar a la base de datos "postgres" para poder crear la base de datos "escuela"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}
	defer db.Close()

	// Crear la base de datos "escuela" si no existe
	_, err = db.Exec("CREATE DATABASE escuela")
	if err != nil {
		log.Println("La base de datos escuela ya existe o no se pudo crear:", err)
	}

	// Pausa breve para asegurar que la base de datos esté lista
	time.Sleep(2 * time.Second)

	// Conectar a la base de datos "escuela" después de crearla o confirmar que existe
	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=escuela sslmode=disable", host, port, user, password)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error al abrir la base de datos escuela:", err)
	}

	// Probar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error conectando a la base de datos escuela:", err)
	}

	// Crear la tabla "estudiantes" si no existe, con ID auto-incremental
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS estudiantes (
		idEstudiante SERIAL PRIMARY KEY,
		nombre VARCHAR(100),
		direccion VARCHAR(100),
		email VARCHAR(50),
		telefono VARCHAR(13),
		altaLocal BOOLEAN,
		altaSep BOOLEAN
	)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error al crear la tabla estudiantes:", err)
	}

	log.Println("Conexión exitosa y tabla estudiantes disponible.")
	return db, nil
}
