package main

import (
	"fmt"
	"strings"
	"time"
)

// TODO 1: Define la interface DataProcessor que tenga métodos:
// - Process(data []byte) ([]byte, error)
// - GetProcessorName() string
type DataProcessor interface {
	// TODO: Definir métodos
}

// TODO 2: Define la interface DataValidator que tenga:
// - Validate(data []byte) error
type DataValidator interface {
	// TODO: Definir método
}

// TODO 3: Define la interface ProcessorWithValidation que combine ambas
type ProcessorWithValidation interface {
	// TODO: Combinar DataProcessor y DataValidator
}

// TODO 4: Define la interface Configurable que tenga:
// - Configure(config map[string]interface{}) error
// - GetConfig() map[string]interface{}
type Configurable interface {
	// TODO: Definir métodos de configuración
}

// UpperCaseProcessor convierte texto a mayúsculas
type UpperCaseProcessor struct {
	name string
}

// TODO 5: Implementa DataProcessor para UpperCaseProcessor
func (u UpperCaseProcessor) Process(data []byte) ([]byte, error) {
	// TODO: Convertir data a mayúsculas
	return nil, nil
}

func (u UpperCaseProcessor) GetProcessorName() string {
	// TODO: Retornar nombre del procesador
	return ""
}

// TODO 6: Implementa DataValidator para UpperCaseProcessor
func (u UpperCaseProcessor) Validate(data []byte) error {
	// TODO: Validar que data no esté vacía
	return nil
}

// CompressProcessor simula compresión de datos
type CompressProcessor struct {
	name           string
	compressionLevel int
	config         map[string]interface{}
}

// TODO 7: Implementa DataProcessor para CompressProcessor
func (c CompressProcessor) Process(data []byte) ([]byte, error) {
	// TODO: Simular compresión (por simplicidad, solo agregar prefijo)
	// Formato: "COMPRESSED[level=X]:" + data
	return nil, nil
}

func (c CompressProcessor) GetProcessorName() string {
	// TODO: Retornar nombre con nivel de compresión
	return ""
}

// TODO 8: Implementa DataValidator para CompressProcessor
func (c CompressProcessor) Validate(data []byte) error {
	// TODO: Validar que data no esté vacía y que tenga al menos 10 bytes
	return nil
}

// TODO 9: Implementa Configurable para CompressProcessor
func (c *CompressProcessor) Configure(config map[string]interface{}) error {
	// TODO: Configurar compressionLevel desde config["level"]
	// Validar que level esté entre 1 y 9
	return nil
}

func (c CompressProcessor) GetConfig() map[string]interface{} {
	// TODO: Retornar configuración actual
	return nil
}

// EncryptProcessor simula encriptación
type EncryptProcessor struct {
	name      string
	algorithm string
	key       string
	config    map[string]interface{}
}

// TODO 10: Implementa DataProcessor para EncryptProcessor
func (e EncryptProcessor) Process(data []byte) ([]byte, error) {
	// TODO: Simular encriptación
	// Formato: "ENCRYPTED[algo]:" + base64-like encoding
	return nil, nil
}

func (e EncryptProcessor) GetProcessorName() string {
	// TODO: Retornar nombre con algoritmo
	return ""
}

// TODO 11: Implementa DataValidator para EncryptProcessor
func (e EncryptProcessor) Validate(data []byte) error {
	// TODO: Validar que data no esté vacía y que key esté configurada
	return nil
}

// TODO 12: Implementa Configurable para EncryptProcessor
func (e *EncryptProcessor) Configure(config map[string]interface{}) error {
	// TODO: Configurar algorithm y key desde config
	// Validar que ambos existan
	return nil
}

func (e EncryptProcessor) GetConfig() map[string]interface{} {
	// TODO: Retornar configuración actual (no incluir key por seguridad)
	return nil
}

// ProcessingPipeline maneja una cadena de procesadores
type ProcessingPipeline struct {
	processors []DataProcessor
	name       string
}

// TODO 13: Implementa NewProcessingPipeline
func NewProcessingPipeline(name string) *ProcessingPipeline {
	// TODO: Crear nueva pipeline
	return nil
}

// TODO 14: Implementa AddProcessor para agregar procesadores
func (p *ProcessingPipeline) AddProcessor(processor DataProcessor) {
	// TODO: Agregar procesador a la pipeline
}

// TODO 15: Implementa Process que ejecute todos los procesadores en secuencia
func (p *ProcessingPipeline) Process(data []byte) ([]byte, error) {
	// TODO: Procesar data a través de todos los procesadores
	// Si alguno falla, retornar error inmediatamente
	return nil, nil
}

// TODO 16: Implementa Validate que valide con todos los procesadores que lo soporten
func (p *ProcessingPipeline) Validate(data []byte) error {
	// TODO: Validar data con todos los procesadores que implementen DataValidator
	// Usar type assertion para verificar si implementan DataValidator
	return nil
}

// TODO 17: Implementa GetProcessorNames que retorne nombres de todos los procesadores
func (p *ProcessingPipeline) GetProcessorNames() []string {
	// TODO: Obtener nombres de todos los procesadores
	return nil
}

// ProcessorFactory crea procesadores
type ProcessorFactory struct{}

// TODO 18: Implementa CreateProcessor que cree procesadores por tipo
func (f ProcessorFactory) CreateProcessor(processorType string, config map[string]interface{}) (DataProcessor, error) {
	// TODO: Crear procesador basado en tipo
	// Tipos soportados: "uppercase", "compress", "encrypt"
	// Si el procesador es Configurable, configurarlo
	return nil, nil
}

// TODO 19: Implementa GetSupportedTypes que retorne tipos soportados
func (f ProcessorFactory) GetSupportedTypes() []string {
	// TODO: Retornar slice con tipos soportados
	return nil
}

// ProcessorManager gestiona múltiples pipelines
type ProcessorManager struct {
	pipelines map[string]*ProcessingPipeline
	factory   ProcessorFactory
}

// TODO 20: Implementa NewProcessorManager
func NewProcessorManager() *ProcessorManager {
	// TODO: Crear nuevo manager
	return nil
}

// TODO 21: Implementa CreatePipeline
func (m *ProcessorManager) CreatePipeline(name string, processorConfigs []map[string]interface{}) error {
	// TODO: Crear pipeline con procesadores configurados
	// Cada config debe tener "type" y otros parámetros
	return nil
}

// TODO 22: Implementa ProcessWithPipeline
func (m *ProcessorManager) ProcessWithPipeline(pipelineName string, data []byte) ([]byte, error) {
	// TODO: Procesar data usando pipeline específica
	return nil, nil
}

// TODO 23: Implementa GetPipelineInfo
func (m *ProcessorManager) GetPipelineInfo(pipelineName string) ([]string, error) {
	// TODO: Obtener información de procesadores en la pipeline
	return nil, nil
}

// TODO 24: Implementa ListPipelines
func (m *ProcessorManager) ListPipelines() []string {
	// TODO: Retornar nombres de todas las pipelines
	return nil
}

// DataProcessingStats para métricas
type DataProcessingStats struct {
	ProcessorName    string
	ProcessingTime   time.Duration
	InputSize        int
	OutputSize       int
	Success          bool
	Error            string
}

// TODO 25: Implementa ProcessWithStats que procese y recolecte métricas
func ProcessWithStats(processor DataProcessor, data []byte) (*DataProcessingStats, []byte, error) {
	// TODO: Medir tiempo de procesamiento y recolectar estadísticas
	return nil, nil, nil
}

func main() {
	fmt.Println("=== Sistema de Procesamiento de Datos ===\n")

	// Crear procesadores individuales
	fmt.Println("--- Procesadores Individuales ---")
	
	upperProcessor := UpperCaseProcessor{name: "UpperCase"}
	
	compressProcessor := CompressProcessor{
		name:   "Compressor",
		config: make(map[string]interface{}),
	}
	compressConfig := map[string]interface{}{
		"level": 5,
	}
	compressProcessor.Configure(compressConfig)
	
	encryptProcessor := EncryptProcessor{
		name:   "Encryptor",
		config: make(map[string]interface{}),
	}
	encryptConfig := map[string]interface{}{
		"algorithm": "AES",
		"key":       "secret123",
	}
	encryptProcessor.Configure(encryptConfig)

	// Test de procesadores individuales
	testData := []byte("Hello, World! This is test data for processing.")
	
	fmt.Printf("Datos originales: %s\n", testData)
	
	// Test UpperCase
	if err := upperProcessor.Validate(testData); err != nil {
		fmt.Printf("Error de validación: %v\n", err)
	} else {
		processed, err := upperProcessor.Process(testData)
		if err != nil {
			fmt.Printf("Error de procesamiento: %v\n", err)
		} else {
			fmt.Printf("%s: %s\n", upperProcessor.GetProcessorName(), processed)
		}
	}

	// Test Compress
	if err := compressProcessor.Validate(testData); err != nil {
		fmt.Printf("Error de validación: %v\n", err)
	} else {
		processed, err := compressProcessor.Process(testData)
		if err != nil {
			fmt.Printf("Error de procesamiento: %v\n", err)
		} else {
			fmt.Printf("%s: %s\n", compressProcessor.GetProcessorName(), processed)
		}
	}

	// Test Encrypt
	if err := encryptProcessor.Validate(testData); err != nil {
		fmt.Printf("Error de validación: %v\n", err)
	} else {
		processed, err := encryptProcessor.Process(testData)
		if err != nil {
			fmt.Printf("Error de procesamiento: %v\n", err)
		} else {
			fmt.Printf("%s: %s\n", encryptProcessor.GetProcessorName(), processed)
		}
	}

	// Test de pipeline
	fmt.Println("\n--- Pipeline de Procesamiento ---")
	pipeline := NewProcessingPipeline("Main Pipeline")
	pipeline.AddProcessor(upperProcessor)
	pipeline.AddProcessor(compressProcessor)
	pipeline.AddProcessor(encryptProcessor)
	
	fmt.Printf("Procesadores en pipeline: %v\n", pipeline.GetProcessorNames())
	
	if err := pipeline.Validate(testData); err != nil {
		fmt.Printf("Error de validación de pipeline: %v\n", err)
	} else {
		finalResult, err := pipeline.Process(testData)
		if err != nil {
			fmt.Printf("Error en pipeline: %v\n", err)
		} else {
			fmt.Printf("Resultado final: %s\n", finalResult)
		}
	}

	// Test de factory
	fmt.Println("\n--- Factory Pattern ---")
	factory := ProcessorFactory{}
	fmt.Printf("Tipos soportados: %v\n", factory.GetSupportedTypes())
	
	newProcessor, err := factory.CreateProcessor("compress", map[string]interface{}{
		"level": 9,
	})
	if err != nil {
		fmt.Printf("Error creando procesador: %v\n", err)
	} else {
		fmt.Printf("Procesador creado: %s\n", newProcessor.GetProcessorName())
	}

	// Test de manager
	fmt.Println("\n--- Manager de Pipelines ---")
	manager := NewProcessorManager()
	
	pipelineConfigs := []map[string]interface{}{
		{"type": "uppercase"},
		{"type": "compress", "level": 3},
		{"type": "encrypt", "algorithm": "DES", "key": "mykey"},
	}
	
	err = manager.CreatePipeline("TestPipeline", pipelineConfigs)
	if err != nil {
		fmt.Printf("Error creando pipeline: %v\n", err)
	} else {
		fmt.Println("Pipeline creada exitosamente")
		
		info, _ := manager.GetPipelineInfo("TestPipeline")
		fmt.Printf("Info de pipeline: %v\n", info)
		
		result, err := manager.ProcessWithPipeline("TestPipeline", testData)
		if err != nil {
			fmt.Printf("Error procesando: %v\n", err)
		} else {
			fmt.Printf("Resultado con manager: %s\n", result)
		}
	}

	fmt.Printf("\nPipelines disponibles: %v\n", manager.ListPipelines())

	// Test de estadísticas
	fmt.Println("\n--- Estadísticas de Procesamiento ---")
	stats, result, err := ProcessWithStats(upperProcessor, testData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Estadísticas:\n")
		fmt.Printf("  Procesador: %s\n", stats.ProcessorName)
		fmt.Printf("  Tiempo: %v\n", stats.ProcessingTime)
		fmt.Printf("  Tamaño entrada: %d bytes\n", stats.InputSize)
		fmt.Printf("  Tamaño salida: %d bytes\n", stats.OutputSize)
		fmt.Printf("  Exitoso: %v\n", stats.Success)
		fmt.Printf("  Resultado: %s\n", result)
	}
}