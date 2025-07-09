# Go Bootcamp - De Básico a Experto 🚀

Un bootcamp completo de Go diseñado para llevarte desde nivel básico hasta experto en 12 semanas, con enfoque práctico y proyectos reales.

## 🎯 Objetivos del Bootcamp

- **Dominar Go completamente**: Desde sintaxis básica hasta arquitectura de sistemas
- **Aprendizaje práctico**: Cada concepto se aplica en proyectos reales
- **Progresión estructurada**: Contenido organizado en 6 fases incrementales
- **Desarrollo profesional**: Buenas prácticas y patrones idiomáticos de Go

## 📚 Estructura del Bootcamp

### 🏗️ Fase 1: Fundamentos (Semanas 1-2)
- **Ubicación**: [`fase1-fundamentos/`](./fase1-fundamentos/)
- **Duración**: 2 semanas
- **Temas**: Sintaxis básica, variables, funciones, control de flujo, tipos básicos
- **Proyecto**: CLI Calculator - Calculadora de línea de comandos

### 🔧 Fase 2: Go Intermedio (Semanas 3-4)
- **Ubicación**: [`fase2-intermedio/`](./fase2-intermedio/)
- **Duración**: 2 semanas  
- **Temas**: Structs, methods, interfaces, paquetes, manejo de errores
- **Proyecto**: Task Manager - Sistema de gestión de tareas (CRUD)

### ⚡ Fase 3: Concurrencia (Semanas 5-6)
- **Ubicación**: [`fase3-concurrencia/`](./fase3-concurrencia/)
- **Duración**: 2 semanas
- **Temas**: Goroutines, channels, select, patrones de concurrencia
- **Proyecto**: Web Scraper - Scraper web concurrente

### 🌐 Fase 4: Desarrollo Web (Semanas 7-8)
- **Ubicación**: [`fase4-web/`](./fase4-web/)
- **Duración**: 2 semanas
- **Temas**: HTTP, JSON, APIs REST, frameworks web (Gin/Echo)
- **Proyecto**: REST API - API completa con base de datos

### 🚀 Fase 5: Nivel Avanzado (Semanas 9-10)
- **Ubicación**: [`fase5-avanzado/`](./fase5-avanzado/)
- **Duración**: 2 semanas
- **Temas**: Testing avanzado, benchmarking, profiling, optimización
- **Proyecto**: Microservice - Microservicio con métricas y monitoreo

### 🏆 Fase 6: Proyectos Expertos (Semanas 11-12)
- **Ubicación**: [`fase6-experto/`](./fase6-experto/)
- **Duración**: 2 semanas
- **Temas**: Arquitectura de sistemas, deployment, containers, patterns
- **Proyecto**: Distributed System - Sistema distribuido completo

## 🎓 Metodología de Aprendizaje

### 📖 Estructura Semanal
1. **Lunes-Martes**: Teoría con ejemplos prácticos
2. **Miércoles-Jueves**: Ejercicios guiados y práctica
3. **Viernes**: Desarrollo del proyecto semanal
4. **Fin de semana**: Revisión, refactoring y mejoras

### 🔄 Ciclo de Aprendizaje
```
📚 Estudiar → 💻 Practicar → 🚀 Proyectar → 🔍 Revisar → ➡️ Avanzar
```

### 📁 Organización de Contenido
```
faseX-nombre/
├── semanaY/
│   ├── README.md        # Objetivos y plan de la semana
│   ├── teoria/          # Material teórico con ejemplos
│   ├── ejercicios/      # Ejercicios prácticos guiados
│   └── proyecto/        # Proyecto semanal aplicado
```

## 🚀 Cómo Empezar

### 1. Prerrequisitos
- **Go instalado**: [Descargar Go](https://golang.org/dl/)
- **Editor de código**: VS Code, GoLand, o tu preferido
- **Terminal**: Bash, Zsh, o PowerShell
- **Git**: Para control de versiones

### 2. Verificar Instalación
```bash
go version          # Debe mostrar Go 1.19 o superior
git --version       # Verificar Git
```

### 3. Comenzar el Bootcamp
```bash
# Navegar al proyecto
cd GoBootCamp

# Iniciar con Fase 1
cd fase1-fundamentos/semana1

# Leer el plan de la semana
cat README.md

# Estudiar la teoría
cat teoria/01-fundamentos.md

# Practicar con ejercicios
cd ejercicios
go run ejercicio1.go
```

## 📊 Seguimiento del Progreso

### ✅ Estado Actual
- [x] **Proyecto configurado**
- [x] **Fase 1 - Semana 1**: Material creado
- [ ] **Fase 1 - Semana 1**: En progreso
- [ ] **Fase 1 - Semana 2**: Pendiente
- [ ] **Fase 2**: Pendiente
- [ ] **Fase 3**: Pendiente
- [ ] **Fase 4**: Pendiente
- [ ] **Fase 5**: Pendiente
- [ ] **Fase 6**: Pendiente

### 🎯 Criterios de Evaluación
- **Funcionalidad**: ¿El código funciona correctamente?
- **Estilo**: ¿Sigue las convenciones de Go?
- **Buenas Prácticas**: ¿Usa patrones idiomáticos?
- **Completitud**: ¿Cumple todos los requisitos?

## 🛠️ Comandos Útiles

### Comandos Básicos de Go
```bash
go run main.go          # Ejecutar programa
go build               # Compilar proyecto
go test ./...          # Ejecutar tests
go mod init <nombre>    # Inicializar módulo
go mod tidy            # Limpiar dependencias
go fmt ./...           # Formatear código
go vet ./...           # Análisis estático
```

### Comandos Específicos del Bootcamp
```bash
# Ejecutar proyecto de semana específica
go run fase1-fundamentos/semana1/proyecto/main.go

# Ejecutar todos los tests del bootcamp
go test ./fase*/...

# Formatear todo el código del bootcamp
go fmt ./fase*/...
```

## 📋 Proyectos del Bootcamp

| Fase | Proyecto | Descripción | Tecnologías |
|------|----------|-------------|-------------|
| 1 | **CLI Calculator** | Calculadora de línea de comandos | Go básico, os.Args, fmt |
| 2 | **Task Manager** | Sistema CRUD de tareas | Structs, interfaces, archivos |
| 3 | **Web Scraper** | Scraper web concurrente | Goroutines, channels, HTTP |
| 4 | **REST API** | API completa con BD | Gin/Echo, SQL/MongoDB, JSON |
| 5 | **Microservice** | Servicio con métricas | Docker, Prometheus, Testing |
| 6 | **Distributed System** | Sistema distribuido | Kubernetes, gRPC, Patterns |

## 🎯 Objetivos de Aprendizaje por Fase

### 📚 Fase 1 - Fundamentos
- [x] Instalación y configuración de Go
- [x] Sintaxis básica y tipos de datos
- [x] Funciones y control de flujo
- [x] Arrays, slices, y maps
- [x] Manejo básico de strings
- [x] Primer proyecto CLI

### 🔧 Fase 2 - Go Intermedio
- [ ] Structs y métodos
- [ ] Interfaces y polimorfismo
- [ ] Paquetes y módulos
- [ ] Manejo de errores idiomático
- [ ] Trabajo con archivos
- [ ] Testing básico

### ⚡ Fase 3 - Concurrencia
- [ ] Goroutines y runtime
- [ ] Channels y comunicación
- [ ] Select statements
- [ ] Patrones de concurrencia
- [ ] Sincronización
- [ ] Proyecto concurrente

### 🌐 Fase 4 - Desarrollo Web
- [ ] HTTP client/server
- [ ] JSON y APIs REST
- [ ] Frameworks web
- [ ] Middleware
- [ ] Bases de datos
- [ ] Autenticación

### 🚀 Fase 5 - Nivel Avanzado
- [ ] Testing avanzado
- [ ] Benchmarking
- [ ] Profiling y optimización
- [ ] Reflection
- [ ] Generics
- [ ] Métricas y monitoreo

### 🏆 Fase 6 - Proyectos Expertos
- [ ] Arquitectura de microservicios
- [ ] Deployment y containers
- [ ] CI/CD
- [ ] Observabilidad
- [ ] Patrones avanzados
- [ ] Proyecto final

## 🤝 Sistema de Revisión

### Para el Estudiante:
1. **Completa** ejercicios y proyectos
2. **Comparte** código para revisión
3. **Recibe** feedback y sugerencias
4. **Implementa** mejoras
5. **Avanza** al siguiente nivel

### Para el Mentor (Claude):
1. **Revisa** funcionalidad y estilo
2. **Sugiere** mejoras y alternativas
3. **Explica** conceptos avanzados
4. **Actualiza** progreso
5. **Prepara** siguiente fase

## 📖 Recursos Adicionales

### Documentación Oficial
- [Go Tour](https://go.dev/tour/) - Tutorial interactivo
- [Go Documentation](https://go.dev/doc/) - Documentación oficial
- [Effective Go](https://go.dev/doc/effective_go) - Mejores prácticas

### Libros Recomendados
- "The Go Programming Language" - Alan Donovan & Brian Kernighan
- "Go in Action" - William Kennedy
- "Concurrency in Go" - Katherine Cox-Buday

### Comunidad
- [Go Forum](https://forum.golangbridge.org/)
- [Reddit r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

## 🎉 Al Completar el Bootcamp

### 🏆 Habilidades Adquiridas
- **Go Experto**: Dominio completo del lenguaje
- **Desarrollo Web**: APIs y servicios web
- **Concurrencia**: Programación paralela eficiente
- **Arquitectura**: Sistemas distribuidos
- **DevOps**: Deployment y monitoreo
- **Testing**: Código confiable y mantenible

### 🚀 Próximos Pasos
- **Contribuir** a proyectos open source
- **Crear** tus propios proyectos
- **Mentorear** a otros desarrolladores
- **Especializarse** en áreas específicas
- **Certificaciones** y reconocimientos

---

## 📞 Soporte

Para dudas, revisiones o ayuda:
- **Comparte tu código** para revisión
- **Pregunta conceptos** específicos
- **Solicita explicaciones** adicionales
- **Pide feedback** en proyectos

**¡Comencemos este viaje hacia la maestría en Go!** 🎯

---

*Bootcamp creado con ❤️ para aprender Go de forma práctica y efectiva*