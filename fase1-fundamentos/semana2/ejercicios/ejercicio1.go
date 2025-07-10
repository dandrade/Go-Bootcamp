package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Ejercicio 1: Arrays y Slices ===")
	
	// 1. Crear un array de 5 números
	numbers := [5]int{23, 45, 12, 78, 34}
	fmt.Printf("Array original: %v\n", numbers)
	
	// 2. Convertir a slice y agregar elementos
	slice := numbers[:]
	slice = append(slice, 89, 56, 67)
	fmt.Printf("Slice después de agregar elementos: %v\n", slice)
	
	// 3. Encontrar el mayor y menor valor
	min, max := findMinMax(slice)
	fmt.Printf("Valor mínimo: %d\n", min)
	fmt.Printf("Valor máximo: %d\n", max)
	
	// 4. Calcular promedio
	average := calculateAverage(slice)
	fmt.Printf("Promedio: %.2f\n", average)
	
	// 5. Operaciones adicionales
	fmt.Println("\n--- Operaciones Adicionales ---")
	
	// Ordenar slice
	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Ints(sortedSlice)
	fmt.Printf("Slice ordenado: %v\n", sortedSlice)
	
	// Filtrar números pares
	evenNumbers := filterEven(slice)
	fmt.Printf("Números pares: %v\n", evenNumbers)
	
	// Crear un sub-slice
	if len(slice) >= 5 {
		subSlice := slice[2:5]
		fmt.Printf("Sub-slice [2:5]: %v\n", subSlice)
	}
	
	// Ejemplo con slice de strings
	fmt.Println("\n--- Trabajando con Strings ---")
	fruits := []string{"manzana", "banana", "naranja"}
	fruits = append(fruits, "uva", "pera")
	fmt.Printf("Frutas: %v\n", fruits)
	
	// Buscar elemento
	if contains(fruits, "banana") {
		fmt.Println("Se encontró 'banana' en la lista")
	}
}

// Función para encontrar el valor mínimo y máximo
func findMinMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	
	min := numbers[0]
	max := numbers[0]
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	
	return min, max
}

// Función para calcular el promedio
func calculateAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	
	return float64(sum) / float64(len(numbers))
}

// Función para filtrar números pares
func filterEven(numbers []int) []int {
	var even []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}

// Función para verificar si un slice contiene un elemento
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}