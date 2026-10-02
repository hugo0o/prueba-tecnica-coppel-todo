package models

import "time"

// Tarea debe empezar con mayúscula para que sea pública en todo tu proyecto
type Tarea struct {
	ID                 int       `json:"id_tarea"`
	Titulo             string    `json:"titulo"`
	Descripcion        string    `json:"descripcion"`
	Completada         bool      `json:"completada"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}
