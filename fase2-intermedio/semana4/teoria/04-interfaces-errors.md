# Interfaces y Manejo de Errores en Go

## Introducción
Las interfaces son una de las características más poderosas de Go. Proporcionan una forma de definir comportamiento sin especificar implementación, permitiendo escribir código flexible y reutilizable. Junto con el sistema de errores de Go, forman la base para crear aplicaciones robustas.

## 1. Interfaces Básicas

### Definición
Una interface define un conjunto de métodos que un tipo debe implementar. Son implícitas - no necesitas declarar que un tipo implementa una interface.

```go
// Definición de interface
type Writer interface {
    Write(data []byte) (int, error)
}

type Reader interface {
    Read(data []byte) (int, error)
}

// Interface compuesta
type ReadWriter interface {
    Reader
    Writer
}
```

### Implementación Implícita
En Go, no declares que implementas una interface. Si tu tipo tiene los métodos requeridos, automáticamente implementa la interface.

```go
// FileWriter implementa Writer automáticamente
type FileWriter struct {
    filename string
}

func (fw *FileWriter) Write(data []byte) (int, error) {
    file, err := os.OpenFile(fw.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    
    return file.Write(data)
}

// ConsoleWriter también implementa Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
    return fmt.Print(string(data))
}
```

### Uso de Interfaces
```go
func WriteData(w Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

func main() {
    fileWriter := &FileWriter{filename: "output.txt"}
    consoleWriter := ConsoleWriter{}
    
    data := []byte("Hello, World!")
    
    // Mismo código, diferentes implementaciones
    WriteData(fileWriter, data)    // Escribe a archivo
    WriteData(consoleWriter, data) // Escribe a consola
}
```

## 2. Interface Vacía

La interface vacía `interface{}` puede contener cualquier valor:

```go
func PrintAny(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func main() {
    PrintAny(42)
    PrintAny("hello")
    PrintAny([]int{1, 2, 3})
}
```

### Type Assertions
Para extraer el valor concreto de una interface:

```go
func ProcessValue(value interface{}) {
    // Type assertion simple
    if str, ok := value.(string); ok {
        fmt.Printf("String: %s\n", str)
        return
    }
    
    // Type switch
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case []int:
        fmt.Printf("Int slice: %v\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## 3. Interfaces Comunes de Go

### io.Reader y io.Writer
```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// Ejemplo de uso
func CopyData(dst Writer, src Reader) error {
    buffer := make([]byte, 1024)
    for {
        n, err := src.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            return err
        }
        
        _, writeErr := dst.Write(buffer[:n])
        if writeErr != nil {
            return writeErr
        }
    }
    return nil
}
```

### fmt.Stringer
```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d años)", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Ana", Age: 25}
    fmt.Println(p) // Usa automáticamente String()
}
```

## 4. Sistema de Errores en Go

### Error Interface
```go
type error interface {
    Error() string
}
```

### Errores Básicos
```go
import "errors"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

// Usando fmt.Errorf
func ValidateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("edad inválida: %d (debe ser >= 0)", age)
    }
    return nil
}
```

### Custom Errors
```go
// Error simple
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// Error con más información
type DatabaseError struct {
    Operation string
    Table     string
    Cause     error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on table %s: %v", 
        e.Operation, e.Table, e.Cause)
}

func (e DatabaseError) Unwrap() error {
    return e.Cause
}
```

### Error Wrapping (Go 1.13+)
```go
import (
    "errors"
    "fmt"
)

func ProcessFile(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("failed to process file %s: %w", filename, err)
    }
    
    if err := validateData(data); err != nil {
        return fmt.Errorf("validation failed for %s: %w", filename, err)
    }
    
    return nil
}

func main() {
    err := ProcessFile("nonexistent.txt")
    if err != nil {
        // Verificar errores específicos
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("File does not exist")
        }
        
        // Extraer errores específicos
        var pathErr *os.PathError
        if errors.As(err, &pathErr) {
            fmt.Printf("Path error: %s\n", pathErr.Path)
        }
    }
}
```

## 5. Patrones de Error Handling

### Múltiples Return Values
```go
func ReadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("cannot read config file: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("invalid config format: %w", err)
    }
    
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    
    return &config, nil
}
```

### Error Accumulation
```go
type MultiError struct {
    Errors []error
}

func (me MultiError) Error() string {
    if len(me.Errors) == 0 {
        return "no errors"
    }
    
    var parts []string
    for _, err := range me.Errors {
        parts = append(parts, err.Error())
    }
    return strings.Join(parts, "; ")
}

func ValidateUser(user User) error {
    var errors []error
    
    if user.Name == "" {
        errors = append(errors, ValidationError{Field: "name", Message: "cannot be empty"})
    }
    
    if user.Age < 0 {
        errors = append(errors, ValidationError{Field: "age", Message: "must be non-negative"})
    }
    
    if !isValidEmail(user.Email) {
        errors = append(errors, ValidationError{Field: "email", Message: "invalid format"})
    }
    
    if len(errors) > 0 {
        return MultiError{Errors: errors}
    }
    
    return nil
}
```

### Panic y Recover
```go
func SafeOperation() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    
    // Operación que puede hacer panic
    result = riskyOperation()
    return result, nil
}

func riskyOperation() string {
    // Simular operación riesgosa
    panic("something went wrong")
}
```

## 6. Interfaces para Arquitectura

### Repository Pattern
```go
type UserRepository interface {
    Create(user User) error
    GetByID(id string) (*User, error)
    Update(user User) error
    Delete(id string) error
    List() ([]User, error)
}

// Implementación en memoria
type InMemoryUserRepository struct {
    users map[string]User
    mu    sync.RWMutex
}

func (r *InMemoryUserRepository) Create(user User) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.users[user.ID]; exists {
        return fmt.Errorf("user with ID %s already exists", user.ID)
    }
    
    r.users[user.ID] = user
    return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    user, exists := r.users[id]
    if !exists {
        return nil, fmt.Errorf("user with ID %s not found", id)
    }
    
    return &user, nil
}

// Implementación con base de datos
type DatabaseUserRepository struct {
    db *sql.DB
}

func (r *DatabaseUserRepository) Create(user User) error {
    query := `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`
    _, err := r.db.Exec(query, user.ID, user.Name, user.Email)
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}

// Service usando la interface
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
    user := User{
        ID:    generateID(),
        Name:  name,
        Email: email,
    }
    
    if err := s.repo.Create(user); err != nil {
        return nil, fmt.Errorf("cannot create user: %w", err)
    }
    
    return &user, nil
}
```

### Strategy Pattern
```go
type Formatter interface {
    Format(data interface{}) ([]byte, error)
}

type JSONFormatter struct{}

func (f JSONFormatter) Format(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

type XMLFormatter struct{}

func (f XMLFormatter) Format(data interface{}) ([]byte, error) {
    return xml.Marshal(data)
}

type DataExporter struct {
    formatter Formatter
}

func (e *DataExporter) Export(data interface{}) ([]byte, error) {
    return e.formatter.Format(data)
}

func main() {
    data := map[string]string{"message": "hello"}
    
    // Usar JSON
    jsonExporter := &DataExporter{formatter: JSONFormatter{}}
    jsonData, _ := jsonExporter.Export(data)
    
    // Usar XML
    xmlExporter := &DataExporter{formatter: XMLFormatter{}}
    xmlData, _ := xmlExporter.Export(data)
}
```

## 7. Testing con Interfaces

### Mocking
```go
type EmailSender interface {
    SendEmail(to, subject, body string) error
}

type UserService struct {
    emailSender EmailSender
}

func (s *UserService) RegisterUser(user User) error {
    // Lógica de registro...
    
    return s.emailSender.SendEmail(
        user.Email,
        "Welcome!",
        "Thanks for registering",
    )
}

// Mock para testing
type MockEmailSender struct {
    SentEmails []EmailData
}

type EmailData struct {
    To      string
    Subject string
    Body    string
}

func (m *MockEmailSender) SendEmail(to, subject, body string) error {
    m.SentEmails = append(m.SentEmails, EmailData{
        To:      to,
        Subject: subject,
        Body:    body,
    })
    return nil
}

func TestUserRegistration(t *testing.T) {
    mockSender := &MockEmailSender{}
    service := &UserService{emailSender: mockSender}
    
    user := User{Email: "test@example.com"}
    err := service.RegisterUser(user)
    
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if len(mockSender.SentEmails) != 1 {
        t.Fatalf("Expected 1 email, got %d", len(mockSender.SentEmails))
    }
    
    email := mockSender.SentEmails[0]
    if email.To != user.Email {
        t.Errorf("Expected email to %s, got %s", user.Email, email.To)
    }
}
```

## 8. Ejemplo Completo: Sistema de Logging

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
    "time"
)

// Logger interface
type Logger interface {
    Info(message string, fields ...Field)
    Error(message string, err error, fields ...Field)
    Debug(message string, fields ...Field)
}

// Field para logging estructurado
type Field struct {
    Key   string
    Value interface{}
}

// LogEntry representa una entrada de log
type LogEntry struct {
    Timestamp time.Time              `json:"timestamp"`
    Level     string                 `json:"level"`
    Message   string                 `json:"message"`
    Error     string                 `json:"error,omitempty"`
    Fields    map[string]interface{} `json:"fields,omitempty"`
}

// ConsoleLogger escribe logs a la consola
type ConsoleLogger struct {
    writer io.Writer
}

func NewConsoleLogger() *ConsoleLogger {
    return &ConsoleLogger{writer: os.Stdout}
}

func (l *ConsoleLogger) Info(message string, fields ...Field) {
    l.log("INFO", message, "", fields...)
}

func (l *ConsoleLogger) Error(message string, err error, fields ...Field) {
    errorMsg := ""
    if err != nil {
        errorMsg = err.Error()
    }
    l.log("ERROR", message, errorMsg, fields...)
}

func (l *ConsoleLogger) Debug(message string, fields ...Field) {
    l.log("DEBUG", message, "", fields...)
}

func (l *ConsoleLogger) log(level, message, errorMsg string, fields ...Field) {
    entry := LogEntry{
        Timestamp: time.Now(),
        Level:     level,
        Message:   message,
        Error:     errorMsg,
        Fields:    make(map[string]interface{}),
    }
    
    for _, field := range fields {
        entry.Fields[field.Key] = field.Value
    }
    
    data, _ := json.Marshal(entry)
    fmt.Fprintf(l.writer, "%s\n", data)
}

// FileLogger escribe logs a archivo
type FileLogger struct {
    filename string
    file     *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, fmt.Errorf("cannot open log file: %w", err)
    }
    
    return &FileLogger{
        filename: filename,
        file:     file,
    }, nil
}

func (l *FileLogger) Info(message string, fields ...Field) {
    l.log("INFO", message, "", fields...)
}

func (l *FileLogger) Error(message string, err error, fields ...Field) {
    errorMsg := ""
    if err != nil {
        errorMsg = err.Error()
    }
    l.log("ERROR", message, errorMsg, fields...)
}

func (l *FileLogger) Debug(message string, fields ...Field) {
    l.log("DEBUG", message, "", fields...)
}

func (l *FileLogger) log(level, message, errorMsg string, fields ...Field) {
    entry := LogEntry{
        Timestamp: time.Now(),
        Level:     level,
        Message:   message,
        Error:     errorMsg,
        Fields:    make(map[string]interface{}),
    }
    
    for _, field := range fields {
        entry.Fields[field.Key] = field.Value
    }
    
    data, _ := json.Marshal(entry)
    fmt.Fprintf(l.file, "%s\n", data)
}

func (l *FileLogger) Close() error {
    return l.file.Close()
}

// MultiLogger escribe a múltiples destinos
type MultiLogger struct {
    loggers []Logger
}

func NewMultiLogger(loggers ...Logger) *MultiLogger {
    return &MultiLogger{loggers: loggers}
}

func (l *MultiLogger) Info(message string, fields ...Field) {
    for _, logger := range l.loggers {
        logger.Info(message, fields...)
    }
}

func (l *MultiLogger) Error(message string, err error, fields ...Field) {
    for _, logger := range l.loggers {
        logger.Error(message, err, fields...)
    }
}

func (l *MultiLogger) Debug(message string, fields ...Field) {
    for _, logger := range l.loggers {
        logger.Debug(message, fields...)
    }
}

// Service usando el logger
type UserService struct {
    logger Logger
}

func NewUserService(logger Logger) *UserService {
    return &UserService{logger: logger}
}

func (s *UserService) CreateUser(name, email string) error {
    s.logger.Info("Creating user", 
        Field{Key: "name", Value: name},
        Field{Key: "email", Value: email})
    
    // Simular error
    if name == "" {
        err := fmt.Errorf("name cannot be empty")
        s.logger.Error("Failed to create user", err,
            Field{Key: "email", Value: email})
        return err
    }
    
    s.logger.Info("User created successfully",
        Field{Key: "name", Value: name},
        Field{Key: "email", Value: email})
    
    return nil
}

func main() {
    // Crear diferentes loggers
    consoleLogger := NewConsoleLogger()
    
    fileLogger, err := NewFileLogger("app.log")
    if err != nil {
        panic(err)
    }
    defer fileLogger.Close()
    
    // Logger que escribe a ambos destinos
    multiLogger := NewMultiLogger(consoleLogger, fileLogger)
    
    // Usar el service
    service := NewUserService(multiLogger)
    
    service.CreateUser("Ana", "ana@example.com")
    service.CreateUser("", "invalid@example.com") // Error
}
```

## Best Practices

### 1. Interfaces Pequeñas
Prefiere interfaces pequeñas y enfocadas:
```go
// Bien - interface específica
type Reader interface {
    Read([]byte) (int, error)
}

// Evita - interface muy grande
type FileManager interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
    Close() error
    Seek(int64, int) (int64, error)
    Stat() (os.FileInfo, error)
    // ... muchos más métodos
}
```

### 2. Acepta Interfaces, Retorna Structs
```go
// Bien - acepta interface
func ProcessData(r io.Reader) (*Result, error) {
    // ...
}

// Bien - retorna struct concreto
func NewFileProcessor(filename string) *FileProcessor {
    // ...
}
```

### 3. Nombres de Interfaces
```go
// Bien - verbo + er
type Reader interface { Read([]byte) (int, error) }
type Writer interface { Write([]byte) (int, error) }
type Stringer interface { String() string }

// Interfaces más complejas
type UserRepository interface { /* ... */ }
type EmailSender interface { /* ... */ }
```

### 4. Error Handling
```go
// Bien - información específica del contexto
func SaveUser(user User) error {
    if err := validateUser(user); err != nil {
        return fmt.Errorf("invalid user data: %w", err)
    }
    
    if err := db.Save(user); err != nil {
        return fmt.Errorf("failed to save user %s: %w", user.ID, err)
    }
    
    return nil
}
```

## Resumen

Las interfaces y el manejo de errores son fundamentales para escribir código Go idiomático:

- **Interfaces** proporcionan polimorfismo y flexibilidad
- **Implementación implícita** hace el código menos acoplado
- **Error interface** permite errores ricos en información
- **Custom errors** mejoran la depuración y manejo
- **Wrapping errors** preserva contexto importante
- **Interfaces pequeñas** son más fáciles de implementar y testear

Dominar estos conceptos te permitirá escribir código Go robusto, flexible y mantenible.