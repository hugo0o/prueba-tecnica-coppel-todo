import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
// 1. Importamos FormsModule y FormControl para el buscador y la edición
import { ReactiveFormsModule, FormsModule, FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';
import { TareaService, Tarea } from './services/tarea.service';
// 2. IMPORTANTE: Agregamos startWith, combineLatest, map, debounceTime, distinctUntilChanged
import { Observable, BehaviorSubject, combineLatest, switchMap, map, debounceTime, distinctUntilChanged, startWith } from 'rxjs';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, FormsModule],
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  private refresh$ = new BehaviorSubject<void>(undefined);
  
  // --- VARIABLES PARA EL BUSCADOR ---
  buscadorControl = new FormControl('');

  // 3. starwith('') es necesario para que el buscador funcione correctamente desde el inicio, ya que de lo contrario no emitirá ningún valor hasta que el usuario escriba algo.
  tareas$: Observable<Tarea[]> = combineLatest([
    this.refresh$.pipe(switchMap(() => this.tareaService.getTareas())),
    this.buscadorControl.valueChanges.pipe(
      startWith(''), 
      debounceTime(300),
      distinctUntilChanged()
    )
  ]).pipe(
    map(([tareas, filtro]) => {
      // Si el filtro está vacío (o es null), regresa todas las tareas
      if (!filtro) return tareas;
      
      const textoFiltro = filtro.toLowerCase();
      // Filtra si el título o la descripción incluyen el texto buscado
      return tareas.filter(t => 
        t.titulo.toLowerCase().includes(textoFiltro) || 
        t.descripcion.toLowerCase().includes(textoFiltro)
      );
    })
  );

  // --- VARIABLES PARA EL FORMULARIO DE CREACIÓN ---
  tareaForm: FormGroup;

  // --- VARIABLES PARA LA EDICIÓN ---
  editandoId: number | null = null;
  tituloEditado: string = '';
  descripcionEditada: string = '';

  constructor(
    private tareaService: TareaService,
    private fb: FormBuilder
  ) {
    this.tareaForm = this.fb.group({
      titulo: ['', [Validators.required]],
      descripcion: ['', [Validators.required]]
    });
  }    

  ngOnInit(): void {}

  private recargarTabla(): void {
    this.refresh$.next();
  }

  // --- MÉTODOS DE CREACIÓN Y ESTADO ---
  guardarTarea(): void {
    if (this.tareaForm.invalid) return;

    const nuevaTarea = this.tareaForm.value;

    this.tareaService.crearTarea(nuevaTarea).subscribe({
      next: () => {
        this.tareaForm.reset();
        this.recargarTabla();
      },
      error: (err) => console.error('Error al crear la tarea:', err)
    });
  }

  toggleCompletada(tarea: Tarea): void {
    if (tarea.id_tarea === undefined) return;

    const tareaActualizada: Tarea = {
      ...tarea,
      completada: !tarea.completada
    };

    this.tareaService.editarTarea(tarea.id_tarea, tareaActualizada).subscribe({
      next: () => this.recargarTabla(),
      error: (err) => console.error('Error al actualizar el estatus:', err)
    });
  }

  eliminarTarea(id: number | undefined): void {
    if (id === undefined) return;
    if (!confirm('¿Estás seguro de que deseas eliminar esta tarea?')) return;

    this.tareaService.eliminarTarea(id).subscribe({
      next: () => this.recargarTabla(),
      error: (err) => console.error('Error al eliminar la tarea:', err)
    });
  }

  // --- MÉTODOS DE EDICIÓN ---
  iniciarEdicion(tarea: Tarea): void {
    if (tarea.id_tarea === undefined) return;
    this.editandoId = tarea.id_tarea;
    this.tituloEditado = tarea.titulo;
    this.descripcionEditada = tarea.descripcion;
  }

  cancelarEdicion(): void {
    this.editandoId = null;
  }

  guardarEdicion(tarea: Tarea): void {
    if (tarea.id_tarea === undefined) return;

    const tareaActualizada: Tarea = {
      ...tarea,
      titulo: this.tituloEditado,
      descripcion: this.descripcionEditada
    };

    this.tareaService.editarTarea(tarea.id_tarea, tareaActualizada).subscribe({
      next: () => {
        this.editandoId = null; // Quita el modo edición
        this.recargarTabla(); // Refresca los datos
      },
      error: (err) => console.error('Error al actualizar la tarea:', err)
    });
  }
}