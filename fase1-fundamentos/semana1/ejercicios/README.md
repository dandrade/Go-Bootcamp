# Ejercicios Prácticos - Semana 1

## Instrucciones

1. Cada ejercicio debe ejecutarse con: `go run ejercicioX.go`
2. Asegúrate de que compile sin errores
3. Usa `go fmt` para formatear tu código
4. Revisa las soluciones después de intentar cada ejercicio

## Ejercicio 1: Variables y Tipos 🔢

**Archivo**: `ejercicio1.go`

**Objetivo**: Practicar declaración de variables y tipos básicos

**Tareas**:
```go
// Completa el código
package main

import "fmt"

func main() {
    // 1. Declara una variable nombre con tu nombre
    
    // 2. Declara una variable edad con tu edad
    
    // 3. Declara una variable altura (float64) con tu altura en metros
    
    // 4. Declara una variable esEstudiante (bool) 
    
    // 5. Imprime todas las variables con fmt.Printf
    
    // 6. Crea una constante PI con valor 3.14159
    
    // 7. Calcula e imprime el área de un círculo con radio 5
}
```

## Ejercicio 2: Funciones 🔧

**Archivo**: `ejercicio2.go`

**Objetivo**: Crear y usar funciones

**Tareas**:
```go
package main

import "fmt"

// 1. Función que recibe dos enteros y retorna la suma
func sumar(a, b int) int {
    // Completa aquí
}

// 2. Función que recibe un nombre y retorna un saludo
func saludar(nombre string) string {
    // Completa aquí
}

// 3. Función que calcula el área de un rectángulo
func areaRectangulo(base, altura float64) float64 {
    // Completa aquí
}

// 4. Función que determina si un número es par
func esPar(numero int) bool {
    // Completa aquí
}

// 5. Función que retorna el mayor de dos números
func mayor(a, b int) int {
    // Completa aquí
}

func main() {
    // Usa todas las funciones aquí
}
```

## Ejercicio 3: Control de Flujo 🔄

**Archivo**: `ejercicio3.go`

**Objetivo**: Usar if/else, for, switch

**Tareas**:
```go
package main

import "fmt"

func main() {
    // 1. Programa que clasifica una edad
    edad := 25
    // Si es menor de 18: "Menor de edad"
    // Si es de 18 a 65: "Adulto"
    // Si es mayor de 65: "Jubilado"
    
    // 2. Programa que imprime números del 1 al 10
    
    // 3. Programa que imprime solo números pares del 1 al 20
    
    // 4. Switch que clasifica un día de la semana
    dia := 3
    // 1-5: "Día laboral"
    // 6-7: "Fin de semana"
    
    // 5. Programa que encuentra el primer número divisible por 7 entre 1 y 100
    
    // 6. Programa que cuenta cuántos números impares hay entre 1 y 50
}
```

## Ejercicio 4: Arrays y Slices 📋

**Archivo**: `ejercicio4.go`

**Objetivo**: Trabajar con arrays y slices

**Tareas**:
```go
package main

import "fmt"

func main() {
    // 1. Crear un array de 5 enteros con valores 1, 2, 3, 4, 5
    
    // 2. Crear un slice de strings con nombres de frutas
    
    // 3. Agregar más frutas al slice usando append
    
    // 4. Imprimir todos los elementos usando range
    
    // 5. Crear un slice con los primeros 3 elementos del array
    
    // 6. Función que encuentra el número mayor en un slice
    
    // 7. Función que calcula la suma de todos los elementos de un slice
    
    // 8. Programa que invierte un slice de enteros
}
```

## Ejercicio 5: Strings 📝

**Archivo**: `ejercicio5.go`

**Objetivo**: Manipular strings

**Tareas**:
```go
package main

import "fmt"

func main() {
    // 1. Crear una variable con tu nombre completo
    
    // 2. Imprimir la longitud del nombre
    
    // 3. Imprimir el primer y último carácter
    
    // 4. Crear un slice con tus nombres separados
    
    // 5. Función que cuenta vocales en una string
    
    // 6. Función que determina si una string es un palíndromo
    
    // 7. Función que capitaliza la primera letra de cada palabra
    
    // 8. Programa que crea iniciales de un nombre completo
}
```

## Ejercicio 6: Desafío - Validador de Contraseñas 🔒

**Archivo**: `ejercicio6.go`

**Objetivo**: Combinar todos los conceptos aprendidos

**Tareas**:
```go
package main

import "fmt"

// Función que valida si una contraseña es segura
func validarContrasena(contrasena string) bool {
    // Requisitos:
    // - Al menos 8 caracteres
    // - Al menos una mayúscula
    // - Al menos una minúscula
    // - Al menos un número
    // - Al menos un carácter especial (!@#$%^&*)
    
    // Completa aquí
}

// Función que genera un informe de la contraseña
func informeContrasena(contrasena string) {
    // Imprime:
    // - Longitud
    // - Número de mayúsculas
    // - Número de minúsculas
    // - Número de dígitos
    // - Número de caracteres especiales
    
    // Completa aquí
}

func main() {
    // Prueba con diferentes contraseñas
    contraseñas := []string{
        "123456",
        "password",
        "Password123!",
        "MiContraseña123!",
    }
    
    for _, pwd := range contraseñas {
        fmt.Printf("Contraseña: %s\n", pwd)
        if validarContrasena(pwd) {
            fmt.Println("✅ Contraseña segura")
        } else {
            fmt.Println("❌ Contraseña no segura")
        }
        informeContrasena(pwd)
        fmt.Println("---")
    }
}
```

## Ejercicio 7: Calculadora de Notas 📊

**Archivo**: `ejercicio7.go`

**Objetivo**: Crear un programa más complejo

**Tareas**:
```go
package main

import "fmt"

// Función que calcula el promedio de un slice de notas
func calcularPromedio(notas []float64) float64 {
    // Completa aquí
}

// Función que determina la letra de calificación
func obtenerLetra(promedio float64) string {
    // 90-100: A
    // 80-89: B
    // 70-79: C
    // 60-69: D
    // 0-59: F
    
    // Completa aquí
}

// Función que encuentra la nota más alta
func notaMaxima(notas []float64) float64 {
    // Completa aquí
}

// Función que encuentra la nota más baja
func notaMinima(notas []float64) float64 {
    // Completa aquí
}

// Función que cuenta cuántas notas están por encima del promedio
func notasSobrePromedio(notas []float64) int {
    // Completa aquí
}

func main() {
    // Slice de notas de ejemplo
    notas := []float64{85.5, 92.0, 78.5, 96.0, 88.0, 82.5, 90.0}
    
    // Calcular e imprimir estadísticas
    promedio := calcularPromedio(notas)
    fmt.Printf("Promedio: %.2f\n", promedio)
    fmt.Printf("Calificación: %s\n", obtenerLetra(promedio))
    fmt.Printf("Nota más alta: %.2f\n", notaMaxima(notas))
    fmt.Printf("Nota más baja: %.2f\n", notaMinima(notas))
    fmt.Printf("Notas sobre promedio: %d\n", notasSobrePromedio(notas))
}
```

## Cómo Verificar tu Progreso

1. **Compila sin errores**: `go build ejercicioX.go`
2. **Ejecuta correctamente**: `go run ejercicioX.go`
3. **Código bien formateado**: `go fmt ejercicioX.go`
4. **Resultados esperados**: Compara con las soluciones

## Soluciones

Las soluciones están en la carpeta `soluciones/` - ¡Intenta resolver cada ejercicio antes de verlas!

## Siguiente Paso

Una vez completados todos los ejercicios, continúa con el **Proyecto CLI Calculator** en la carpeta `proyecto/`.