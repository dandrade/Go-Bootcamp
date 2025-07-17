package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TODO 1: Define ValidationError como struct que implemente error
type ValidationError struct {
	// TODO: Agregar campos Field, Message, Value
}

// TODO 2: Implementa Error() para ValidationError
func (e ValidationError) Error() string {
	// TODO: Retornar mensaje formateado
	return ""
}

// TODO 3: Define MultiValidationError para múltiples errores
type MultiValidationError struct {
	// TODO: Agregar slice de errores
}

// TODO 4: Implementa Error() para MultiValidationError
func (e MultiValidationError) Error() string {
	// TODO: Combinar todos los errores en un mensaje
	return ""
}

// TODO 5: Implementa AddError para MultiValidationError
func (e *MultiValidationError) AddError(err error) {
	// TODO: Agregar error al slice
}

// TODO 6: Implementa HasErrors para MultiValidationError
func (e MultiValidationError) HasErrors() bool {
	// TODO: Verificar si hay errores
	return false
}

// TODO 7: Define DatabaseError con información específica
type DatabaseError struct {
	// TODO: Agregar Operation, Table, Cause
}

// TODO 8: Implementa Error() para DatabaseError
func (e DatabaseError) Error() string {
	// TODO: Mensaje con contexto de base de datos
	return ""
}

// TODO 9: Implementa Unwrap() para DatabaseError
func (e DatabaseError) Unwrap() error {
	// TODO: Retornar el error causa
	return nil
}

// TODO 10: Define NetworkError con timeout y retry info
type NetworkError struct {
	// TODO: Agregar URL, Timeout, Retries, Cause
}

// TODO 11: Implementa Error() para NetworkError
func (e NetworkError) Error() string {
	// TODO: Mensaje con info de red
	return ""
}

// TODO 12: Implementa Unwrap() para NetworkError
func (e NetworkError) Unwrap() error {
	// TODO: Retornar error causa
	return nil
}

// TODO 13: Implementa IsTimeout() para verificar si es timeout
func (e NetworkError) IsTimeout() bool {
	// TODO: Verificar si el error es por timeout
	return false
}

// User representa un usuario del sistema
type User struct {
	ID       string
	Username string
	Email    string
	Age      int
	Password string
	CreatedAt time.Time
}

// TODO 14: Implementa ValidateUser que retorne ValidationError específicos
func ValidateUser(user User) error {
	// TODO: Validar todos los campos del usuario
	// - ID no vacío
	// - Username entre 3-20 caracteres, solo alfanumérico
	// - Email formato válido
	// - Age entre 13-120
	// - Password al menos 8 caracteres
	// Usar MultiValidationError para acumular errores
	return nil
}

// TODO 15: Implementa validateUsername
func validateUsername(username string) error {
	// TODO: Validar username (3-20 chars, alfanumérico)
	return nil
}

// TODO 16: Implementa validateEmail
func validateEmail(email string) error {
	// TODO: Validar formato de email básico
	return nil
}

// TODO 17: Implementa validateAge
func validateAge(age int) error {
	// TODO: Validar edad entre 13-120
	return nil
}

// TODO 18: Implementa validatePassword
func validatePassword(password string) error {
	// TODO: Validar password (mínimo 8 caracteres)
	return nil
}

// UserRepository simula operaciones de base de datos
type UserRepository struct {
	users map[string]User
}

// TODO 19: Implementa NewUserRepository
func NewUserRepository() *UserRepository {
	// TODO: Crear nuevo repositorio
	return nil
}

// TODO 20: Implementa Create que retorne DatabaseError en caso de error
func (r *UserRepository) Create(user User) error {
	// TODO: Crear usuario, validar que no exista
	// Retornar DatabaseError si ya existe
	return nil
}

// TODO 21: Implementa GetByID que retorne DatabaseError si no existe
func (r *UserRepository) GetByID(id string) (*User, error) {
	// TODO: Buscar usuario por ID
	// Retornar DatabaseError si no existe
	return nil, nil
}

// TODO 22: Implementa Update que valide existencia
func (r *UserRepository) Update(user User) error {
	// TODO: Actualizar usuario existente
	// Retornar DatabaseError si no existe
	return nil
}

// TODO 23: Implementa Delete que valide existencia
func (r *UserRepository) Delete(id string) error {
	// TODO: Eliminar usuario
	// Retornar DatabaseError si no existe
	return nil
}

// UserService maneja lógica de negocio
type UserService struct {
	repo *UserRepository
}

// TODO 24: Implementa NewUserService
func NewUserService(repo *UserRepository) *UserService {
	// TODO: Crear nuevo service
	return nil
}

// TODO 25: Implementa RegisterUser que valide y cree usuario
func (s *UserService) RegisterUser(username, email, password string, age int) (*User, error) {
	// TODO: Validar datos, crear usuario, guardarlo
	// Manejar tanto ValidationError como DatabaseError
	return nil, nil
}

// TODO 26: Implementa UpdateUser que valide y actualice
func (s *UserService) UpdateUser(id, username, email string, age int) (*User, error) {
	// TODO: Buscar usuario, validar nuevos datos, actualizar
	return nil, nil
}

// NetworkClient simula cliente de red
type NetworkClient struct {
	baseURL string
	timeout time.Duration
}

// TODO 27: Implementa NewNetworkClient
func NewNetworkClient(baseURL string, timeout time.Duration) *NetworkClient {
	// TODO: Crear cliente
	return nil
}

// TODO 28: Implementa SendRequest que simule error de red
func (c *NetworkClient) SendRequest(endpoint string, data []byte) ([]byte, error) {
	// TODO: Simular request que puede fallar
	// Simular diferentes tipos de errores de red
	return nil, nil
}

// RetryableOperation maneja operaciones con retry
type RetryableOperation struct {
	maxRetries int
	delay      time.Duration
}

// TODO 29: Implementa NewRetryableOperation
func NewRetryableOperation(maxRetries int, delay time.Duration) *RetryableOperation {
	// TODO: Crear operación con retry
	return nil
}

// TODO 30: Implementa Execute que reintente en caso de error
func (r *RetryableOperation) Execute(operation func() error) error {
	// TODO: Ejecutar operación con reintentos
	// Solo reintentar en NetworkError que no sea timeout
	return nil
}

// ErrorReporter recolecta errores para análisis
type ErrorReporter struct {
	errors []error
}

// TODO 31: Implementa NewErrorReporter
func NewErrorReporter() *ErrorReporter {
	// TODO: Crear reporter
	return nil
}

// TODO 32: Implementa ReportError
func (r *ErrorReporter) ReportError(err error) {
	// TODO: Agregar error a la lista
}

// TODO 33: Implementa GetErrorSummary que categorice errores
func (r *ErrorReporter) GetErrorSummary() map[string]int {
	// TODO: Contar errores por tipo
	// Usar type assertions para clasificar
	return nil
}

// TODO 34: Implementa GetValidationErrors que extraiga solo errores de validación
func (r *ErrorReporter) GetValidationErrors() []ValidationError {
	// TODO: Filtrar y retornar solo ValidationError
	return nil
}

func main() {
	fmt.Println("=== Sistema de Manejo de Errores ===\n")

	// Test de validación de usuarios
	fmt.Println("--- Validación de Usuarios ---")
	
	// Usuario válido
	validUser := User{
		ID:       "user1",
		Username: "johnDoe",
		Email:    "john@example.com",
		Age:      25,
		Password: "securePassword123",
		CreatedAt: time.Now(),
	}
	
	if err := ValidateUser(validUser); err != nil {
		fmt.Printf("Error inesperado: %v\n", err)
	} else {
		fmt.Println("Usuario válido ✓")
	}
	
	// Usuario inválido
	invalidUser := User{
		ID:       "",
		Username: "ab", // Muy corto
		Email:    "invalid-email",
		Age:      12, // Muy joven
		Password: "123", // Muy corta
	}
	
	if err := ValidateUser(invalidUser); err != nil {
		fmt.Printf("Errores de validación encontrados:\n%v\n", err)
		
		// Test de tipo específico
		if multiErr, ok := err.(MultiValidationError); ok {
			fmt.Printf("Número total de errores: %d\n", len(multiErr.Errors))
		}
	}

	// Test de repositorio
	fmt.Println("\n--- Repositorio de Usuarios ---")
	repo := NewUserRepository()
	
	// Crear usuario
	err := repo.Create(validUser)
	if err != nil {
		fmt.Printf("Error creando usuario: %v\n", err)
	} else {
		fmt.Println("Usuario creado exitosamente")
	}
	
	// Intentar crear usuario duplicado
	err = repo.Create(validUser)
	if err != nil {
		fmt.Printf("Error esperado (usuario duplicado): %v\n", err)
		
		// Verificar tipo de error
		if dbErr, ok := err.(DatabaseError); ok {
			fmt.Printf("Operación de BD: %s, Tabla: %s\n", dbErr.Operation, dbErr.Table)
		}
	}
	
	// Buscar usuario existente
	user, err := repo.GetByID("user1")
	if err != nil {
		fmt.Printf("Error buscando usuario: %v\n", err)
	} else {
		fmt.Printf("Usuario encontrado: %s\n", user.Username)
	}
	
	// Buscar usuario inexistente
	_, err = repo.GetByID("nonexistent")
	if err != nil {
		fmt.Printf("Error esperado (usuario no encontrado): %v\n", err)
	}

	// Test de service
	fmt.Println("\n--- Service de Usuarios ---")
	service := NewUserService(repo)
	
	newUser, err := service.RegisterUser("alice", "alice@example.com", "password123", 30)
	if err != nil {
		fmt.Printf("Error registrando usuario: %v\n", err)
	} else {
		fmt.Printf("Usuario registrado: %s\n", newUser.Username)
	}
	
	// Intentar registrar con datos inválidos
	_, err = service.RegisterUser("", "invalid", "123", 10)
	if err != nil {
		fmt.Printf("Error esperado (datos inválidos): %v\n", err)
	}

	// Test de cliente de red
	fmt.Println("\n--- Cliente de Red ---")
	client := NewNetworkClient("https://api.example.com", 5*time.Second)
	
	_, err = client.SendRequest("/users", []byte(`{"name": "test"}`))
	if err != nil {
		fmt.Printf("Error de red: %v\n", err)
		
		// Verificar tipo de error
		if netErr, ok := err.(NetworkError); ok {
			fmt.Printf("URL: %s, Reintentos: %d, ¿Timeout?: %v\n", 
				netErr.URL, netErr.Retries, netErr.IsTimeout())
		}
	}

	// Test de operaciones con retry
	fmt.Println("\n--- Operaciones con Retry ---")
	retryOp := NewRetryableOperation(3, 100*time.Millisecond)
	
	// Operación que siempre falla
	err = retryOp.Execute(func() error {
		return NetworkError{
			URL:     "https://unreachable.com",
			Timeout: false,
			Retries: 0,
			Cause:   fmt.Errorf("connection refused"),
		}
	})
	
	if err != nil {
		fmt.Printf("Operación falló después de reintentos: %v\n", err)
	}

	// Test de error reporter
	fmt.Println("\n--- Reporte de Errores ---")
	reporter := NewErrorReporter()
	
	// Simular varios errores
	reporter.ReportError(ValidationError{Field: "email", Message: "invalid format"})
	reporter.ReportError(DatabaseError{Operation: "INSERT", Table: "users"})
	reporter.ReportError(NetworkError{URL: "api.com", Timeout: true})
	reporter.ReportError(ValidationError{Field: "age", Message: "too young"})
	
	summary := reporter.GetErrorSummary()
	fmt.Println("Resumen de errores:")
	for errorType, count := range summary {
		fmt.Printf("  %s: %d\n", errorType, count)
	}
	
	validationErrors := reporter.GetValidationErrors()
	fmt.Printf("\nErrores de validación encontrados: %d\n", len(validationErrors))
	for _, valErr := range validationErrors {
		fmt.Printf("  - %s: %s\n", valErr.Field, valErr.Message)
	}

	// Demostración de error wrapping
	fmt.Println("\n--- Error Wrapping ---")
	originalErr := fmt.Errorf("connection timeout")
	wrappedErr := NetworkError{
		URL:     "api.example.com",
		Timeout: true,
		Cause:   originalErr,
	}
	
	fmt.Printf("Error original: %v\n", originalErr)
	fmt.Printf("Error wrapeado: %v\n", wrappedErr)
	fmt.Printf("Error unwrapeado: %v\n", wrappedErr.Unwrap())
}