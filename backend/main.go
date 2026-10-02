package main

import (
	"fmt"
	"log"
	"net/http"

	"todolist-backend/database"
	"todolist-backend/handlers"
)

func main() {
	// 1. Iniciar conexión a BD
	database.Connect()
	defer database.DB.Close()

	// 2. Definir rutas
	http.HandleFunc("/tareas", handlers.TareasHandler)

	// 3. Levantar servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
