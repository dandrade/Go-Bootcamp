package main

import (
	"fmt"
	"time"
)

// TODO 1: Define un struct 'Book' con los siguientes campos:
// - ISBN (string)
// - Title (string)
// - Author (string)
// - Pages (int)
// - PublishedYear (int)
// - Available (bool)
type Book struct {
	// TODO: Implementar campos
}

// TODO 2: Define un struct 'Library' con:
// - Name (string)
// - Address (string)
// - Books (slice de Book)
// - OpenTime (string, ej: "09:00")
// - CloseTime (string, ej: "18:00")
type Library struct {
	// TODO: Implementar campos
}

// TODO 3: Crea una función NewBook que actúe como constructor
// Debe recibir todos los parámetros necesarios y retornar un Book
// El campo Available debe inicializarse como true por defecto
func NewBook(isbn, title, author string, pages, year int) Book {
	// TODO: Implementar constructor
	return Book{}
}

// TODO 4: Crea una función NewLibrary como constructor
// Debe inicializar el slice de Books vacío
func NewLibrary(name, address, openTime, closeTime string) Library {
	// TODO: Implementar constructor
	return Library{}
}

// TODO 5: Implementa una función que agregue un libro a la biblioteca
// Debe recibir un puntero a Library y un Book
func AddBook(lib *Library, book Book) {
	// TODO: Agregar el libro al slice de Books de la biblioteca
}

// TODO 6: Implementa una función que cuente los libros disponibles
// Debe recibir Library (por valor) y retornar el número de libros con Available = true
func CountAvailableBooks(lib Library) int {
	// TODO: Contar y retornar libros disponibles
	return 0
}

// TODO 7: Implementa una función que busque un libro por ISBN
// Debe retornar un puntero al Book si lo encuentra, o nil si no existe
func FindBookByISBN(lib Library, isbn string) *Book {
	// TODO: Buscar libro por ISBN
	return nil
}

// TODO 8: Implementa una función que marque un libro como no disponible
// Debe recibir un puntero a Library y el ISBN del libro
// Retorna true si se pudo marcar como no disponible, false si no se encontró
func CheckoutBook(lib *Library, isbn string) bool {
	// TODO: Buscar el libro y marcarlo como no disponible
	return false
}

// TODO 9: Implementa una función que retorne todos los libros de un autor
// Debe retornar un slice con todos los libros del autor especificado
func GetBooksByAuthor(lib Library, author string) []Book {
	// TODO: Filtrar y retornar libros por autor
	return nil
}

// TODO 10: Define un struct 'Member' para representar miembros de la biblioteca
// - ID (string)
// - Name (string)
// - Email (string)
// - JoinDate (time.Time)
// - BorrowedBooks (slice de strings con ISBNs)
type Member struct {
	// TODO: Implementar campos
}

func main() {
	fmt.Println("=== Sistema de Biblioteca ===\n")

	// Crear biblioteca
	library := NewLibrary(
		"Biblioteca Central",
		"Calle Principal 123",
		"09:00",
		"20:00",
	)

	// Crear y agregar libros
	book1 := NewBook("978-0134190440", "The Go Programming Language", "Alan Donovan", 380, 2015)
	book2 := NewBook("978-1491941959", "Go in Practice", "Matt Butcher", 312, 2016)
	book3 := NewBook("978-1484256664", "Pro Go", "Adam Freeman", 1076, 2021)
	book4 := NewBook("978-0134190501", "Another Go Book", "Alan Donovan", 250, 2020)

	AddBook(&library, book1)
	AddBook(&library, book2)
	AddBook(&library, book3)
	AddBook(&library, book4)

	// Mostrar información de la biblioteca
	fmt.Printf("Biblioteca: %s\n", library.Name)
	fmt.Printf("Dirección: %s\n", library.Address)
	fmt.Printf("Horario: %s - %s\n", library.OpenTime, library.CloseTime)
	fmt.Printf("Total de libros: %d\n", len(library.Books))
	fmt.Printf("Libros disponibles: %d\n\n", CountAvailableBooks(library))

	// Buscar un libro específico
	isbn := "978-0134190440"
	if book := FindBookByISBN(library, isbn); book != nil {
		fmt.Printf("Libro encontrado: %s por %s\n", book.Title, book.Author)
	}

	// Prestar un libro
	fmt.Println("\n--- Prestando libro ---")
	if CheckoutBook(&library, isbn) {
		fmt.Printf("Libro con ISBN %s prestado exitosamente\n", isbn)
		fmt.Printf("Libros disponibles ahora: %d\n", CountAvailableBooks(library))
	}

	// Buscar libros por autor
	fmt.Println("\n--- Libros de Alan Donovan ---")
	donovanBooks := GetBooksByAuthor(library, "Alan Donovan")
	for _, book := range donovanBooks {
		status := "Disponible"
		if !book.Available {
			status = "Prestado"
		}
		fmt.Printf("- %s (%d) - %s\n", book.Title, book.PublishedYear, status)
	}

	// Crear un miembro (si implementaste el TODO 10)
	member := Member{
		ID:            "M001",
		Name:          "Juan Pérez",
		Email:         "juan@email.com",
		JoinDate:      time.Now(),
		BorrowedBooks: []string{isbn},
	}
	fmt.Printf("\n--- Nuevo miembro ---\n")
	fmt.Printf("Nombre: %s\n", member.Name)
	fmt.Printf("Email: %s\n", member.Email)
	fmt.Printf("Libros prestados: %d\n", len(member.BorrowedBooks))
}