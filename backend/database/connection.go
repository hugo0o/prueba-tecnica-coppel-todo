package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB es global y con mayúscula para poder usarla desde otras carpetas
var DB *sql.DB

func Connect() {
	var err error

	connStr := "user=postgres password=correcaminos dbname=todo_coppel sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error de configuración de BD: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error al conectar a PostgreSQL: ", err)
	}

	fmt.Println("✅ ¡Conexión a PostgreSQL exitosa!")
}
