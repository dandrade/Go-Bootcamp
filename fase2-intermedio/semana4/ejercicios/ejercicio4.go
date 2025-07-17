package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO 1: Define CircuitBreakerError
type CircuitBreakerError struct {
	// TODO: Agregar State, LastError, OpenedAt
}

// TODO 2: Implementa Error() para CircuitBreakerError
func (e CircuitBreakerError) Error() string {
	// TODO: Mensaje describiendo el estado del circuit breaker
	return ""
}

// TODO 3: Define RetryableError interface
type RetryableError interface {
	error
	IsRetryable() bool
	GetRetryAfter() time.Duration
}

// TODO 4: Define TemporaryError que implemente RetryableError
type TemporaryError struct {
	// TODO: Agregar Message, RetryAfter
}

// TODO 5: Implementa métodos para TemporaryError
func (e TemporaryError) Error() string {
	// TODO: Retornar mensaje
	return ""
}

func (e TemporaryError) IsRetryable() bool {
	// TODO: Siempre retryable
	return true
}

func (e TemporaryError) GetRetryAfter() time.Duration {
	// TODO: Retornar RetryAfter
	return 0
}

// TODO 6: Define PermanentError que implemente RetryableError
type PermanentError struct {
	// TODO: Agregar Message
}

// TODO 7: Implementa métodos para PermanentError
func (e PermanentError) Error() string {
	// TODO: Retornar mensaje
	return ""
}

func (e PermanentError) IsRetryable() bool {
	// TODO: Nunca retryable
	return false
}

func (e PermanentError) GetRetryAfter() time.Duration {
	// TODO: Cero porque no es retryable
	return 0
}

// TODO 8: Define ErrorAccumulator para coleccionar errores
type ErrorAccumulator struct {
	// TODO: Agregar errors slice y mutex para thread safety
}

// TODO 9: Implementa NewErrorAccumulator
func NewErrorAccumulator() *ErrorAccumulator {
	// TODO: Crear nuevo acumulador
	return nil
}

// TODO 10: Implementa AddError thread-safe
func (ea *ErrorAccumulator) AddError(err error) {
	// TODO: Agregar error de forma thread-safe
}

// TODO 11: Implementa GetErrors que retorne copia de errores
func (ea *ErrorAccumulator) GetErrors() []error {
	// TODO: Retornar copia thread-safe
	return nil
}

// TODO 12: Implementa HasErrors
func (ea *ErrorAccumulator) HasErrors() bool {
	// TODO: Verificar si hay errores
	return false
}

// TODO 13: Implementa Clear que limpie errores
func (ea *ErrorAccumulator) Clear() {
	// TODO: Limpiar errores de forma thread-safe
}

// TODO 14: Define OperationResult para resultados async
type OperationResult struct {
	// TODO: Agregar Value, Error, Duration
}

// AsyncOperation representa operación asíncrona
type AsyncOperation struct {
	name string
	fn   func() (interface{}, error)
}

// TODO 15: Implementa NewAsyncOperation
func NewAsyncOperation(name string, fn func() (interface{}, error)) *AsyncOperation {
	// TODO: Crear nueva operación
	return nil
}

// TODO 16: Implementa Execute que ejecute async y retorne channel
func (ao *AsyncOperation) Execute() <-chan OperationResult {
	// TODO: Ejecutar función en goroutine y retornar resultado via channel
	// Medir tiempo de ejecución
	return nil
}

// TODO 17: Define OperationExecutor para ejecutar múltiples operaciones
type OperationExecutor struct {
	// TODO: Agregar operations slice
}

// TODO 18: Implementa NewOperationExecutor
func NewOperationExecutor() *OperationExecutor {
	// TODO: Crear nuevo executor
	return nil
}

// TODO 19: Implementa AddOperation
func (oe *OperationExecutor) AddOperation(op *AsyncOperation) {
	// TODO: Agregar operación
}

// TODO 20: Implementa ExecuteAll que ejecute todas las operaciones
func (oe *OperationExecutor) ExecuteAll() map[string]OperationResult {
	// TODO: Ejecutar todas las operaciones en paralelo
	// Usar channels para recibir resultados
	return nil
}

// TODO 21: Implementa ExecuteWithTimeout
func (oe *OperationExecutor) ExecuteWithTimeout(timeout time.Duration) map[string]OperationResult {
	// TODO: Ejecutar con timeout global
	// Operaciones que no terminen a tiempo deben tener error de timeout
	return nil
}

// RecoveryHandler maneja panics
type RecoveryHandler struct {
	onPanic func(interface{})
}

// TODO 22: Implementa NewRecoveryHandler
func NewRecoveryHandler(onPanic func(interface{})) *RecoveryHandler {
	// TODO: Crear handler
	return nil
}

// TODO 23: Implementa SafeExecute que ejecute función con recovery
func (rh *RecoveryHandler) SafeExecute(fn func() error) error {
	// TODO: Ejecutar función con defer recover
	// Si hay panic, llamar onPanic y retornar error
	return nil
}

// TODO 24: Implementa SafeExecuteWithResult
func (rh *RecoveryHandler) SafeExecuteWithResult(fn func() (interface{}, error)) (interface{}, error) {
	// TODO: Similar a SafeExecute pero con resultado
	return nil, nil
}

// CircuitBreakerState representa estados del circuit breaker
type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	Open
	HalfOpen
)

// TODO 25: Implementa String() para CircuitBreakerState
func (s CircuitBreakerState) String() string {
	// TODO: Retornar nombre del estado
	return ""
}

// CircuitBreaker implementa patrón circuit breaker
type CircuitBreaker struct {
	// TODO: Agregar campos necesarios:
	// maxFailures, timeout, state, failures, lastFailure, mutex
}

// TODO 26: Implementa NewCircuitBreaker
func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	// TODO: Crear circuit breaker
	return nil
}

// TODO 27: Implementa Execute para CircuitBreaker
func (cb *CircuitBreaker) Execute(fn func() error) error {
	// TODO: Ejecutar función según estado del circuit breaker
	// - Closed: ejecutar normalmente, contar fallos
	// - Open: retornar inmediatamente CircuitBreakerError si no ha pasado timeout
	// - HalfOpen: intentar una ejecución para probar si el servicio se recuperó
	return nil
}

// TODO 28: Implementa métodos privados para cambiar estado
func (cb *CircuitBreaker) fail() {
	// TODO: Incrementar fallos y cambiar a Open si se alcanza el límite
}

func (cb *CircuitBreaker) success() {
	// TODO: Resetear fallos y cambiar a Closed
}

func (cb *CircuitBreaker) canAttempt() bool {
	// TODO: Verificar si se puede intentar la operación
	return false
}

// BulkOperationProcessor procesa operaciones en lotes
type BulkOperationProcessor struct {
	batchSize     int
	maxConcurrent int
	errorHandler  func(error)
}

// TODO 29: Implementa NewBulkOperationProcessor
func NewBulkOperationProcessor(batchSize, maxConcurrent int, errorHandler func(error)) *BulkOperationProcessor {
	// TODO: Crear processor
	return nil
}

// TODO 30: Implementa Process que procese items en lotes
func (bop *BulkOperationProcessor) Process(items []interface{}, processor func(interface{}) error) *ErrorAccumulator {
	// TODO: Procesar items en lotes de batchSize
	// Usar maxConcurrent goroutines
	// Acumular errores y continuar procesando
	return nil
}

// Simuladores para testing
func simulateNetworkCall() error {
	// Simula call de red que puede fallar
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	
	if rand.Float32() < 0.3 { // 30% chance de fallo
		if rand.Float32() < 0.5 {
			return TemporaryError{
				Message:    "network timeout",
				RetryAfter: time.Second,
			}
		} else {
			return PermanentError{
				Message: "unauthorized access",
			}
		}
	}
	return nil
}

func simulateDataProcessing(data interface{}) error {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	
	if rand.Float32() < 0.2 { // 20% chance de fallo
		return fmt.Errorf("processing failed for: %v", data)
	}
	return nil
}

func simulatePanicOperation() (interface{}, error) {
	if rand.Float32() < 0.3 { // 30% chance de panic
		panic("unexpected panic occurred!")
	}
	return "success", nil
}

func main() {
	fmt.Println("=== Patrones Avanzados de Error Handling ===\n")
	rand.Seed(time.Now().UnixNano())

	// Test de errores retryables
	fmt.Println("--- Errores Retryables ---")
	tempErr := TemporaryError{Message: "service unavailable", RetryAfter: 2 * time.Second}
	permErr := PermanentError{Message: "invalid credentials"}
	
	fmt.Printf("Error temporal: %v, Retryable: %v, Retry after: %v\n", 
		tempErr, tempErr.IsRetryable(), tempErr.GetRetryAfter())
	fmt.Printf("Error permanente: %v, Retryable: %v\n", 
		permErr, permErr.IsRetryable())

	// Test de error accumulator
	fmt.Println("\n--- Error Accumulator ---")
	accumulator := NewErrorAccumulator()
	
	// Simular múltiples goroutines agregando errores
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				err := fmt.Errorf("error from goroutine %d, iteration %d", id, j)
				accumulator.AddError(err)
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Printf("Total errors collected: %d\n", len(accumulator.GetErrors()))
	fmt.Printf("Has errors: %v\n", accumulator.HasErrors())

	// Test de operaciones asíncronas
	fmt.Println("\n--- Operaciones Asíncronas ---")
	executor := NewOperationExecutor()
	
	// Agregar operaciones
	executor.AddOperation(NewAsyncOperation("fast", func() (interface{}, error) {
		time.Sleep(100 * time.Millisecond)
		return "fast result", nil
	}))
	
	executor.AddOperation(NewAsyncOperation("slow", func() (interface{}, error) {
		time.Sleep(300 * time.Millisecond)
		return "slow result", nil
	}))
	
	executor.AddOperation(NewAsyncOperation("error", func() (interface{}, error) {
		time.Sleep(50 * time.Millisecond)
		return nil, fmt.Errorf("operation failed")
	}))
	
	// Ejecutar todas las operaciones
	fmt.Println("Ejecutando operaciones...")
	results := executor.ExecuteAll()
	
	for name, result := range results {
		if result.Error != nil {
			fmt.Printf("%s: ERROR - %v (took %v)\n", name, result.Error, result.Duration)
		} else {
			fmt.Printf("%s: SUCCESS - %v (took %v)\n", name, result.Value, result.Duration)
		}
	}

	// Test con timeout
	fmt.Println("\n--- Ejecución con Timeout ---")
	executor2 := NewOperationExecutor()
	executor2.AddOperation(NewAsyncOperation("quick", func() (interface{}, error) {
		time.Sleep(50 * time.Millisecond)
		return "quick done", nil
	}))
	
	executor2.AddOperation(NewAsyncOperation("timeout", func() (interface{}, error) {
		time.Sleep(500 * time.Millisecond) // Más largo que el timeout
		return "should timeout", nil
	}))
	
	resultsWithTimeout := executor2.ExecuteWithTimeout(200 * time.Millisecond)
	for name, result := range resultsWithTimeout {
		if result.Error != nil {
			fmt.Printf("%s: ERROR - %v\n", name, result.Error)
		} else {
			fmt.Printf("%s: SUCCESS - %v\n", name, result.Value)
		}
	}

	// Test de recovery handler
	fmt.Println("\n--- Recovery Handler ---")
	recovery := NewRecoveryHandler(func(panicValue interface{}) {
		fmt.Printf("Panic capturado: %v\n", panicValue)
	})
	
	// Operación normal
	err := recovery.SafeExecute(func() error {
		return nil
	})
	fmt.Printf("Operación normal: %v\n", err)
	
	// Operación que hace panic
	err = recovery.SafeExecute(func() error {
		panic("something went wrong!")
	})
	fmt.Printf("Operación con panic: %v\n", err)
	
	// Con resultado
	result, err := recovery.SafeExecuteWithResult(simulatePanicOperation)
	if err != nil {
		fmt.Printf("Error en operación: %v\n", err)
	} else {
		fmt.Printf("Resultado: %v\n", result)
	}

	// Test de circuit breaker
	fmt.Println("\n--- Circuit Breaker ---")
	cb := NewCircuitBreaker(3, 2*time.Second)
	
	// Simular múltiples llamadas para activar el circuit breaker
	for i := 0; i < 10; i++ {
		err := cb.Execute(simulateNetworkCall)
		if err != nil {
			fmt.Printf("Llamada %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Llamada %d: exitosa\n", i+1)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Test de bulk processor
	fmt.Println("\n--- Bulk Operation Processor ---")
	errorHandler := func(err error) {
		fmt.Printf("Error en batch: %v\n", err)
	}
	
	bulkProcessor := NewBulkOperationProcessor(3, 2, errorHandler)
	
	// Crear datos de prueba
	items := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		items[i] = fmt.Sprintf("item-%d", i)
	}
	
	errors := bulkProcessor.Process(items, simulateDataProcessing)
	
	fmt.Printf("Procesamiento completado. Errores: %d\n", len(errors.GetErrors()))
	for _, err := range errors.GetErrors() {
		fmt.Printf("  - %v\n", err)
	}

	fmt.Println("\n--- Resumen ---")
	fmt.Println("✅ Patrones de error handling implementados:")
	fmt.Println("  - Custom errors con contexto")
	fmt.Println("  - Error accumulation thread-safe")
	fmt.Println("  - Operaciones asíncronas con manejo de errores")
	fmt.Println("  - Recovery de panics")
	fmt.Println("  - Circuit breaker pattern")
	fmt.Println("  - Bulk processing con error handling")
}