# Proyecto Semana 4: Sistema de Gestión de Tareas v2.0

## Descripción
Esta es la versión avanzada del sistema de gestión de tareas que comenzaste en la Semana 3. Ahora incorporarás interfaces para mayor flexibilidad, manejo robusto de errores, persistencia en archivos JSON, y validación completa.

## Nuevas Funcionalidades v2.0

### Mejoras Principales
- ✨ **Interfaces**: Arquitectura desacoplada con Repository pattern
- 🔒 **Validación robusta**: Custom errors informativos
- 💾 **Persistencia**: Almacenamiento en archivos JSON
- 🔄 **Auto-save**: Guardado automático de cambios
- 📊 **Métricas**: Estadísticas avanzadas
- 🎯 **Filtros avanzados**: Búsquedas y filtros complejos

### Arquitectura con Interfaces
```
main.go           # CLI mejorado
│
├── models/       # Definiciones de datos
│   ├── task.go   # Task struct (de Semana 3)
│   └── errors.go # Custom errors
│
├── interfaces/   # Definiciones de interfaces
│   └── repository.go
│
├── repository/   # Implementaciones de storage
│   ├── memory.go     # En memoria (de Semana 3)
│   └── json.go       # Persistencia JSON
│
├── service/      # Lógica de negocio
│   └── task_service.go
│
└── utils/        # Utilidades
    └── validation.go
```

## Especificaciones Técnicas

### Interfaces Principales
```go
type TaskRepository interface {
    Save(task Task) error
    GetByID(id string) (*Task, error)
    GetAll() ([]Task, error)
    Update(task Task) error
    Delete(id string) error
    FindByStatus(status Status) ([]Task, error)
    Search(query string) ([]Task, error)
}

type TaskService interface {
    CreateTask(title, description string, priority Priority) (*Task, error)
    UpdateTask(id, title, description string) error
    DeleteTask(id string) error
    GetTask(id string) (*Task, error)
    ListTasks() ([]Task, error)
    // ... más métodos
}
```

### Custom Errors
```go
type TaskNotFoundError struct {
    ID string
}

type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

type PersistenceError struct {
    Operation string
    Cause     error
}
```

### JSON Storage
- Archivo: `tasks.json`
- Auto-save en cada operación
- Backup automático
- Recovery en caso de corrupción

## Funcionalidades Implementadas

### Básicas (de Semana 3)
- [x] Crear, leer, actualizar, eliminar tareas
- [x] Estados y prioridades
- [x] Filtros básicos
- [x] CLI interactivo

### Nuevas (Semana 4)
- [ ] Repository pattern con interfaces
- [ ] Persistencia en JSON
- [ ] Validación robusta con custom errors
- [ ] Auto-save y backup
- [ ] Métricas de uso
- [ ] Búsqueda avanzada
- [ ] Import/Export de tareas
- [ ] Configuración del sistema

## Instrucciones de Implementación

### Paso 1: Migrar a nueva estructura
```bash
cd proyecto
mkdir -p {models,interfaces,repository,service,utils}
```

### Paso 2: Definir interfaces
- Crear `interfaces/repository.go`
- Definir contratos claros
- Documentar comportamientos esperados

### Paso 3: Implementar custom errors
- Crear `models/errors.go`
- Errors específicos del dominio
- Información rica para debugging

### Paso 4: Repository implementations
- `repository/memory.go` - Migrar código de Semana 3
- `repository/json.go` - Nueva persistencia
- Manejo de errores específicos

### Paso 5: Service layer
- `service/task_service.go`
- Lógica de negocio centralizada
- Validación y orquestación

### Paso 6: CLI mejorado
- Mejorar `main.go`
- Manejo de errores en UI
- Nuevas opciones de menú

## Comandos del CLI v2.0

```
=== Task Manager v2.0 ===
1.  Crear nueva tarea
2.  Listar todas las tareas
3.  Ver tarea específica
4.  Actualizar tarea
5.  Cambiar estado de tarea
6.  Eliminar tarea
7.  Buscar tareas
8.  Filtros avanzados
9.  Gestionar etiquetas
10. Fechas límite
11. Estadísticas
12. Import/Export
13. Configuración
14. Backup/Restore
0.  Salir
```

### Nuevos Filtros
- Por rango de fechas
- Por múltiples etiquetas
- Combinación estado + prioridad
- Tareas vencidas
- Tareas sin fecha límite

### Métricas Avanzadas
- Productividad por día/semana
- Tiempo promedio de completitud
- Distribución por prioridad
- Trends de creación

## Ejemplos de Uso

### Manejo de Errores
```go
task, err := service.GetTask("invalid-id")
if err != nil {
    switch e := err.(type) {
    case TaskNotFoundError:
        fmt.Printf("Tarea '%s' no encontrada\n", e.ID)
    case PersistenceError:
        fmt.Printf("Error de almacenamiento: %v\n", e.Cause)
    default:
        fmt.Printf("Error inesperado: %v\n", err)
    }
    return
}
```

### Repository Pattern
```go
// Usar memoria para testing
memRepo := repository.NewMemoryRepository()
service := service.NewTaskService(memRepo)

// Cambiar a persistencia para producción
jsonRepo := repository.NewJSONRepository("tasks.json")
service := service.NewTaskService(jsonRepo)
```

### Validación
```go
err := service.CreateTask("", "description", Low)
if err != nil {
    if valErr, ok := err.(ValidationError); ok {
        fmt.Printf("Campo '%s' inválido: %s\n", valErr.Field, valErr.Message)
    }
}
```

## Testing

### Unit Tests
```bash
# Test de cada componente
go test ./models/...
go test ./repository/...
go test ./service/...

# Test de integración
go test ./...
```

### Test de Persistencia
- Crear tareas
- Reiniciar programa
- Verificar que persisten

## Criterios de Evaluación

### Arquitectura (40%)
- [ ] Interfaces bien definidas
- [ ] Separación de responsabilidades
- [ ] Código desacoplado
- [ ] Patrones bien implementados

### Error Handling (30%)
- [ ] Custom errors informativos
- [ ] Manejo completo de casos edge
- [ ] Recovery de errores
- [ ] UX mejorado con errores

### Persistencia (20%)
- [ ] JSON storage funcional
- [ ] Auto-save confiable
- [ ] Backup y recovery
- [ ] Manejo de corrupción

### Usabilidad (10%)
- [ ] CLI intuitivo
- [ ] Funcionalidades completas
- [ ] Performance aceptable
- [ ] Documentación clara

## Comandos de Desarrollo

```bash
# Ejecutar
go run .

# Tests
go test ./...

# Verificar código
go vet ./...
go fmt ./...

# Build de producción
go build -o taskmanager-v2

# Limpiar datos
rm tasks.json tasks.backup.json
```

## Estructura de Archivos Final
```
proyecto/
├── go.mod
├── main.go                    # CLI principal
├── models/
│   ├── task.go               # Task struct
│   └── errors.go             # Custom errors
├── interfaces/
│   └── repository.go         # Repository interface
├── repository/
│   ├── memory.go             # In-memory storage
│   └── json.go               # JSON persistence
├── service/
│   └── task_service.go       # Business logic
├── utils/
│   └── validation.go         # Validation helpers
├── tasks.json                # Data file
└── README.md                 # Este archivo
```

## Próximos Pasos (Fase 3)
En la siguiente fase trabajarás con:
- Concurrencia para operaciones pesadas
- Web API REST
- Base de datos real
- Testing avanzado

---

¡Este proyecto consolidará tu conocimiento de Go y te preparará para desarrollo profesional! 🚀