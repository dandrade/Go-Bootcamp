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
	// Validar que el número sea positivo o cero
	// Retornar math.Sqrt(a)
	return 0, nil
}

// TODO: Implementar función para valor absoluto
func calcularValorAbsoluto(a float64) float64 {
	// Retornar math.Abs(a)
	return 0
}

// Función auxiliar para validar números
func esNumeroValido(valor string) bool {
	_, err := strconv.ParseFloat(valor, 64)
	return err == nil
}

// Función auxiliar para operadores válidos
func esOperadorValido(operador string) bool {
	operadores := []string{"+", "-", "*", "/", "^", "%"}
	for _, op := range operadores {
		if op == operador {
			return true
		}
	}
	return false
}