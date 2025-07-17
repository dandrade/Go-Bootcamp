package main

import (
	"fmt"
	"time"
)

// Engine representa un motor
type Engine struct {
	Type        string  // "Gasoline", "Electric", "Hybrid"
	Horsepower  int
	Cylinders   int     // 0 para eléctricos
	FuelType    string
}

// TODO 1: Implementa un method Start() para Engine
func (e Engine) Start() string {
	// TODO: Retornar mensaje diferente según el tipo de motor
	// Gasoline: "Vroom! Engine started"
	// Electric: "Silent start... Ready to go"
	// Hybrid: "Hybrid system activated"
	return ""
}

// TODO 2: Implementa un method GetInfo() para Engine
func (e Engine) GetInfo() string {
	// TODO: Retornar información formateada del motor
	return ""
}

// Vehicle es la estructura base para todos los vehículos
type Vehicle struct {
	Brand        string
	Model        string
	Year         int
	Color        string
	VIN          string // Vehicle Identification Number
	Mileage      float64
	IsRunning    bool
	LastService  time.Time
}

// TODO 3: Implementa methods básicos para Vehicle
// GetAge() - retorna años desde Year hasta ahora
func (v Vehicle) GetAge() int {
	// TODO: Calcular edad del vehículo
	return 0
}

// NeedsService() - true si han pasado más de 6 meses desde LastService
func (v Vehicle) NeedsService() bool {
	// TODO: Verificar si necesita servicio
	return false
}

// GetDescription() - retorna "Year Brand Model (Color)"
func (v Vehicle) GetDescription() string {
	// TODO: Crear descripción del vehículo
	return ""
}

// Car embebe Vehicle y Engine
type Car struct {
	Vehicle // Embedding
	Engine  // Embedding
	Doors   int
	Seats   int
	Type    string // "Sedan", "SUV", "Coupe", etc.
}

// TODO 4: Implementa un constructor para Car
func NewCar(brand, model string, year int, engineType string, hp int) *Car {
	// TODO: Crear y retornar nuevo Car
	// Inicializar con valores apropiados
	return nil
}

// TODO 5: Override el method Start() para Car
// Debe verificar que el carro no esté ya encendido
func (c *Car) Start() string {
	// TODO: Si ya está encendido, retornar "Car is already running"
	// Si no, encender y retornar el mensaje del motor
	return ""
}

// TODO 6: Implementa Stop() para Car
func (c *Car) Stop() string {
	// TODO: Si no está encendido, retornar "Car is not running"
	// Si está encendido, apagar y retornar "Engine stopped"
	return ""
}

// Motorcycle embebe Vehicle y Engine
type Motorcycle struct {
	Vehicle
	Engine
	Type       string // "Sport", "Cruiser", "Touring"
	HasSidecar bool
}

// TODO 7: Implementa un constructor para Motorcycle
func NewMotorcycle(brand, model string, year int, engineType string, hp int, motoType string) *Motorcycle {
	// TODO: Crear y retornar nueva Motorcycle
	return nil
}

// TODO 8: Override GetDescription() para Motorcycle
// Debe incluir el tipo de moto
func (m Motorcycle) GetDescription() string {
	// TODO: Retornar descripción incluyendo tipo
	// Formato: "Year Brand Model Type Motorcycle (Color)"
	return ""
}

// ElectricScooter embebe solo Vehicle (no tiene Engine tradicional)
type ElectricScooter struct {
	Vehicle
	BatteryCapacity float64 // en kWh
	Range           float64 // en km
	MaxSpeed        float64 // en km/h
	ChargingTime    float64 // en horas
}

// TODO 9: Implementa Start() para ElectricScooter
func (e *ElectricScooter) Start() string {
	// TODO: Implementar encendido de scooter eléctrico
	return ""
}

// TODO 10: Implementa GetRange() que calcule autonomía restante
// basado en el porcentaje de batería (asume 100% al inicio)
func (e ElectricScooter) GetRange() float64 {
	// TODO: Por simplicidad, retorna Range completo
	return 0
}

// Fleet representa una flota de vehículos
type Fleet struct {
	Name     string
	Vehicles []interface{} // Puede contener Car, Motorcycle, ElectricScooter
}

// TODO 11: Implementa AddVehicle para Fleet
func (f *Fleet) AddVehicle(vehicle interface{}) {
	// TODO: Agregar vehículo a la flota
}

// TODO 12: Implementa CountByType que cuente vehículos por tipo
func (f Fleet) CountByType() map[string]int {
	// TODO: Contar cuántos Car, Motorcycle y ElectricScooter hay
	// Pista: usa type switch
	return nil
}

// TODO 13: Implementa GetOldestVehicle
func (f Fleet) GetOldestVehicle() interface{} {
	// TODO: Encontrar el vehículo más antiguo de la flota
	// Pista: necesitarás type assertion para acceder a Vehicle
	return nil
}

// Garage puede realizar mantenimiento a cualquier vehículo
type Garage struct {
	Name     string
	Location string
}

// TODO 14: Implementa PerformService que actualice LastService
// Debe funcionar con Car, Motorcycle y ElectricScooter
func (g Garage) PerformService(vehicle interface{}) string {
	// TODO: Usar type switch para manejar diferentes tipos
	// Actualizar LastService y retornar mensaje apropiado
	return ""
}

func main() {
	fmt.Println("=== Sistema de Gestión de Vehículos ===\n")

	// Crear diferentes vehículos
	car1 := NewCar("Toyota", "Camry", 2022, "Gasoline", 203)
	if car1 != nil {
		car1.Color = "Silver"
		car1.VIN = "1HGCM82633A123456"
		car1.Doors = 4
		car1.Seats = 5
		car1.Type = "Sedan"
		car1.LastService = time.Now().AddDate(0, -8, 0) // Hace 8 meses
	}

	car2 := NewCar("Tesla", "Model 3", 2023, "Electric", 283)
	if car2 != nil {
		car2.Color = "Blue"
		car2.VIN = "5YJ3E1EA1JF123456"
		car2.Doors = 4
		car2.Seats = 5
		car2.Type = "Sedan"
	}

	moto := NewMotorcycle("Harley-Davidson", "Street 750", 2021, "Gasoline", 53, "Cruiser")
	if moto != nil {
		moto.Color = "Black"
		moto.VIN = "1HD1KB4119Y123456"
		moto.Cylinders = 2
		moto.FuelType = "Gasoline"
	}

	scooter := &ElectricScooter{
		Vehicle: Vehicle{
			Brand:   "Xiaomi",
			Model:   "Mi Electric Scooter Pro 2",
			Year:    2023,
			Color:   "Black",
			VIN:     "XMI2023PRO123456",
			Mileage: 150,
		},
		BatteryCapacity: 0.474,  // kWh
		Range:           45,     // km
		MaxSpeed:        25,     // km/h
		ChargingTime:    8.5,    // hours
	}

	// Mostrar información de vehículos
	fmt.Println("--- Información de Vehículos ---")
	if car1 != nil {
		fmt.Printf("Car 1: %s\n", car1.GetDescription())
		fmt.Printf("  Edad: %d años\n", car1.GetAge())
		fmt.Printf("  Motor: %s\n", car1.Engine.GetInfo())
		fmt.Printf("  Necesita servicio: %v\n", car1.NeedsService())
	}

	if moto != nil {
		fmt.Printf("\nMotorcycle: %s\n", moto.GetDescription())
		fmt.Printf("  Motor: %s\n", moto.Engine.GetInfo())
	}

	fmt.Printf("\nScooter: %s\n", scooter.GetDescription())
	fmt.Printf("  Autonomía: %.1f km\n", scooter.GetRange())

	// Operaciones con vehículos
	fmt.Println("\n--- Operaciones ---")
	if car1 != nil {
		fmt.Println(car1.Start())
		fmt.Println(car1.Start()) // Intentar encender de nuevo
		fmt.Println(car1.Stop())
	}

	fmt.Println(scooter.Start())

	// Crear flota
	fleet := Fleet{Name: "Flota Corporativa"}
	fleet.AddVehicle(car1)
	fleet.AddVehicle(car2)
	fleet.AddVehicle(moto)
	fleet.AddVehicle(scooter)

	fmt.Println("\n--- Análisis de Flota ---")
	fmt.Printf("Total de vehículos: %d\n", len(fleet.Vehicles))
	
	counts := fleet.CountByType()
	for vType, count := range counts {
		fmt.Printf("%s: %d\n", vType, count)
	}

	if oldest := fleet.GetOldestVehicle(); oldest != nil {
		// Type assertion para mostrar información
		switch v := oldest.(type) {
		case *Car:
			fmt.Printf("\nVehículo más antiguo: %s\n", v.GetDescription())
		case *Motorcycle:
			fmt.Printf("\nVehículo más antiguo: %s\n", v.GetDescription())
		case *ElectricScooter:
			fmt.Printf("\nVehículo más antiguo: %s\n", v.GetDescription())
		}
	}

	// Servicio de mantenimiento
	garage := Garage{
		Name:     "SuperGarage",
		Location: "Downtown",
	}

	fmt.Println("\n--- Servicio de Mantenimiento ---")
	if car1 != nil && car1.NeedsService() {
		fmt.Println(garage.PerformService(car1))
		fmt.Printf("¿Necesita servicio ahora? %v\n", car1.NeedsService())
	}
}