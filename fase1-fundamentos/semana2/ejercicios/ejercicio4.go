package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== Ejercicio 4: Contador de Palabras ===")
	
	// Texto de ejemplo
	text := `Go es un lenguaje de programación desarrollado por Google.
Go es simple, rápido y eficiente.
Los programas en Go se compilan rápidamente.
Go tiene una excelente gestión de concurrencia.
Muchas empresas usan Go para desarrollo backend.
Go es perfecto para microservicios y APIs.`
	
	fmt.Println("Texto a analizar:")
	fmt.Println(text)
	fmt.Println()
	
	// 1. Contar frecuencia de palabras
	fmt.Println("--- Análisis de Frecuencia de Palabras ---")
	wordCount := countWordFrequency(text)
	
	fmt.Printf("Total de palabras únicas: %d\n", len(wordCount))
	fmt.Printf("Total de palabras: %d\n", getTotalWords(wordCount))
	
	// 2. Mostrar todas las palabras y sus frecuencias
	fmt.Println("\n--- Todas las Palabras ---")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
	
	// 3. Encontrar las palabras más comunes
	fmt.Println("\n--- Top 5 Palabras Más Comunes ---")
	topWords := getTopWords(wordCount, 5)
	for i, wordFreq := range topWords {
		fmt.Printf("%d. %s: %d veces\n", i+1, wordFreq.Word, wordFreq.Count)
	}
	
	// 4. Estadísticas adicionales
	fmt.Println("\n--- Estadísticas Adicionales ---")
	
	// Palabras que aparecen solo una vez
	uniqueWords := getUniqueWords(wordCount)
	fmt.Printf("Palabras que aparecen solo una vez: %d\n", len(uniqueWords))
	if len(uniqueWords) > 0 {
		fmt.Printf("Ejemplos: %v\n", uniqueWords[:min(5, len(uniqueWords))])
	}
	
	// Longitud promedio de palabras
	avgLength := getAverageWordLength(wordCount)
	fmt.Printf("Longitud promedio de palabras: %.2f caracteres\n", avgLength)
	
	// Palabras más largas
	longestWords := getLongestWords(wordCount, 3)
	fmt.Printf("Palabras más largas: %v\n", longestWords)
	
	// 5. Análisis por longitud
	fmt.Println("\n--- Análisis por Longitud de Palabras ---")
	lengthDistribution := analyzeLengthDistribution(wordCount)
	for length, count := range lengthDistribution {
		fmt.Printf("Palabras de %d caracteres: %d\n", length, count)
	}
	
	// 6. Buscar palabras específicas
	fmt.Println("\n--- Búsqueda de Palabras ---")
	searchWords := []string{"go", "programación", "google", "rápido"}
	
	for _, word := range searchWords {
		if count, exists := wordCount[word]; exists {
			fmt.Printf("'%s' aparece %d veces\n", word, count)
		} else {
			fmt.Printf("'%s' no se encuentra en el texto\n", word)
		}
	}
	
	// 7. Generar reporte
	fmt.Println("\n--- Reporte de Análisis ---")
	generateReport(text, wordCount)
	
	// 8. Comparar dos textos
	fmt.Println("\n--- Comparación de Textos ---")
	text2 := `Python es un lenguaje de programación popular.
Python es fácil de aprender y usar.
Muchas empresas usan Python para ciencia de datos.`
	
	wordCount2 := countWordFrequency(text2)
	commonWords := findCommonWords(wordCount, wordCount2)
	fmt.Printf("Palabras en común entre los textos: %v\n", commonWords)
}

// Estructura para almacenar palabra y frecuencia
type WordFrequency struct {
	Word  string
	Count int
}

// Función para contar frecuencia de palabras
func countWordFrequency(text string) map[string]int {
	// Convertir a minúsculas y dividir en palabras
	words := strings.Fields(strings.ToLower(text))
	wordCount := make(map[string]int)
	
	for _, word := range words {
		// Limpiar signos de puntuación
		cleaned := cleanWord(word)
		if cleaned != "" {
			wordCount[cleaned]++
		}
	}
	
	return wordCount
}

// Función para limpiar una palabra de signos de puntuación
func cleanWord(word string) string {
	// Eliminar signos de puntuación al inicio y final
	word = strings.Trim(word, ".,!?;:\"'()[]{}*")
	return word
}

// Función para obtener el total de palabras
func getTotalWords(wordCount map[string]int) int {
	total := 0
	for _, count := range wordCount {
		total += count
	}
	return total
}

// Función para obtener las palabras más comunes
func getTopWords(wordCount map[string]int, n int) []WordFrequency {
	var wordFreqs []WordFrequency
	
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, WordFrequency{Word: word, Count: count})
	}
	
	// Ordenar por frecuencia (descendente)
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Count > wordFreqs[j].Count
	})
	
	// Retornar solo los primeros n
	if n > len(wordFreqs) {
		n = len(wordFreqs)
	}
	
	return wordFreqs[:n]
}

// Función para obtener palabras únicas (que aparecen solo una vez)
func getUniqueWords(wordCount map[string]int) []string {
	var uniqueWords []string
	
	for word, count := range wordCount {
		if count == 1 {
			uniqueWords = append(uniqueWords, word)
		}
	}
	
	return uniqueWords
}

// Función para calcular longitud promedio de palabras
func getAverageWordLength(wordCount map[string]int) float64 {
	totalLength := 0
	totalCount := 0
	
	for word, count := range wordCount {
		totalLength += len(word) * count
		totalCount += count
	}
	
	if totalCount == 0 {
		return 0
	}
	
	return float64(totalLength) / float64(totalCount)
}

// Función para obtener las palabras más largas
func getLongestWords(wordCount map[string]int, n int) []string {
	var words []string
	
	for word := range wordCount {
		words = append(words, word)
	}
	
	// Ordenar por longitud (descendente)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	
	// Retornar solo los primeros n
	if n > len(words) {
		n = len(words)
	}
	
	return words[:n]
}

// Función para analizar distribución por longitud
func analyzeLengthDistribution(wordCount map[string]int) map[int]int {
	lengthDist := make(map[int]int)
	
	for word, count := range wordCount {
		length := len(word)
		lengthDist[length] += count
	}
	
	return lengthDist
}

// Función para generar reporte
func generateReport(text string, wordCount map[string]int) {
	totalWords := getTotalWords(wordCount)
	uniqueWords := len(wordCount)
	avgLength := getAverageWordLength(wordCount)
	
	fmt.Printf("Caracteres totales: %d\n", len(text))
	fmt.Printf("Palabras totales: %d\n", totalWords)
	fmt.Printf("Palabras únicas: %d\n", uniqueWords)
	fmt.Printf("Diversidad léxica: %.2f%%\n", float64(uniqueWords)/float64(totalWords)*100)
	fmt.Printf("Longitud promedio: %.2f caracteres\n", avgLength)
}

// Función para encontrar palabras comunes entre dos textos
func findCommonWords(wordCount1, wordCount2 map[string]int) []string {
	var commonWords []string
	
	for word := range wordCount1 {
		if _, exists := wordCount2[word]; exists {
			commonWords = append(commonWords, word)
		}
	}
	
	return commonWords
}

// Función auxiliar para obtener el mínimo de dos números
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}