# Fundamentos de Go - Teoría

## 1. ¿Qué es Go?

Go es un lenguaje de programación desarrollado por Google en 2009. Es conocido por:
- **Simplicidad**: Sintaxis limpia y minimalista
- **Concurrencia**: Soporte nativo para programación concurrente
- **Rendimiento**: Compilación rápida y ejecución eficiente
- **Productividad**: Herramientas integradas excelentes

## 2. Configuración del Entorno

### Instalación
```bash
# Descargar desde https://golang.org/dl/
# O usando un gestor de paquetes:
brew install go  # macOS
sudo apt install golang-go  # Ubuntu
```

### Verificación
```bash
go version
```

### Estructura de un Proyecto Go
```
mi-proyecto/
├── go.mod          # Información del módulo
├── main.go         # Archivo principal
└── utils/
    └── helpers.go  # Paquetes auxiliares
```

## 3. Tu Primer Programa

```go
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, mundo!")
}
```

**Explicación línea por línea:**
- `package main`: Define el paquete principal
- `import "fmt"`: Importa el paquete para formatear texto
- `func main()`: Función principal, punto de entrada
- `fmt.Println()`: Imprime texto con salto de línea

## 4. Variables y Tipos de Datos

### Declaración de Variables

```go
// Declaración explícita
var nombre string = "Juan"
var edad int = 25

// Declaración con inferencia de tipo
var nombre = "Juan"
var edad = 25

// Declaración corta (solo dentro de funciones)
nombre := "Juan"
edad := 25

// Múltiples variables
var x, y int = 10, 20
a, b := "hello", "world"
```

### Tipos Básicos

```go
// Números enteros
var entero int = 42
var entero8 int8 = 127
var entero64 int64 = 9223372036854775807

// Números flotantes
var flotante float32 = 3.14
var doble float64 = 3.141592653589793

// Booleanos
var verdadero bool = true
var falso bool = false

// Strings
var texto string = "¡Hola Go!"
var multilinea string = `Este es un
texto de múltiples
líneas`

// Caracteres (runes)
var caracter rune = 'A'
```

### Constantes

```go
// Constantes simples
const PI = 3.14159
const MENSAJE = "Bienvenido"

// Bloque de constantes
const (
    LUNES = 1
    MARTES = 2
    MIERCOLES = 3
)

// Constantes con iota (enumeración)
const (
    PEQUENO = iota  // 0
    MEDIANO         // 1
    GRANDE          // 2
)
```

## 5. Funciones

### Sintaxis Básica

```go
// Función sin parámetros ni retorno
func saludar() {
    fmt.Println("¡Hola!")
}

// Función con parámetros
func sumar(a int, b int) int {
    return a + b
}

// Parámetros del mismo tipo (forma corta)
func multiplicar(a, b int) int {
    return a * b
}

// Múltiples valores de retorno
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return a / b, nil
}

// Valores de retorno nombrados
func calculadora(a, b int) (suma, resta int) {
    suma = a + b
    resta = a - b
    return // return implícito
}
```

### Ejemplos Prácticos

```go
package main

import "fmt"

// Función que saluda a una persona
func saludarPersona(nombre string) {
    fmt.Printf("¡Hola, %s!\n", nombre)
}

// Función que calcula el área de un círculo
func areaCirculo(radio float64) float64 {
    const PI = 3.14159
    return PI * radio * radio
}

// Función que determina si un número es par
func esPar(numero int) bool {
    return numero%2 == 0
}

func main() {
    saludarPersona("Carlos")
    
    area := areaCirculo(5.0)
    fmt.Printf("El área del círculo es: %.2f\n", area)
    
    if esPar(10) {
        fmt.Println("10 es par")
    }
}
```

## 6. Control de Flujo

### if/else

```go
// if básico
if edad >= 18 {
    fmt.Println("Eres mayor de edad")
}

// if con else
if edad >= 18 {
    fmt.Println("Eres mayor de edad")
} else {
    fmt.Println("Eres menor de edad")
}

// if con múltiples condiciones
if edad >= 18 && edad < 65 {
    fmt.Println("Estás en edad laboral")
} else if edad >= 65 {
    fmt.Println("Estás jubilado")
} else {
    fmt.Println("Eres menor de edad")
}

// if con declaración inicial
if resultado := calcular(); resultado > 0 {
    fmt.Println("Resultado positivo")
}
```

### for (único tipo de bucle en Go)

```go
// Bucle clásico
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Bucle while
contador := 0
for contador < 5 {
    fmt.Println(contador)
    contador++
}

// Bucle infinito
for {
    // Código que se ejecuta infinitamente
    // Usar break para salir
    if condicion {
        break
    }
}

// Iteración sobre arrays/slices
numeros := []int{1, 2, 3, 4, 5}
for indice, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
}

// Solo el valor (ignorar índice)
for _, valor := range numeros {
    fmt.Println(valor)
}
```

### switch

```go
// Switch básico
dia := 3
switch dia {
case 1:
    fmt.Println("Lunes")
case 2:
    fmt.Println("Martes")
case 3:
    fmt.Println("Miércoles")
default:
    fmt.Println("Día desconocido")
}

// Switch con múltiples valores
switch dia {
case 1, 2, 3, 4, 5:
    fmt.Println("Día laboral")
case 6, 7:
    fmt.Println("Fin de semana")
}

// Switch sin expresión (como if-else)
switch {
case edad < 18:
    fmt.Println("Menor de edad")
case edad < 65:
    fmt.Println("Adulto")
default:
    fmt.Println("Jubilado")
}
```

## 7. Arrays y Slices

### Arrays (tamaño fijo)

```go
// Declaración
var numeros [5]int
var nombres [3]string = [3]string{"Ana", "Juan", "María"}

// Inicialización con valores
colores := [4]string{"rojo", "verde", "azul", "amarillo"}

// Tamaño automático
frutas := [...]string{"manzana", "banana", "naranja"}

// Acceso a elementos
fmt.Println(colores[0])  // "rojo"
colores[1] = "púrpura"
```

### Slices (tamaño dinámico)

```go
// Creación de slices
var numeros []int
numeros = append(numeros, 1, 2, 3)

// Slice literal
colores := []string{"rojo", "verde", "azul"}

// Crear slice con make
puntuaciones := make([]float64, 5)  // [0 0 0 0 0]
capacidades := make([]int, 3, 10)   // longitud 3, capacidad 10

// Operaciones con slices
numeros = append(numeros, 4, 5)     // Agregar elementos
subslice := numeros[1:3]            // Subslice [2, 3]
longitud := len(numeros)            // Longitud
capacidad := cap(numeros)           // Capacidad

// Iteración
for i, valor := range numeros {
    fmt.Printf("Índice %d: %d\n", i, valor)
}
```

## 8. Strings

```go
// Declaración
nombre := "Go Programming"

// Operaciones básicas
longitud := len(nombre)
primera := nombre[0]  // byte, no rune
subcadena := nombre[3:6]  // "Pro"

// Concatenación
saludo := "Hola, " + nombre + "!"

// Strings multilinea
mensaje := `Este es un
mensaje de múltiples
líneas`

// Formateo con fmt
edad := 25
texto := fmt.Sprintf("Tengo %d años", edad)
fmt.Printf("Mi nombre es %s y tengo %d años\n", nombre, edad)
```

## 9. Ejemplo Completo: Calculadora Simple

```go
package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Uso: programa <número1> <operador> <número2>")
        fmt.Println("Ejemplo: programa 5 + 3")
        return
    }
    
    // Convertir argumentos
    num1, err1 := strconv.ParseFloat(os.Args[1], 64)
    operador := os.Args[2]
    num2, err2 := strconv.ParseFloat(os.Args[3], 64)
    
    if err1 != nil || err2 != nil {
        fmt.Println("Error: Los argumentos deben ser números válidos")
        return
    }
    
    // Realizar operación
    var resultado float64
    switch operador {
    case "+":
        resultado = num1 + num2
    case "-":
        resultado = num1 - num2
    case "*":
        resultado = num1 * num2
    case "/":
        if num2 == 0 {
            fmt.Println("Error: División por cero")
            return
        }
        resultado = num1 / num2
    default:
        fmt.Println("Error: Operador no válido. Use +, -, *, /")
        return
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operador, num2, resultado)
}
```

## 10. Mejores Prácticas

1. **Nombres descriptivos**: `calcularEdad()` mejor que `calc()`
2. **Usa go fmt**: Mantén el código bien formateado
3. **Maneja errores**: Siempre verifica valores de error
4. **Usa const para valores fijos**: `const PI = 3.14159`
5. **Inicializa variables**: Evita valores zero no deseados

## 11. Herramientas Útiles

```bash
go run main.go      # Ejecutar programa
go build           # Compilar
go fmt             # Formatear código
go vet             # Análisis estático
go mod init nombre  # Inicializar módulo
```

## Ejercicios Recomendados

1. Crear un programa que calcule el área de diferentes formas geométricas
2. Implementar un conversor de unidades (metros a pies, celsius a fahrenheit)
3. Crear un programa que determine si un año es bisiesto
4. Implementar un generador de números primos

¡Practica estos conceptos con los ejercicios en la carpeta `ejercicios/`!