import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

// Definimos la interfaz Tarea para tipar los objetos que manejaremos
export interface Tarea {
  id_tarea?: number;
  titulo: string;
  descripcion: string;
  completada?: boolean;
  fecha_creacion?: string;
  fecha_actualizacion?: string;
}

@Injectable({
  providedIn: 'root'
})
export class TareaService {
  // 2. Definimos la URL base de nuestra API
  private apiUrl = 'http://localhost:8080/tareas';

  // 3. Inyectamos HttpClient en el constructor para poder hacer peticiones HTTP
  constructor(private http: HttpClient) { }

  // GET: Consultar tareas (usando las 2 funciones extra que agregamos)
  getTareas(estado?: string, buscar?: string): Observable<Tarea[]> {
    let params = new HttpParams();
    
    // Angular armará los Query Parameters (?estado=...&buscar=...) automáticamente
    if (estado) {
      params = params.set('estado', estado);
    }
    if (buscar) {
      params = params.set('buscar', buscar);
    }

    return this.http.get<Tarea[]>(this.apiUrl, { params });
  }

  // POST: Crear una nueva tarea
  crearTarea(tarea: Tarea): Observable<any> {
    return this.http.post(this.apiUrl, tarea);
  }

  // PUT: Editar una tarea existente
  editarTarea(id: number, tarea: Tarea): Observable<any> {
    return this.http.put(`${this.apiUrl}?id=${id}`, tarea);
  }

  // DELETE: Eliminar una tarea
  eliminarTarea(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}?id=${id}`);
  }
}