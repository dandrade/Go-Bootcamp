# Arrays, Slices, Maps y Strings en Go

## Arrays

Los arrays en Go tienen un tamaño fijo y se declaran con un tipo específico:

```go
// Declaración de array
var numbers [5]int              // Array de 5 enteros (valores cero)
var names [3]string = [3]string{"Juan", "Ana", "Pedro"}

// Inicialización directa
scores := [4]int{95, 87, 92, 78}

// Tamaño automático
fruits := [...]string{"manzana", "banana", "naranja"}
```

### Características de Arrays:
- Tamaño fijo en tiempo de compilación
- Valores inicializados a cero por defecto
- Acceso por índice: `array[index]`
- Función `len()` para obtener longitud

## Slices

Los slices son más flexibles que los arrays:

```go
// Declaración de slice
var numbers []int                    // Slice vacío
var names []string = []string{"Juan", "Ana"}

// Usando make
scores := make([]int, 5)            // Slice de 5 elementos (ceros)
buffer := make([]int, 0, 10)        // Longitud 0, capacidad 10

// Literales de slice
fruits := []string{"manzana", "banana", "naranja"}
```

### Operaciones con Slices:

```go
// Agregar elementos
fruits = append(fruits, "uva")
fruits = append(fruits, "pera", "kiwi")

// Crear slice desde array/slice
numbers := []int{1, 2, 3, 4, 5}
subset := numbers[1:4]              // [2, 3, 4]
prefix := numbers[:3]               // [1, 2, 3]
suffix := numbers[2:]               // [3, 4, 5]

// Copiar slices
dest := make([]int, len(numbers))
copy(dest, numbers)
```

## Maps

Los maps almacenan pares key-value:

```go
// Declaración de map
var ages map[string]int
ages = make(map[string]int)

// Inicialización directa
ages := map[string]int{
    "Juan": 25,
    "Ana":  30,
    "Pedro": 28,
}

// Usando make
scores := make(map[string]int)
```

### Operaciones con Maps:

```go
// Agregar/actualizar
ages["Maria"] = 27
ages["Juan"] = 26

// Acceder a valores
age := ages["Juan"]
age, exists := ages["Carlos"]      // Verificar existencia

// Eliminar
delete(ages, "Pedro")

// Iterar
for name, age := range ages {
    fmt.Printf("%s tiene %d años\n", name, age)
}
```

## Strings

Los strings en Go son inmutables:

```go
// Declaración
var message string = "Hola mundo"
greeting := "¡Hola, Go!"

// Concatenación
fullMessage := message + " " + greeting
```

### Paquete strings:

```go
import "strings"

text := "Hola Mundo Go"

// Funciones útiles
strings.ToUpper(text)           // "HOLA MUNDO GO"
strings.ToLower(text)           // "hola mundo go"
strings.Contains(text, "Go")    // true
strings.Split(text, " ")        // ["Hola", "Mundo", "Go"]
strings.Join([]string{"a", "b"}, "-")  // "a-b"
strings.Replace(text, "Go", "Golang", 1)
```

### Runes y UTF-8:

```go
text := "Hola 世界"
runes := []rune(text)
fmt.Println(len(text))      // Bytes
fmt.Println(len(runes))     // Caracteres Unicode
```

## Iteración con range

```go
// Arrays y slices
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Solo valores
for _, value := range numbers {
    fmt.Println(value)
}

// Maps
ages := map[string]int{"Juan": 25, "Ana": 30}
for name, age := range ages {
    fmt.Printf("%s: %d\n", name, age)
}

// Strings (itera sobre runes)
for index, char := range "Hola" {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}
```

## Ejemplos Prácticos

### Lista de Tareas:
```go
type Task struct {
    ID   int
    Text string
    Done bool
}

var tasks []Task
var nextID int = 1

func addTask(text string) {
    task := Task{ID: nextID, Text: text, Done: false}
    tasks = append(tasks, task)
    nextID++
}
```

### Contador de Palabras:
```go
func countWords(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    
    return counts
}
```

## Mejores Prácticas

1. **Usa slices** en lugar de arrays cuando necesites flexibilidad
2. **Verifica la existencia** al acceder a maps
3. **Usa strings.Builder** para concatenar muchos strings
4. **Inicializa maps** antes de usarlos
5. **Considera la capacidad** al crear slices con make()

## Próximos Pasos

En los ejercicios aplicarás estos conceptos para:
- Manipular colecciones de datos
- Procesar texto
- Crear estructuras de datos simples
- Implementar algoritmos básicos