package main

import (
	"fmt"
	"time"
)

// BankAccount representa una cuenta bancaria
type BankAccount struct {
	AccountNumber string
	Owner         string
	Balance       float64
	CreatedAt     time.Time
	IsActive      bool
}

// TODO 1: Implementa un method con VALUE RECEIVER que retorne el balance formateado
// Ejemplo: "$1,234.56" para un balance de 1234.56
func (b BankAccount) GetFormattedBalance() string {
	// TODO: Implementar formateo del balance
	// Pista: usa fmt.Sprintf con %.2f para 2 decimales
	return ""
}

// TODO 2: Implementa un method con POINTER RECEIVER para depositar dinero
// - Debe validar que el monto sea positivo
// - Retornar error si el monto es <= 0
// - Actualizar el balance si es válido
func (b *BankAccount) Deposit(amount float64) error {
	// TODO: Implementar depósito
	return nil
}

// TODO 3: Implementa un method con POINTER RECEIVER para retirar dinero
// - Validar que el monto sea positivo
// - Validar que haya suficiente balance
// - Retornar error apropiado si no se puede realizar
// - Actualizar el balance si es válido
func (b *BankAccount) Withdraw(amount float64) error {
	// TODO: Implementar retiro
	return nil
}

// TODO 4: Implementa un method con VALUE RECEIVER que indique si la cuenta puede retirar un monto
// No modifica nada, solo verifica
func (b BankAccount) CanWithdraw(amount float64) bool {
	// TODO: Verificar si se puede retirar el monto
	return false
}

// TODO 5: Implementa un method con POINTER RECEIVER para transferir a otra cuenta
// - Validar que el monto sea positivo
// - Validar que haya suficiente balance
// - Retirar de esta cuenta y depositar en la otra
// - Si falla el depósito, no hacer nada (transacción atómica)
func (b *BankAccount) TransferTo(other *BankAccount, amount float64) error {
	// TODO: Implementar transferencia
	// Pista: usa los methods Withdraw y Deposit ya implementados
	return nil
}

// Transaction representa una transacción
type Transaction struct {
	ID        string
	Type      string // "deposit", "withdrawal", "transfer"
	Amount    float64
	Timestamp time.Time
	FromAccount string
	ToAccount   string
}

// SavingsAccount extiende BankAccount con funcionalidad adicional
type SavingsAccount struct {
	BankAccount
	InterestRate   float64 // Tasa de interés anual (ej: 0.05 para 5%)
	MinimumBalance float64
}

// TODO 6: Implementa un constructor para SavingsAccount
// Debe inicializar todos los campos necesarios
func NewSavingsAccount(accountNumber, owner string, initialBalance, interestRate, minBalance float64) *SavingsAccount {
	// TODO: Crear y retornar nueva cuenta de ahorros
	return nil
}

// TODO 7: Override el method Withdraw para SavingsAccount
// Debe verificar que el balance no quede por debajo del mínimo
func (s *SavingsAccount) Withdraw(amount float64) error {
	// TODO: Implementar retiro con validación de balance mínimo
	return nil
}

// TODO 8: Implementa un method que calcule y aplique intereses
// Basado en el balance actual y la tasa de interés
// El período es en meses
func (s *SavingsAccount) ApplyInterest(months int) float64 {
	// TODO: Calcular y aplicar interés
	// Fórmula: interés = balance * (tasa_anual / 12) * meses
	return 0
}

// TODO 9: Define un tipo personalizado Money como float64
// e implementa methods para formateo
type Money float64

// TODO: Implementa String() para Money que retorne formato de moneda
func (m Money) String() string {
	// TODO: Retornar formato "$X.XX"
	return ""
}

// TODO 10: Implementa un method que convierta BankAccount balance a Money
func (b BankAccount) GetBalanceAsMoney() Money {
	// TODO: Convertir balance a Money
	return 0
}

func main() {
	fmt.Println("=== Sistema Bancario ===\n")

	// Crear cuentas
	account1 := &BankAccount{
		AccountNumber: "1001",
		Owner:         "Juan Pérez",
		Balance:       1000.00,
		CreatedAt:     time.Now(),
		IsActive:      true,
	}

	account2 := &BankAccount{
		AccountNumber: "1002",
		Owner:         "María García",
		Balance:       500.00,
		CreatedAt:     time.Now(),
		IsActive:      true,
	}

	// Mostrar balances iniciales
	fmt.Println("--- Balances Iniciales ---")
	fmt.Printf("%s: %s\n", account1.Owner, account1.GetFormattedBalance())
	fmt.Printf("%s: %s\n\n", account2.Owner, account2.GetFormattedBalance())

	// Realizar depósito
	fmt.Println("--- Operaciones ---")
	if err := account1.Deposit(250.50); err != nil {
		fmt.Printf("Error en depósito: %v\n", err)
	} else {
		fmt.Printf("Depósito exitoso. Nuevo balance de %s: %s\n", 
			account1.Owner, account1.GetFormattedBalance())
	}

	// Intentar retiro
	if account2.CanWithdraw(100) {
		if err := account2.Withdraw(100); err != nil {
			fmt.Printf("Error en retiro: %v\n", err)
		} else {
			fmt.Printf("Retiro exitoso. Nuevo balance de %s: %s\n", 
				account2.Owner, account2.GetFormattedBalance())
		}
	}

	// Transferencia
	fmt.Println("\n--- Transferencia ---")
	fmt.Printf("Transfiriendo $200 de %s a %s\n", account1.Owner, account2.Owner)
	if err := account1.TransferTo(account2, 200); err != nil {
		fmt.Printf("Error en transferencia: %v\n", err)
	} else {
		fmt.Println("Transferencia exitosa!")
		fmt.Printf("%s: %s\n", account1.Owner, account1.GetFormattedBalance())
		fmt.Printf("%s: %s\n", account2.Owner, account2.GetFormattedBalance())
	}

	// Cuenta de ahorros
	fmt.Println("\n--- Cuenta de Ahorros ---")
	savings := NewSavingsAccount("2001", "Carlos López", 5000, 0.05, 1000)
	if savings != nil {
		fmt.Printf("Cuenta de ahorros creada: %s\n", savings.GetFormattedBalance())
		
		// Aplicar intereses por 6 meses
		interest := savings.ApplyInterest(6)
		fmt.Printf("Intereses aplicados (6 meses): $%.2f\n", interest)
		fmt.Printf("Nuevo balance: %s\n", savings.GetFormattedBalance())
		
		// Intentar retiro que dejaría por debajo del mínimo
		fmt.Println("\nIntentando retirar $4200 (dejaría $%.2f)...", 
			savings.Balance - 4200)
		if err := savings.Withdraw(4200); err != nil {
			fmt.Printf("Retiro denegado: %v\n", err)
		}
	}

	// Usar tipo Money
	fmt.Println("\n--- Tipo Money ---")
	balance := account1.GetBalanceAsMoney()
	fmt.Printf("Balance de %s como Money: %s\n", account1.Owner, balance)
}