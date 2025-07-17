# Fase 2: Go Intermedio

## Descripción General
En esta fase profundizarás en las características que hacen a Go un lenguaje poderoso para el desarrollo de software robusto. Aprenderás sobre estructuras de datos personalizadas, programación orientada a objetos al estilo Go, y manejo profesional de errores.

## Objetivos de Aprendizaje
- Dominar el uso de structs y methods
- Entender la composición sobre la herencia
- Trabajar con interfaces para crear código flexible
- Implementar manejo robusto de errores
- Crear aplicaciones con arquitectura bien estructurada

## Contenido

### Semana 3: Structs y Methods
- **Teoría**: Estructuras de datos personalizadas
- **Práctica**: Methods, receivers, y composición
- **Proyecto**: Inicio del sistema de gestión de tareas

### Semana 4: Interfaces y Manejo de Errores
- **Teoría**: Interfaces y polimorfismo en Go
- **Práctica**: Error handling patterns
- **Proyecto**: Sistema CRUD completo con persistencia

## Proyecto Final: Sistema de Gestión de Tareas
Desarrollarás un sistema completo de gestión de tareas que incluye:
- Crear, leer, actualizar y eliminar tareas
- Diferentes estados y prioridades
- Persistencia de datos
- Interfaz de línea de comandos interactiva
- Manejo profesional de errores

## Requisitos Previos
- Haber completado la Fase 1 (Fundamentos)
- Conocimiento de arrays, slices, maps y strings
- Familiaridad con funciones y control de flujo

## Estructura del Proyecto
```
fase2-intermedio/
├── semana3/           # Structs y Methods
│   ├── teoria/        # Conceptos y ejemplos
│   ├── ejercicios/    # Práctica guiada
│   └── proyecto/      # Sistema de tareas v1
└── semana4/           # Interfaces y Errores
    ├── teoria/        # Conceptos avanzados
    ├── ejercicios/    # Práctica avanzada
    └── proyecto/      # Sistema de tareas v2
```

## Cómo Estudiar
1. Lee el material teórico de cada semana
2. Completa los ejercicios en orden
3. Implementa el proyecto incrementalmente
4. Revisa tu código con `go fmt` y `go vet`
5. Comparte tu progreso para recibir feedback

## Comandos Útiles
```bash
# Inicializar módulo para el proyecto
cd fase2-intermedio/semana3/proyecto
go mod init taskmanager

# Ejecutar el proyecto
go run .

# Verificar tu código
go fmt ./...
go vet ./...

# Compilar el proyecto
go build -o taskmanager
```

## Recursos Adicionales
- [Effective Go - Structs](https://golang.org/doc/effective_go#composite_literals)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Error handling in Go](https://blog.golang.org/error-handling-and-go)

---

¡Prepárate para llevar tu conocimiento de Go al siguiente nivel! 🚀