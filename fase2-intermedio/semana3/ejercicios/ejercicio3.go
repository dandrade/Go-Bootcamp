package main

import (
	"fmt"
	"strings"
	"time"
)

// Address representa una dirección
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

// TODO 1: Implementa un method String() para Address que retorne la dirección formateada
// Formato: "Street, City, State ZipCode, Country"
func (a Address) String() string {
	// TODO: Implementar formato de dirección
	return ""
}

// Person representa una persona
type Person struct {
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Email       string
	Phone       string
}

// TODO 2: Implementa un method FullName() que retorne el nombre completo
func (p Person) FullName() string {
	// TODO: Retornar FirstName + " " + LastName
	return ""
}

// TODO 3: Implementa un method Age() que calcule la edad actual
func (p Person) Age() int {
	// TODO: Calcular edad basada en DateOfBirth
	// Pista: usa time.Now() y considera años completos
	return 0
}

// Department representa un departamento
type Department struct {
	ID       string
	Name     string
	Manager  string
	Budget   float64
	Location Address
}

// Employee representa un empleado (composición explícita)
type Employee struct {
	ID         string
	Person     Person     // Composición: un empleado TIENE una persona
	Department Department // Composición: un empleado TIENE un departamento
	Address    Address    // Composición: un empleado TIENE una dirección
	Position   string
	Salary     float64
	HireDate   time.Time
	IsActive   bool
}

// TODO 4: Implementa un constructor para Employee
func NewEmployee(id string, person Person, dept Department, address Address, position string, salary float64) *Employee {
	// TODO: Crear y retornar nuevo empleado
	// Inicializar HireDate con time.Now() e IsActive como true
	return nil
}

// TODO 5: Implementa un method que calcule años de servicio
func (e Employee) YearsOfService() int {
	// TODO: Calcular años desde HireDate hasta ahora
	return 0
}

// TODO 6: Implementa un method que retorne información resumida del empleado
// Formato: "FullName (Position) - Department - Years años"
func (e Employee) Summary() string {
	// TODO: Usar los métodos de Person y Employee para crear resumen
	return ""
}

// TODO 7: Implementa un method para dar aumento de salario (porcentaje)
func (e *Employee) GiveRaise(percentage float64) error {
	// TODO: Validar que el porcentaje sea positivo
	// Calcular y aplicar el nuevo salario
	return nil
}

// Company representa una empresa
type Company struct {
	Name         string
	Address      Address
	Founded      time.Time
	Employees    []Employee
	Departments  []Department
	TotalRevenue float64
}

// TODO 8: Implementa un constructor para Company
func NewCompany(name string, address Address) *Company {
	// TODO: Inicializar Company con slices vacíos
	return nil
}

// TODO 9: Implementa un method para agregar empleado
func (c *Company) AddEmployee(emp Employee) error {
	// TODO: Validar que el empleado no exista (por ID)
	// Agregar al slice si es válido
	return nil
}

// TODO 10: Implementa un method que encuentre empleados por departamento
func (c *Company) GetEmployeesByDepartment(deptID string) []Employee {
	// TODO: Filtrar y retornar empleados del departamento especificado
	return nil
}

// TODO 11: Implementa un method que calcule el salario total por departamento
func (c *Company) GetDepartmentPayroll(deptID string) float64 {
	// TODO: Sumar salarios de todos los empleados del departamento
	return 0
}

// TODO 12: Implementa un method que encuentre el empleado con más años de servicio
func (c *Company) GetMostSeniorEmployee() *Employee {
	// TODO: Encontrar y retornar el empleado con más años
	// Retornar nil si no hay empleados
	return nil
}

// Project representa un proyecto
type Project struct {
	ID          string
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Budget      float64
	TeamMembers []Employee // Relación muchos a muchos
}

// TODO 13: Implementa un method que verifique si el proyecto está activo
func (p Project) IsActive() bool {
	// TODO: Un proyecto está activo si la fecha actual está entre StartDate y EndDate
	return false
}

// TODO 14: Implementa un method que calcule el costo del equipo (suma de salarios)
// Considerar el período del proyecto (meses)
func (p Project) GetTeamCost() float64 {
	// TODO: Calcular costo total del equipo para la duración del proyecto
	return 0
}

func main() {
	fmt.Println("=== Sistema de Gestión de Empresa ===\n")

	// Crear direcciones
	companyAddr := Address{
		Street:  "123 Business Ave",
		City:    "Tech City",
		State:   "TC",
		ZipCode: "12345",
		Country: "USA",
	}

	// Crear empresa
	company := NewCompany("TechCorp Solutions", companyAddr)
	
	// Crear departamentos
	engineering := Department{
		ID:       "ENG",
		Name:     "Engineering",
		Manager:  "Alice Johnson",
		Budget:   1000000,
		Location: companyAddr,
	}

	marketing := Department{
		ID:       "MKT",
		Name:     "Marketing",
		Manager:  "Bob Smith",
		Budget:   500000,
		Location: companyAddr,
	}

	// Agregar departamentos
	company.Departments = append(company.Departments, engineering, marketing)

	// Crear empleados
	birthDate1, _ := time.Parse("2006-01-02", "1990-05-15")
	person1 := Person{
		FirstName:   "Carlos",
		LastName:    "García",
		DateOfBirth: birthDate1,
		Email:       "carlos@techcorp.com",
		Phone:       "+1234567890",
	}

	addr1 := Address{
		Street:  "456 Oak St",
		City:    "Tech City",
		State:   "TC",
		ZipCode: "12346",
		Country: "USA",
	}

	emp1 := NewEmployee("E001", person1, engineering, addr1, "Senior Developer", 85000)

	birthDate2, _ := time.Parse("2006-01-02", "1985-08-20")
	person2 := Person{
		FirstName:   "Diana",
		LastName:    "Martinez",
		DateOfBirth: birthDate2,
		Email:       "diana@techcorp.com",
		Phone:       "+1234567891",
	}

	emp2 := NewEmployee("E002", person2, engineering, addr1, "Tech Lead", 95000)
	if emp2 != nil {
		// Simular que fue contratada hace 3 años
		emp2.HireDate = time.Now().AddDate(-3, 0, 0)
	}

	// Agregar empleados
	if emp1 != nil {
		company.AddEmployee(*emp1)
	}
	if emp2 != nil {
		company.AddEmployee(*emp2)
	}

	// Mostrar información de la empresa
	fmt.Printf("Empresa: %s\n", company.Name)
	fmt.Printf("Dirección: %s\n", company.Address)
	fmt.Printf("Total de empleados: %d\n", len(company.Employees))
	fmt.Printf("Total de departamentos: %d\n\n", len(company.Departments))

	// Mostrar empleados
	fmt.Println("--- Empleados ---")
	for _, emp := range company.Employees {
		fmt.Printf("- %s\n", emp.Summary())
		fmt.Printf("  Edad: %d años\n", emp.Person.Age())
		fmt.Printf("  Email: %s\n", emp.Person.Email)
		fmt.Printf("  Salario: $%.2f\n\n", emp.Salary)
	}

	// Dar aumento
	fmt.Println("--- Aumentos de Salario ---")
	if len(company.Employees) > 0 {
		emp := &company.Employees[0]
		oldSalary := emp.Salary
		if err := emp.GiveRaise(10); err == nil {
			fmt.Printf("%s: $%.2f -> $%.2f (10%% aumento)\n", 
				emp.Person.FullName(), oldSalary, emp.Salary)
		}
	}

	// Análisis por departamento
	fmt.Println("\n--- Análisis por Departamento ---")
	engEmployees := company.GetEmployeesByDepartment("ENG")
	fmt.Printf("Empleados en Engineering: %d\n", len(engEmployees))
	fmt.Printf("Nómina total Engineering: $%.2f\n", company.GetDepartmentPayroll("ENG"))

	// Empleado más antiguo
	if senior := company.GetMostSeniorEmployee(); senior != nil {
		fmt.Printf("\nEmpleado más antiguo: %s (%d años de servicio)\n", 
			senior.Person.FullName(), senior.YearsOfService())
	}

	// Crear proyecto
	project := Project{
		ID:          "P001",
		Name:        "Nueva Plataforma Web",
		Description: "Desarrollo de plataforma web empresarial",
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(0, 6, 0), // 6 meses
		Budget:      500000,
		TeamMembers: company.Employees,
	}

	fmt.Println("\n--- Información del Proyecto ---")
	fmt.Printf("Proyecto: %s\n", project.Name)
	fmt.Printf("Activo: %v\n", project.IsActive())
	fmt.Printf("Duración: %v\n", project.EndDate.Sub(project.StartDate))
	fmt.Printf("Presupuesto: $%.2f\n", project.Budget)
	fmt.Printf("Costo del equipo: $%.2f\n", project.GetTeamCost())
}