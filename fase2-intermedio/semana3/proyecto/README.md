# Proyecto Semana 3: Sistema de Gestión de Tareas v1.0

## Descripción
En esta semana crearás la primera versión de un sistema de gestión de tareas usando structs y methods. Esta es la base del proyecto que completarás en la Semana 4 con interfaces y manejo de errores.

## Objetivos
- Diseñar e implementar structs para representar tareas
- Usar methods para operaciones sobre tareas
- Crear un CLI básico para interactuar con las tareas
- Implementar almacenamiento en memoria

## Funcionalidades a Implementar

### Básicas
- [ ] Crear nuevas tareas
- [ ] Listar todas las tareas
- [ ] Mostrar detalles de una tarea
- [ ] Marcar tareas como completadas
- [ ] Eliminar tareas

### Avanzadas
- [ ] Filtrar tareas por estado
- [ ] Buscar tareas por título
- [ ] Ordenar tareas por prioridad/fecha
- [ ] Mostrar estadísticas básicas

## Estructura del Proyecto

```
proyecto/
├── README.md          # Este archivo
├── main.go           # CLI y menú principal
├── task.go           # Definición de Task y methods
├── taskmanager.go    # Gestor de tareas (slice en memoria)
└── go.mod            # Módulo Go
```

## Especificaciones

### Task Struct
```go
type Task struct {
    ID          string
    Title       string
    Description string
    Priority    Priority    // enum: Low, Medium, High
    Status      Status      // enum: Pending, InProgress, Completed
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DueDate     *time.Time  // Opcional
    Tags        []string
}
```

### Priority y Status
```go
type Priority int
const (
    Low Priority = iota
    Medium
    High
)

type Status int
const (
    Pending Status = iota
    InProgress
    Completed
)
```

### TaskManager
```go
type TaskManager struct {
    tasks  []Task
    nextID int
}
```

## Methods Requeridos

### Para Task
- `String() string` - Representación de texto
- `IsOverdue() bool` - Si la tarea está vencida
- `MarkComplete()` - Marcar como completada
- `AddTag(tag string)` - Agregar etiqueta
- `RemoveTag(tag string)` - Remover etiqueta

### Para TaskManager
- `AddTask(title, description string, priority Priority) *Task`
- `GetTask(id string) *Task`
- `GetAllTasks() []Task`
- `GetTasksByStatus(status Status) []Task`
- `DeleteTask(id string) bool`
- `SearchTasks(query string) []Task`

## CLI Menú
```
=== Task Manager ===
1. Crear nueva tarea
2. Listar todas las tareas
3. Ver tarea específica
4. Marcar tarea como completada
5. Eliminar tarea
6. Buscar tareas
7. Filtrar por estado
8. Mostrar estadísticas
9. Salir
```

## Instrucciones de Implementación

### Paso 1: Configurar el proyecto
```bash
cd proyecto
go mod init taskmanager
```

### Paso 2: Implementar task.go
- Define los tipos Task, Priority, Status
- Implementa todos los methods de Task
- Usa constructores como `NewTask()`

### Paso 3: Implementar taskmanager.go
- Define TaskManager struct
- Implementa todos los methods del manager
- Maneja la generación de IDs únicos

### Paso 4: Implementar main.go
- Crea el CLI interactivo
- Conecta el menú con TaskManager
- Maneja input del usuario

### Paso 5: Testing
- Crea tareas de ejemplo
- Prueba todas las funcionalidades
- Verifica edge cases

## Ejemplos de Uso

### Crear Tarea
```
Título: Revisar documentación
Descripción: Revisar docs de Go para el proyecto
Prioridad: (1) Baja, (2) Media, (3) Alta: 2
Fecha límite (YYYY-MM-DD) [opcional]: 2024-01-15
```

### Listar Tareas
```
ID   | Título                | Estado      | Prioridad | Fecha Límite
-----|----------------------|-------------|-----------|-------------
T001 | Revisar documentación | Pendiente   | Media     | 2024-01-15
T002 | Implementar API       | En Progreso | Alta      | -
```

## Criterios de Evaluación

### Funcionalidad (60%)
- [ ] Todas las operaciones CRUD funcionan
- [ ] CLI es intuitivo y funcional
- [ ] Filtros y búsquedas funcionan correctamente

### Código (30%)
- [ ] Structs bien diseñados
- [ ] Methods apropiados (value vs pointer receivers)
- [ ] Código limpio y bien organizado
- [ ] Nombres descriptivos

### Extras (10%)
- [ ] Funcionalidades adicionales
- [ ] Interfaz de usuario mejorada
- [ ] Validación robusta de datos

## Comandos Útiles

```bash
# Ejecutar el programa
go run .

# Verificar el código
go vet ./...
go fmt ./...

# Compilar
go build -o taskmanager

# Ejecutar compilado
./taskmanager
```

## Próximos Pasos (Semana 4)
En la Semana 4 mejorarás este sistema con:
- Interfaces para flexibilidad
- Manejo robusto de errores
- Persistencia en archivos JSON
- Validación avanzada
- Testing unitario

## Tips de Desarrollo

1. **Comienza simple**: Implementa lo básico primero
2. **Itera**: Mejora gradualmente la funcionalidad
3. **Testa frecuentemente**: Ejecuta el programa constantemente
4. **Usa go fmt**: Mantén el código formateado
5. **Valida inputs**: Siempre verifica las entradas del usuario

---

¡Buena suerte con tu primer proyecto de Go! 🚀