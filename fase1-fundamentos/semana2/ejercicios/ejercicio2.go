package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Ejercicio 2: Trabajo con Maps ===")
	
	// 1. Crear un map de estudiantes y calificaciones
	students := map[string]float64{
		"Juan":    85.5,
		"Ana":     92.0,
		"Pedro":   78.5,
		"Maria":   88.0,
		"Carlos":  95.5,
	}
	
	fmt.Println("Calificaciones iniciales:")
	printStudents(students)
	
	// 2. Agregar nuevos estudiantes
	students["Sofia"] = 89.5
	students["Diego"] = 91.0
	fmt.Println("\nDespués de agregar nuevos estudiantes:")
	printStudents(students)
	
	// 3. Modificar calificaciones existentes
	students["Juan"] = 87.0
	fmt.Println("\nDespués de modificar calificación de Juan:")
	printStudents(students)
	
	// 4. Eliminar un estudiante
	delete(students, "Pedro")
	fmt.Println("\nDespués de eliminar a Pedro:")
	printStudents(students)
	
	// 5. Calcular estadísticas
	fmt.Println("\n--- Estadísticas ---")
	avg := calculateClassAverage(students)
	fmt.Printf("Promedio de la clase: %.2f\n", avg)
	
	highest, lowest := findHighestLowest(students)
	fmt.Printf("Calificación más alta: %.2f\n", highest)
	fmt.Printf("Calificación más baja: %.2f\n", lowest)
	
	// 6. Buscar estudiantes
	fmt.Println("\n--- Búsquedas ---")
	
	// Buscar por nombre
	if grade, exists := students["Ana"]; exists {
		fmt.Printf("Ana tiene una calificación de: %.2f\n", grade)
	} else {
		fmt.Println("Ana no está en la lista")
	}
	
	// Buscar estudiantes aprobados (>= 80)
	passed := findPassedStudents(students, 80.0)
	fmt.Printf("Estudiantes aprobados: %v\n", passed)
	
	// 7. Ordenar estudiantes por calificación
	fmt.Println("\n--- Ranking ---")
	ranking := sortStudentsByGrade(students)
	for i, student := range ranking {
		fmt.Printf("%d. %s: %.2f\n", i+1, student.Name, student.Grade)
	}
	
	// 8. Ejemplo con map de contadores
	fmt.Println("\n--- Contador de Letras ---")
	text := "Hola Mundo"
	letterCount := countLetters(text)
	fmt.Printf("Conteo de letras en '%s':\n", text)
	for letter, count := range letterCount {
		fmt.Printf("%c: %d\n", letter, count)
	}
}

// Función para imprimir todos los estudiantes
func printStudents(students map[string]float64) {
	for name, grade := range students {
		fmt.Printf("%s: %.2f\n", name, grade)
	}
}

// Función para calcular el promedio de la clase
func calculateClassAverage(students map[string]float64) float64 {
	if len(students) == 0 {
		return 0
	}
	
	total := 0.0
	for _, grade := range students {
		total += grade
	}
	
	return total / float64(len(students))
}

// Función para encontrar la calificación más alta y más baja
func findHighestLowest(students map[string]float64) (float64, float64) {
	if len(students) == 0 {
		return 0, 0
	}
	
	first := true
	var highest, lowest float64
	
	for _, grade := range students {
		if first {
			highest = grade
			lowest = grade
			first = false
		} else {
			if grade > highest {
				highest = grade
			}
			if grade < lowest {
				lowest = grade
			}
		}
	}
	
	return highest, lowest
}

// Función para encontrar estudiantes aprobados
func findPassedStudents(students map[string]float64, passingGrade float64) []string {
	var passed []string
	for name, grade := range students {
		if grade >= passingGrade {
			passed = append(passed, name)
		}
	}
	return passed
}

// Estructura para ordenar estudiantes
type Student struct {
	Name  string
	Grade float64
}

// Función para ordenar estudiantes por calificación
func sortStudentsByGrade(students map[string]float64) []Student {
	var studentList []Student
	
	for name, grade := range students {
		studentList = append(studentList, Student{Name: name, Grade: grade})
	}
	
	// Ordenar por calificación (descendente)
	sort.Slice(studentList, func(i, j int) bool {
		return studentList[i].Grade > studentList[j].Grade
	})
	
	return studentList
}

// Función para contar letras en un texto
func countLetters(text string) map[rune]int {
	letterCount := make(map[rune]int)
	
	for _, char := range text {
		if char != ' ' {
			letterCount[char]++
		}
	}
	
	return letterCount
}