# Proyecto: CLI Calculator 🧮

## Descripción del Proyecto

Crear una calculadora de línea de comandos que pueda realizar operaciones matemáticas básicas. Este proyecto integra todos los conceptos aprendidos en la Semana 1.

## Objetivos de Aprendizaje

- Usar argumentos de línea de comandos
- Aplicar funciones para organizar código
- Implementar control de flujo con switch
- Manejar errores básicos
- Convertir strings a números
- Formatear salida

## Funcionalidades Requeridas

### Nivel Básico ⭐
- [x] Suma de dos números
- [x] Resta de dos números
- [x] Multiplicación de dos números
- [x] División de dos números (con validación de división por cero)

### Nivel Intermedio ⭐⭐
- [ ] Potencia (a^b)
- [ ] Raíz cuadrada
- [ ] Modulo (resto de división)
- [ ] Valor absoluto

### Nivel Avanzado ⭐⭐⭐
- [ ] Operaciones múltiples (ej: 5 + 3 * 2)
- [ ] Funciones trigonométricas básicas (sin, cos, tan)
- [ ] Modo interactivo (calculadora que no termine después de una operación)

## Uso del Programa

```bash
# Operaciones básicas
go run main.go 5 + 3
go run main.go 10 - 4
go run main.go 7 * 8
go run main.go 15 / 3

# Operaciones avanzadas
go run main.go 2 ^ 8
go run main.go sqrt 16
go run main.go abs -5
```

## Estructura del Proyecto

```
proyecto/
├── main.go           # Archivo principal
├── calculator.go     # Funciones de cálculo
├── utils.go          # Funciones auxiliares
├── README.md         # Este archivo
└── tests/
    └── calculator_test.go  # Tests (opcional)
```

## Guía de Desarrollo

### Paso 1: Estructura Básica

Crea el archivo `main.go` con:
- Función main que valide argumentos
- Análisis de argumentos de línea de comandos
- Llamadas a funciones de cálculo

### Paso 2: Funciones de Cálculo

Crea `calculator.go` con:
- Función para cada operación matemática
- Validación de entrada
- Manejo de errores

### Paso 3: Utilidades

Crea `utils.go` con:
- Función para convertir strings a números
- Función para mostrar ayuda
- Función para validar argumentos

### Paso 4: Mejoras

- Agregar más operaciones
- Mejorar el manejo de errores
- Agregar tests

## Ejemplo de Implementación

### main.go (Plantilla)
```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Validar número de argumentos
    if len(os.Args) < 2 {
        mostrarAyuda()
        return
    }
    
    // Procesar diferentes tipos de operaciones
    switch os.Args[1] {
    case "help", "-h", "--help":
        mostrarAyuda()
    case "sqrt":
        // Manejar raíz cuadrada
        manejarRaizCuadrada()
    case "abs":
        // Manejar valor absoluto
        manejarValorAbsoluto()
    default:
        // Operaciones binarias (necesitan 3 argumentos)
        if len(os.Args) != 4 {
            fmt.Println("Error: Operación binaria requiere 3 argumentos")
            mostrarAyuda()
            return
        }
        manejarOperacionBinaria()
    }
}

func mostrarAyuda() {
    fmt.Println("Calculadora CLI - Uso:")
    fmt.Println("  go run main.go <número1> <operador> <número2>")
    fmt.Println("  go run main.go <operador> <número>")
    fmt.Println("")
    fmt.Println("Operadores binarios:")
    fmt.Println("  +    Suma")
    fmt.Println("  -    Resta")
    fmt.Println("  *    Multiplicación")
    fmt.Println("  /    División")
    fmt.Println("  ^    Potencia")
    fmt.Println("  %    Módulo")
    fmt.Println("")
    fmt.Println("Operadores unarios:")
    fmt.Println("  sqrt Raíz cuadrada")
    fmt.Println("  abs  Valor absoluto")
}

func manejarOperacionBinaria() {
    // Convertir argumentos
    num1, err1 := strconv.ParseFloat(os.Args[1], 64)
    operador := os.Args[2]
    num2, err2 := strconv.ParseFloat(os.Args[3], 64)
    
    if err1 != nil || err2 != nil {
        fmt.Println("Error: Argumentos deben ser números válidos")
        return
    }
    
    // Realizar operación
    resultado, err := calcularBinario(num1, operador, num2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operador, num2, resultado)
}

// TODO: Implementar manejarRaizCuadrada()
// TODO: Implementar manejarValorAbsoluto()
```

### calculator.go (Plantilla)
```go
package main

import (
    "errors"
    "math"
)

// Función para operaciones binarias
func calcularBinario(a float64, operador string, b float64) (float64, error) {
    switch operador {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, errors.New("división por cero")
        }
        return a / b, nil
    case "^":
        return math.Pow(a, b), nil
    case "%":
        if b == 0 {
            return 0, errors.New("módulo por cero")
        }
        return math.Mod(a, b), nil
    default:
        return 0, errors.New("operador no válido")
    }
}

// TODO: Implementar función para raíz cuadrada
func calcularRaizCuadrada(a float64) (float64, error) {
    // Validar que el número sea positivo
    // Retornar math.Sqrt(a)
    return 0, nil
}

// TODO: Implementar función para valor absoluto
func calcularValorAbsoluto(a float64) float64 {
    // Retornar math.Abs(a)
    return 0
}
```

## Tests de Ejemplo

```go
package main

import "testing"

func TestCalcularBinario(t *testing.T) {
    // Test suma
    resultado, err := calcularBinario(5, "+", 3)
    if err != nil || resultado != 8 {
        t.Errorf("5 + 3 = 8, pero obtuve %f", resultado)
    }
    
    // Test división por cero
    _, err = calcularBinario(5, "/", 0)
    if err == nil {
        t.Error("División por cero debería retornar error")
    }
}
```

## Criterios de Evaluación

### Funcionalidad (50%)
- ✅ Operaciones básicas funcionan correctamente
- ✅ Manejo de errores apropiado
- ✅ Validación de entrada

### Código (30%)
- ✅ Código bien estructurado y legible
- ✅ Funciones bien definidas
- ✅ Comentarios apropiados

### Buenas Prácticas (20%)
- ✅ Uso de `go fmt`
- ✅ Manejo de errores idiomático
- ✅ Nombres de variables descriptivos

## Entregables

1. **main.go** - Archivo principal con lógica de entrada
2. **calculator.go** - Funciones de cálculo
3. **utils.go** - Funciones auxiliares (opcional)
4. **README.md** - Documentación del proyecto

## Extensiones Opcionales

1. **Calculadora Científica**: sin, cos, tan, log, ln
2. **Modo Interactivo**: No termine después de una operación
3. **Historial**: Guarde las operaciones anteriores
4. **Archivo de Configuración**: Precisión decimal, formato de salida
5. **Tests Completos**: Cobertura de todas las funciones

## Recursos Útiles

- [Package os](https://pkg.go.dev/os) - Para argumentos de línea de comandos
- [Package strconv](https://pkg.go.dev/strconv) - Para conversión de strings
- [Package math](https://pkg.go.dev/math) - Para operaciones matemáticas
- [Package errors](https://pkg.go.dev/errors) - Para manejo de errores

¡Comienza con el nivel básico y ve agregando funcionalidades gradualmente! 🚀