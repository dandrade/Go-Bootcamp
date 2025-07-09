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
	fmt.Println("")
	fmt.Println("Ejemplos:")
	fmt.Println("  go run main.go 5 + 3")
	fmt.Println("  go run main.go 10 / 2")
	fmt.Println("  go run main.go sqrt 16")
	fmt.Println("  go run main.go abs -5")
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

func manejarRaizCuadrada() {
	// TODO: Implementar manejo de raíz cuadrada
	// Verificar que haya exactamente 2 argumentos
	// Convertir el argumento a float64
	// Llamar a calcularRaizCuadrada()
	// Mostrar resultado
	fmt.Println("TODO: Implementar raíz cuadrada")
}

func manejarValorAbsoluto() {
	// TODO: Implementar manejo de valor absoluto
	// Verificar que haya exactamente 2 argumentos
	// Convertir el argumento a float64
	// Llamar a calcularValorAbsoluto()
	// Mostrar resultado
	fmt.Println("TODO: Implementar valor absoluto")
}