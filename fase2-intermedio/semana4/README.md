# Semana 4: Interfaces y Manejo de Errores

## Objetivos de la Semana
- Dominar el uso de interfaces en Go
- Entender el polimorfismo al estilo Go
- Implementar manejo robusto de errores
- Crear custom errors informativos
- Completar el proyecto con persistencia y validación

## Contenido

### 1. Teoría (teoria/)
- **04-interfaces-errors.md**: Guía completa sobre interfaces, polimorfismo y error handling

### 2. Ejercicios (ejercicios/)
- **ejercicio1.go**: Interfaces básicas - Definir e implementar interfaces
- **ejercicio2.go**: Polimorfismo - Diferentes tipos implementando misma interface
- **ejercicio3.go**: Custom errors - Crear errores personalizados
- **ejercicio4.go**: Error handling patterns - Patrones avanzados de manejo

### 3. Proyecto (proyecto/)
- **Sistema de Gestión de Tareas v2.0**
  - Implementar interfaces para flexibilidad
  - Agregar persistencia en archivos JSON
  - Manejo robusto de errores
  - Validación completa de datos

## Plan de Estudio Recomendado

### Día 1-2: Teoría y Conceptos
1. Lee el material teórico sobre interfaces
2. Comprende el concepto de polimorfismo en Go
3. Estudia el sistema de errores de Go

### Día 3-4: Ejercicios Prácticos
1. Completa ejercicios de interfaces básicas
2. Practica polimorfismo con diferentes tipos
3. Implementa custom errors

### Día 5-7: Proyecto Final
1. Refactoriza el proyecto de la Semana 3
2. Implementa interfaces para storage
3. Agrega persistencia en JSON
4. Mejora el manejo de errores

## Conceptos Clave de la Semana

### Interfaces
```go
type Writer interface {
    Write([]byte) (int, error)
}
```

### Polimorfismo
```go
func ProcessData(w Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}
```

### Custom Errors
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}
```

## Criterios de Éxito
- [ ] Entiendes cómo definir e implementar interfaces
- [ ] Usas polimorfismo efectivamente
- [ ] Creas custom errors informativos
- [ ] Manejas errores de forma robusta
- [ ] Has completado el proyecto con persistencia

## Recursos Adicionales
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go#interfaces)
- [Error handling in Go](https://blog.golang.org/error-handling-and-go)

## Comandos para esta Semana
```bash
# Ejecutar ejercicios
go run ejercicio1.go

# Proyecto final
cd proyecto
go run .

# Testing
go test ./...

# Verificar código
go fmt ./...
go vet ./...
```

---

¡Esta semana completas las bases sólidas de Go! 🎯