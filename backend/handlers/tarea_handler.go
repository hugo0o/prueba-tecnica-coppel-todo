package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// Importamos los otros paquetes de tu proyecto
	"todolist-backend/database"
	"todolist-backend/models"
)

// TareasHandler empieza con mayúscula para que main.go pueda usarla
func TareasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getTareas(w, r)
	case http.MethodPost:
		crearTarea(w, r)
	case http.MethodPut:
		editarTarea(w, r)
	case http.MethodDelete:
		eliminarTarea(w, r)
	default:
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}

// Empieza con minúscula porque solo se usa dentro de este archivo
func getTareas(w http.ResponseWriter, r *http.Request) {
	// Iniciamos con la consulta base y luego agregamos filtros dinámicamente
	baseQuery := `SELECT id_tarea, titulo, descripcion, completada, fecha_creacion, fecha_actualizacion FROM tareas WHERE 1=1`

	// Arreglo para guardar los valores de los parámetros dinámicos
	var args []interface{}
	paramIndex := 1

	// 1. Filtro por estado (?estado=completadas o ?estado=pendientes)
	estado := r.URL.Query().Get("estado")
	if estado == "completadas" {
		baseQuery += fmt.Sprintf(" AND completada = $%d", paramIndex)
		args = append(args, true)
		paramIndex++
	} else if estado == "pendientes" {
		baseQuery += fmt.Sprintf(" AND completada = $%d", paramIndex)
		args = append(args, false)
		paramIndex++
	}

	// 2. Buscador por título (?buscar=palabra)
	buscar := r.URL.Query().Get("buscar")
	if buscar != "" {
		baseQuery += fmt.Sprintf(" AND titulo ILIKE $%d", paramIndex)
		args = append(args, "%"+buscar+"%")
		paramIndex++
	}

	// Siempre ordenamos para que la vista sea consistente
	baseQuery += " ORDER BY id_tarea ASC"

	// Ejecutamos la consulta pasándole el arreglo de argumentos desenpaquetado (args...)
	rows, err := database.DB.Query(baseQuery, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tareas := []models.Tarea{}
	for rows.Next() {
		var t models.Tarea
		if err := rows.Scan(&t.ID, &t.Titulo, &t.Descripcion, &t.Completada, &t.FechaCreacion, &t.FechaActualizacion); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tareas = append(tareas, t)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tareas)
}

func crearTarea(w http.ResponseWriter, r *http.Request) {
	var t models.Tarea
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO tareas (titulo, descripcion) VALUES ($1, $2) RETURNING id_tarea, completada, fecha_creacion, fecha_actualizacion`
	err := database.DB.QueryRow(sqlStatement, t.Titulo, t.Descripcion).Scan(&t.ID, &t.Completada, &t.FechaCreacion, &t.FechaActualizacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func editarTarea(w http.ResponseWriter, r *http.Request) {
	//Extraer el ID de la tarea desde la URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID de tarea no proporcionado", http.StatusBadRequest)
		return
	}

	// Decodificar el JSON con los nuevos datos de la tarea
	var t models.Tarea
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Actualizar la tarea en la base de datos Update
	sqlStatement := `UPDATE tareas SET titulo=$1, descripcion=$2, completada=$3, fecha_actualizacion = CURRENT_TIMESTAMP WHERE id_tarea=$4`

	// Ejecutar la actualización con DB.Exec
	_, err := database.DB.Exec(sqlStatement, t.Titulo, t.Descripcion, t.Completada, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Tarea actualizada correctamente"})

}

// Función para ELIMINAR (DELETE)
func eliminarTarea(w http.ResponseWriter, r *http.Request) {
	// 1. Extraemos el ID de la URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Falta el ID de la tarea en la URL", http.StatusBadRequest)
		return
	}

	// 2. Ejecutamos el DELETE
	sqlStatement := `DELETE FROM tareas WHERE id_tarea = $1`
	_, err := database.DB.Exec(sqlStatement, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Tarea eliminada correctamente"})
}
