# Structs y Methods en Go

## Introducción
En Go, los structs son la forma principal de agrupar datos relacionados. Junto con los methods, proporcionan una forma de programación orientada a objetos, pero sin clases ni herencia tradicional.

## 1. Structs Básicos

### Definición
Un struct es un tipo de dato que agrupa cero o más campos de diferentes tipos.

```go
// Definición básica
type Person struct {
    Name string
    Age  int
}

// Struct con tags
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}
```

### Creación e Inicialización
```go
// Forma 1: Declaración zero-value
var p1 Person
// p1.Name = "" (string vacío)
// p1.Age = 0

// Forma 2: Literal con valores
p2 := Person{
    Name: "Ana",
    Age:  25,
}

// Forma 3: Literal posicional (no recomendado)
p3 := Person{"Carlos", 30}

// Forma 4: Con new (retorna puntero)
p4 := new(Person)
p4.Name = "Diana"
p4.Age = 28
```

### Acceso a Campos
```go
person := Person{Name: "Eva", Age: 22}

// Lectura
fmt.Println(person.Name) // Eva

// Escritura
person.Age = 23

// Con punteros (Go desreferencia automáticamente)
p := &person
fmt.Println(p.Name) // Equivalente a (*p).Name
```

## 2. Methods

### Definición
Un method es una función con un receiver especial entre el keyword `func` y el nombre del método.

```go
// Value receiver
func (p Person) Greet() string {
    return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Name, p.Age)
}

// Pointer receiver
func (p *Person) HaveBirthday() {
    p.Age++
}
```

### Value vs Pointer Receivers

#### Value Receivers
- Operan sobre una copia del valor
- No pueden modificar el struct original
- Útiles para métodos de solo lectura

```go
func (p Person) GetInfo() string {
    return fmt.Sprintf("%s (%d años)", p.Name, p.Age)
}
```

#### Pointer Receivers
- Operan sobre el struct original
- Pueden modificar el struct
- Más eficientes para structs grandes
- Necesarios para modificar el estado

```go
func (p *Person) SetName(name string) {
    p.Name = name
}

func (p *Person) IncrementAge() {
    p.Age++
}
```

### Regla General
Si algún método necesita pointer receiver, todos deberían usarlo para consistencia.

## 3. Constructor Pattern

Go no tiene constructores, pero es común crear funciones `New`:

```go
// Constructor simple
func NewPerson(name string, age int) Person {
    return Person{
        Name: name,
        Age:  age,
    }
}

// Constructor que retorna puntero
func NewUser(username, email string) *User {
    return &User{
        ID:       generateID(),
        Username: username,
        Email:    email,
    }
}

// Constructor con validación
func NewAccount(balance float64) (*Account, error) {
    if balance < 0 {
        return nil, errors.New("balance no puede ser negativo")
    }
    
    return &Account{
        ID:      generateAccountID(),
        Balance: balance,
    }, nil
}
```

## 4. Composición

Go favorece la composición sobre la herencia:

```go
// Structs base
type Address struct {
    Street  string
    City    string
    Country string
}

type Person struct {
    Name string
    Age  int
}

// Composición explícita
type Employee struct {
    Person  Person
    Address Address
    Salary  float64
}

// Uso
emp := Employee{
    Person:  Person{Name: "Juan", Age: 30},
    Address: Address{City: "Madrid"},
    Salary:  50000,
}
fmt.Println(emp.Person.Name) // Acceso explícito
```

## 5. Embedding (Composición Implícita)

El embedding permite "heredar" campos y métodos:

```go
// Embedding de structs
type Employee struct {
    Person  // Embedded struct
    Address // Embedded struct
    Salary  float64
}

// Uso simplificado
emp := Employee{
    Person:  Person{Name: "María", Age: 28},
    Address: Address{City: "Barcelona"},
    Salary:  45000,
}

// Acceso directo a campos embedded
fmt.Println(emp.Name) // No necesitas emp.Person.Name
fmt.Println(emp.City) // No necesitas emp.Address.City

// Los métodos también se "heredan"
fmt.Println(emp.Greet()) // Llama a Person.Greet()
```

### Method Overriding con Embedding
```go
type Manager struct {
    Employee
    TeamSize int
}

// Override del método Greet
func (m Manager) Greet() string {
    return fmt.Sprintf("Hola, soy %s, manager de %d personas", m.Name, m.TeamSize)
}
```

## 6. Structs Anónimos

Útiles para datos temporales:

```go
// Struct anónimo en variable
data := struct {
    X int
    Y int
}{X: 10, Y: 20}

// En slices
points := []struct {
    X, Y float64
}{
    {1.0, 2.0},
    {3.0, 4.0},
}

// Como parámetro de función
func processData(data struct{ Name string; Value int }) {
    fmt.Printf("%s: %d\n", data.Name, data.Value)
}
```

## 7. Comparación de Structs

Los structs son comparables si todos sus campos lo son:

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}
p3 := Point{2, 3}

fmt.Println(p1 == p2) // true
fmt.Println(p1 == p3) // false

// No comparables si contienen slices, maps o functions
type Data struct {
    Values []int // slice hace el struct no comparable
}
```

## 8. Tags de Struct

Las tags proporcionan metadata para reflection:

```go
type Server struct {
    Host string `json:"host" xml:"server-host"`
    Port int    `json:"port" xml:"server-port"`
    TLS  bool   `json:"tls,omitempty" xml:"use-tls"`
}

// Uso común con encoding/json
server := Server{Host: "localhost", Port: 8080}
data, _ := json.Marshal(server)
fmt.Println(string(data)) // {"host":"localhost","port":8080}
```

## 9. Métodos en Tipos No-Struct

Puedes definir métodos en cualquier tipo definido:

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// Uso
temp := Celsius(25)
fmt.Printf("%.2f°C = %.2f°F\n", temp, temp.ToFahrenheit())
```

## 10. Best Practices

### 1. Nombres de Campos
- Usa PascalCase para campos exportados
- Usa camelCase para campos no exportados

### 2. Organización
```go
type User struct {
    // Agrupa campos relacionados
    // Identificación
    ID       string
    Username string
    
    // Información personal
    FirstName string
    LastName  string
    Email     string
    
    // Metadata
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 3. Inicialización
```go
// Prefiere nombres de campos explícitos
good := Person{
    Name: "Alice",
    Age:  30,
}

// Evita inicialización posicional
bad := Person{"Bob", 25}
```

### 4. Zero Values Útiles
Diseña structs para que el zero value sea útil:

```go
type Buffer struct {
    data []byte
}

// El zero value es usable
var buf Buffer
buf.Write([]byte("hello")) // Funciona sin inicialización explícita
```

## Ejemplo Completo: Sistema de Inventario

```go
package main

import (
    "fmt"
    "time"
)

// Producto representa un item en el inventario
type Product struct {
    ID          string
    Name        string
    Description string
    Price       float64
    Stock       int
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// NewProduct crea un nuevo producto con validación
func NewProduct(name, description string, price float64, stock int) (*Product, error) {
    if name == "" {
        return nil, fmt.Errorf("el nombre no puede estar vacío")
    }
    if price < 0 {
        return nil, fmt.Errorf("el precio no puede ser negativo")
    }
    if stock < 0 {
        return nil, fmt.Errorf("el stock no puede ser negativo")
    }
    
    now := time.Now()
    return &Product{
        ID:          generateID(),
        Name:        name,
        Description: description,
        Price:       price,
        Stock:       stock,
        CreatedAt:   now,
        UpdatedAt:   now,
    }, nil
}

// Value receiver - no modifica el producto
func (p Product) GetValue() float64 {
    return p.Price * float64(p.Stock)
}

// Value receiver - método de presentación
func (p Product) String() string {
    return fmt.Sprintf("%s - $%.2f (Stock: %d)", p.Name, p.Price, p.Stock)
}

// Pointer receiver - modifica el producto
func (p *Product) UpdatePrice(newPrice float64) error {
    if newPrice < 0 {
        return fmt.Errorf("el precio no puede ser negativo")
    }
    p.Price = newPrice
    p.UpdatedAt = time.Now()
    return nil
}

// Pointer receiver - modifica el stock
func (p *Product) AddStock(quantity int) error {
    if quantity < 0 {
        return fmt.Errorf("la cantidad debe ser positiva")
    }
    p.Stock += quantity
    p.UpdatedAt = time.Now()
    return nil
}

// Pointer receiver - vende productos
func (p *Product) Sell(quantity int) error {
    if quantity > p.Stock {
        return fmt.Errorf("stock insuficiente: disponible %d, solicitado %d", p.Stock, quantity)
    }
    p.Stock -= quantity
    p.UpdatedAt = time.Now()
    return nil
}

// Almacén agrupa productos
type Warehouse struct {
    Name     string
    Location string
    Products map[string]*Product
}

// NewWarehouse crea un nuevo almacén
func NewWarehouse(name, location string) *Warehouse {
    return &Warehouse{
        Name:     name,
        Location: location,
        Products: make(map[string]*Product),
    }
}

// AddProduct agrega un producto al almacén
func (w *Warehouse) AddProduct(p *Product) error {
    if p == nil {
        return fmt.Errorf("producto no puede ser nil")
    }
    if _, exists := w.Products[p.ID]; exists {
        return fmt.Errorf("producto con ID %s ya existe", p.ID)
    }
    w.Products[p.ID] = p
    return nil
}

// GetTotalValue calcula el valor total del inventario
func (w *Warehouse) GetTotalValue() float64 {
    total := 0.0
    for _, product := range w.Products {
        total += product.GetValue()
    }
    return total
}

// FindLowStock encuentra productos con stock bajo
func (w *Warehouse) FindLowStock(threshold int) []*Product {
    var lowStock []*Product
    for _, product := range w.Products {
        if product.Stock < threshold {
            lowStock = append(lowStock, product)
        }
    }
    return lowStock
}

func generateID() string {
    return fmt.Sprintf("PROD-%d", time.Now().UnixNano())
}

func main() {
    // Crear almacén
    warehouse := NewWarehouse("Almacén Central", "Madrid")
    
    // Crear productos
    laptop, _ := NewProduct("Laptop Pro", "Laptop de alta gama", 1299.99, 10)
    mouse, _ := NewProduct("Mouse Wireless", "Mouse inalámbrico", 29.99, 50)
    
    // Agregar al almacén
    warehouse.AddProduct(laptop)
    warehouse.AddProduct(mouse)
    
    // Operaciones
    laptop.Sell(2)
    mouse.AddStock(20)
    
    // Mostrar información
    fmt.Printf("Inventario total: $%.2f\n", warehouse.GetTotalValue())
    
    // Buscar productos con stock bajo
    lowStock := warehouse.FindLowStock(15)
    fmt.Println("\nProductos con stock bajo:")
    for _, p := range lowStock {
        fmt.Println("-", p)
    }
}
```

## Resumen

Los structs y methods son fundamentales en Go:
- **Structs** agrupan datos relacionados
- **Methods** agregan comportamiento a los tipos
- **Value receivers** para métodos de solo lectura
- **Pointer receivers** para modificar estado
- **Composición** es preferida sobre herencia
- **Embedding** proporciona una forma de "herencia"

Dominar estos conceptos te permitirá escribir código Go idiomático y bien estructurado.