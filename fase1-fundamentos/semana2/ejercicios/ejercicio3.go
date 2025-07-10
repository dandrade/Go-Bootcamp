package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("=== Ejercicio 3: Procesamiento de Strings ===")
	
	// Texto de ejemplo
	text := `Hola Mundo!
Esta es una prueba de procesamiento de texto.
¿Cómo estás? Espero que bien.
Vamos a contar palabras, líneas y caracteres.`
	
	fmt.Println("Texto original:")
	fmt.Println(text)
	fmt.Println()
	
	// 1. Contar caracteres, palabras y líneas
	fmt.Println("--- Estadísticas del Texto ---")
	charCount := len(text)
	wordCount := countWords(text)
	lineCount := countLines(text)
	
	fmt.Printf("Caracteres: %d\n", charCount)
	fmt.Printf("Palabras: %d\n", wordCount)
	fmt.Printf("Líneas: %d\n", lineCount)
	
	// 2. Análisis de caracteres
	fmt.Println("\n--- Análisis de Caracteres ---")
	letters, digits, spaces, punctuation := analyzeCharacters(text)
	fmt.Printf("Letras: %d\n", letters)
	fmt.Printf("Dígitos: %d\n", digits)
	fmt.Printf("Espacios: %d\n", spaces)
	fmt.Printf("Puntuación: %d\n", punctuation)
	
	// 3. Conversiones de formato
	fmt.Println("\n--- Conversiones de Formato ---")
	fmt.Printf("Mayúsculas: %s\n", strings.ToUpper(text))
	fmt.Printf("Minúsculas: %s\n", strings.ToLower(text))
	fmt.Printf("Título: %s\n", strings.Title(strings.ToLower(text)))
	
	// 4. Manipulación de texto
	fmt.Println("\n--- Manipulación de Texto ---")
	
	// Reemplazar texto
	modifiedText := strings.Replace(text, "Hola", "Buenos días", -1)
	fmt.Println("Texto modificado (Hola → Buenos días):")
	fmt.Println(modifiedText)
	
	// Limpiar espacios
	messy := "  Texto con espacios extra   "
	cleaned := strings.TrimSpace(messy)
	fmt.Printf("Original: '%s'\n", messy)
	fmt.Printf("Limpio: '%s'\n", cleaned)
	
	// 5. Dividir y unir texto
	fmt.Println("\n--- Dividir y Unir Texto ---")
	words := strings.Fields(text)
	fmt.Printf("Primeras 5 palabras: %v\n", words[:5])
	
	// Dividir por líneas
	lines := strings.Split(text, "\n")
	fmt.Printf("Número de líneas: %d\n", len(lines))
	
	// Unir palabras
	joined := strings.Join(words[:3], " - ")
	fmt.Printf("Primeras 3 palabras unidas: %s\n", joined)
	
	// 6. Búsqueda en texto
	fmt.Println("\n--- Búsqueda en Texto ---")
	
	searchWord := "texto"
	if strings.Contains(strings.ToLower(text), searchWord) {
		fmt.Printf("La palabra '%s' se encontró en el texto\n", searchWord)
	}
	
	// Buscar prefijo y sufijo
	firstLine := lines[0]
	if strings.HasPrefix(firstLine, "Hola") {
		fmt.Println("El texto comienza con 'Hola'")
	}
	
	// 7. Palíndromo
	fmt.Println("\n--- Verificar Palíndromo ---")
	testWords := []string{"radar", "hello", "level", "golang"}
	for _, word := range testWords {
		if isPalindrome(word) {
			fmt.Printf("'%s' es un palíndromo\n", word)
		} else {
			fmt.Printf("'%s' no es un palíndromo\n", word)
		}
	}
	
	// 8. Formatear texto
	fmt.Println("\n--- Formateo de Texto ---")
	name := "Juan"
	age := 25
	height := 1.75
	
	formatted := fmt.Sprintf("Nombre: %s, Edad: %d años, Altura: %.2f m", name, age, height)
	fmt.Println(formatted)
	
	// 9. Procesamiento de URL
	fmt.Println("\n--- Procesamiento de URL ---")
	url := "https://www.ejemplo.com/pagina?param=valor"
	urlInfo := analyzeURL(url)
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Protocolo: %s\n", urlInfo.Protocol)
	fmt.Printf("Dominio: %s\n", urlInfo.Domain)
	fmt.Printf("Ruta: %s\n", urlInfo.Path)
}

// Función para contar palabras
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// Función para contar líneas
func countLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}

// Función para analizar tipos de caracteres
func analyzeCharacters(text string) (int, int, int, int) {
	var letters, digits, spaces, punctuation int
	
	for _, char := range text {
		switch {
		case unicode.IsLetter(char):
			letters++
		case unicode.IsDigit(char):
			digits++
		case unicode.IsSpace(char):
			spaces++
		case unicode.IsPunct(char):
			punctuation++
		}
	}
	
	return letters, digits, spaces, punctuation
}

// Función para verificar si una palabra es palíndromo
func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	runes := []rune(word)
	
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	
	return true
}

// Estructura para información de URL
type URLInfo struct {
	Protocol string
	Domain   string
	Path     string
}

// Función para analizar una URL simple
func analyzeURL(url string) URLInfo {
	info := URLInfo{}
	
	// Separar protocolo
	if strings.Contains(url, "://") {
		parts := strings.Split(url, "://")
		info.Protocol = parts[0]
		url = parts[1]
	}
	
	// Separar dominio y ruta
	if strings.Contains(url, "/") {
		parts := strings.Split(url, "/")
		info.Domain = parts[0]
		info.Path = "/" + strings.Join(parts[1:], "/")
	} else {
		info.Domain = url
		info.Path = "/"
	}
	
	return info
}