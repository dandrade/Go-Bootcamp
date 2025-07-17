# Semana 3: Structs y Methods

## Objetivos de la Semana
- Entender qué son los structs y cómo definirlos
- Aprender a crear methods con diferentes tipos de receivers
- Dominar la composición de structs
- Implementar patrones de constructor
- Iniciar el desarrollo del proyecto de gestión de tareas

## Contenido

### 1. Teoría (teoria/)
- **03-structs-methods.md**: Guía completa sobre structs, methods, composición y embedding

### 2. Ejercicios (ejercicios/)
- **ejercicio1.go**: Structs básicos - Crear y manipular estructuras simples
- **ejercicio2.go**: Methods y receivers - Value vs pointer receivers
- **ejercicio3.go**: Composición de structs - Construir estructuras complejas
- **ejercicio4.go**: Embedding - Herencia al estilo Go

### 3. Proyecto (proyecto/)
- **Sistema de Gestión de Tareas v1.0**
  - Definir la estructura Task
  - Implementar operaciones básicas
  - Crear un CLI simple
  - Almacenamiento en memoria

## Plan de Estudio Recomendado

### Día 1-2: Teoría y Conceptos
1. Lee el material teórico completo
2. Experimenta con los ejemplos de código
3. Toma notas de los conceptos clave

### Día 3-4: Ejercicios Prácticos
1. Completa cada ejercicio en orden
2. No avances hasta entender completamente cada concepto
3. Experimenta modificando los ejercicios

### Día 5-7: Proyecto
1. Diseña la estructura de tu sistema
2. Implementa funcionalidad básica
3. Prueba exhaustivamente tu código

## Conceptos Clave de la Semana

### Structs
```go
type Person struct {
    Name string
    Age  int
}
```

### Methods
```go
func (p Person) Greet() string {
    return fmt.Sprintf("Hola, soy %s", p.Name)
}
```

### Composición
```go
type Employee struct {
    Person
    Department string
}
```

## Criterios de Éxito
- [ ] Puedes crear y usar structs correctamente
- [ ] Entiendes cuándo usar value vs pointer receivers
- [ ] Dominas la composición de structs
- [ ] Has iniciado el proyecto de gestión de tareas
- [ ] Tu código pasa `go fmt` y `go vet`

## Recursos Adicionales
- [A Tour of Go - Structs](https://tour.golang.org/moretypes/2)
- [Effective Go - Methods](https://golang.org/doc/effective_go#methods)
- [Go Blog - Method Sets](https://golang.org/doc/faq#methods_on_values_or_pointers)

## Comandos para esta Semana
```bash
# Ejecutar ejercicios
go run ejercicio1.go

# Iniciar el proyecto
cd proyecto
go mod init taskmanager
go run main.go

# Verificar código
go fmt ./...
go vet ./...
```

---

¡Adelante! Esta semana sentarás las bases para escribir código Go más estructurado y profesional. 💪