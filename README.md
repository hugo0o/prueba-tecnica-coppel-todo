# 📋 Prueba Técnica Coppel - To-Do List Full Stack

Aplicación web Full Stack para la gestión de tareas (To-Do List), desarrollada como parte de la prueba técnica para Coppel. El sistema permite realizar operaciones CRUD completas y cuenta con una arquitectura desacoplada y reactiva.

## 🚀 Tecnologías Utilizadas (Stack Requerido)
- **Frontend:** Angular 21 (Standalone Components, RxJS, Formularios Reactivos)
- **Backend:** Golang 1.25 (API REST)
- **Base de Datos:** PostgreSQL
- **Estilos:** Bootstrap 5

## ⚙️ Funcionalidades Principales (CRUD)
- ✅ **Agregar:** Registro de nuevas tareas con validaciones en tiempo real.
- ✅ **Editar:** Modificación de tareas existentes.
- ✅ **Eliminar:** Borrado de tareas con confirmación de seguridad.
- ✅ **Consultar:** Listado de tareas actualizadas.

## 🌟 Valor Agregado (Propuestas del Candidato)
Para cumplir con la sección de *Valor Agregado*, se implementaron las siguientes 2 funcionalidades extra para mejorar la experiencia de usuario (UX):
1. **Buscador en Tiempo Real (RxJS):** Filtrado instantáneo de tareas por título o descripción utilizando operadores reactivos (`combineLatest`, `debounceTime`) sin necesidad de hacer múltiples peticiones al servidor ni recargar la página.
2. **Edición en Línea (Inline Editing):** Capacidad de modificar el texto de una tarea directamente en su fila dentro de la tabla, logrando una interfaz mucho más limpia y rápida sin abrir ventanas adicionales.

## 🛠️ Instrucciones de Ejecución

### 1. Base de Datos (PostgreSQL)
Asegúrate de tener PostgreSQL corriendo y ejecuta el script SQL para crear la tabla de tareas (incluido en el código).

### 2. Backend (Golang)
Navega a la carpeta del backend y ejecuta el servidor:
```bash
cd backend
go run main.go