# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a comprehensive Go bootcamp repository designed to take a developer from basic to expert level over 12 weeks. The bootcamp follows a practical, project-based learning approach with incremental complexity.

## Bootcamp Structure

### Fase 1: Fundamentos (Semanas 1-2)
- **Ubicación**: `fase1-fundamentos/`
- **Objetivo**: Dominar sintaxis básica, variables, funciones, control de flujo
- **Proyecto**: CLI tool (calculadora/conversor)

### Fase 2: Go Intermedio (Semanas 3-4)
- **Ubicación**: `fase2-intermedio/`
- **Objetivo**: Structs, methods, interfaces, manejo de errores
- **Proyecto**: Sistema de gestión de tareas (CRUD)

### Fase 3: Concurrencia (Semanas 5-6)
- **Ubicación**: `fase3-concurrencia/`
- **Objetivo**: Goroutines, channels, patrones de concurrencia
- **Proyecto**: Web scraper concurrente

### Fase 4: Desarrollo Web (Semanas 7-8)
- **Ubicación**: `fase4-web/`
- **Objetivo**: HTTP, JSON, APIs REST, frameworks web
- **Proyecto**: API completa con base de datos

### Fase 5: Nivel Avanzado (Semanas 9-10)
- **Ubicación**: `fase5-avanzado/`
- **Objetivo**: Testing avanzado, benchmarking, optimización
- **Proyecto**: Microservicio con métricas

### Fase 6: Proyectos Expertos (Semanas 11-12)
- **Ubicación**: `fase6-experto/`
- **Objetivo**: Arquitectura de sistemas, deployment, containers
- **Proyecto**: Sistema distribuido

## Development Commands

### Comandos básicos de Go:
- `go run main.go` - Ejecutar aplicación
- `go build` - Compilar proyecto
- `go test ./...` - Ejecutar todos los tests
- `go mod init <name>` - Inicializar módulo
- `go mod tidy` - Limpiar dependencias
- `go fmt ./...` - Formatear código
- `go vet ./...` - Análisis estático

### Comandos específicos del bootcamp:
- `go run fase1-fundamentos/semana1/proyecto/main.go` - Ejecutar proyecto de semana específica
- `go test ./fase*/...` - Ejecutar todos los tests del bootcamp

## Metodología de Aprendizaje

Cada semana sigue esta estructura:
1. **Teoría**: Conceptos fundamentales con ejemplos
2. **Ejercicios**: Práctica guiada de conceptos
3. **Proyecto**: Aplicación práctica del conocimiento
4. **Revisión**: Refactoring y mejores prácticas

## Organización de Archivos

```
faseX-nombre/
├── semanaY/
│   ├── teoria/          # Material de estudio
│   ├── ejercicios/      # Prácticas guiadas
│   ├── proyecto/        # Proyecto semanal
│   └── README.md        # Objetivos y instrucciones
```

## Progreso Actual

- ✅ Estructura del proyecto creada
- 🔄 **Fase 1 - Fundamentos**
  - ✅ Semana 1: Material creado
  - ⏳ Semana 1: En progreso
  - ⏳ Semana 2: Pendiente

## Sistema de Revisión

### Para el Estudiante:
1. **Completa ejercicios/proyectos**
2. **Comparte código con Claude** (copy/paste)
3. **Recibe feedback y sugerencias**
4. **Implementa mejoras**
5. **Avanza al siguiente nivel**

### Comandos Útiles:
- `go run ejercicio1.go` - Ejecutar ejercicio
- `go fmt ejercicio1.go` - Formatear código
- `go build` - Verificar compilación

### Criterios de Evaluación:
- **Funcionalidad**: ¿El código funciona correctamente?
- **Estilo**: ¿Sigue las convenciones de Go?
- **Buenas Prácticas**: ¿Usa patrones idiomáticos?
- **Completitud**: ¿Cumple todos los requisitos?