package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Sistema de Inventario Simple ===")
	
	// Inicializar el inventario
	inventory := NewInventory()
	
	// Cargar datos de ejemplo si se especifica
	if len(os.Args) > 1 && os.Args[1] == "-demo" {
		loadSampleData(inventory)
		fmt.Println("Datos de ejemplo cargados.")
	}
	
	// Mostrar menú principal
	for {
		showMainMenu()
		option := getUserInput("Seleccione una opción: ")
		
		switch option {
		case "1":
			addProduct(inventory)
		case "2":
			listProducts(inventory)
		case "3":
			searchProduct(inventory)
		case "4":
			updateProduct(inventory)
		case "5":
			deleteProduct(inventory)
		case "6":
			showReports(inventory)
		case "7":
			showFilterMenu(inventory)
		case "0":
			fmt.Println("¡Gracias por usar el sistema de inventario!")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
		
		fmt.Println() // Línea en blanco para mejor legibilidad
	}
}

// Función para mostrar el menú principal
func showMainMenu() {
	fmt.Println("\n--- MENÚ PRINCIPAL ---")
	fmt.Println("1. Agregar producto")
	fmt.Println("2. Listar productos")
	fmt.Println("3. Buscar producto")
	fmt.Println("4. Actualizar producto")
	fmt.Println("5. Eliminar producto")
	fmt.Println("6. Reportes")
	fmt.Println("7. Filtros avanzados")
	fmt.Println("0. Salir")
}

// Función para agregar un producto
func addProduct(inventory *Inventory) {
	fmt.Println("\n--- AGREGAR PRODUCTO ---")
	
	id := getUserInput("ID del producto: ")
	if inventory.ProductExists(id) {
		fmt.Printf("Error: Ya existe un producto con ID '%s'\n", id)
		return
	}
	
	name := getUserInput("Nombre: ")
	category := getUserInput("Categoría: ")
	priceStr := getUserInput("Precio: ")
	stockStr := getUserInput("Stock: ")
	description := getUserInput("Descripción: ")
	
	// Validar precio
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		fmt.Println("Error: Precio inválido")
		return
	}
	
	// Validar stock
	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		fmt.Println("Error: Stock inválido")
		return
	}
	
	product := Product{
		ID:          id,
		Name:        name,
		Category:    category,
		Price:       price,
		Stock:       stock,
		Description: description,
	}
	
	inventory.AddProduct(product)
	fmt.Printf("Producto '%s' agregado exitosamente.\n", name)
}

// Función para listar productos
func listProducts(inventory *Inventory) {
	fmt.Println("\n--- LISTA DE PRODUCTOS ---")
	
	products := inventory.GetAllProducts()
	if len(products) == 0 {
		fmt.Println("No hay productos en el inventario.")
		return
	}
	
	fmt.Printf("%-10s %-20s %-15s %-10s %-8s %s\n", "ID", "Nombre", "Categoría", "Precio", "Stock", "Descripción")
	fmt.Println(strings.Repeat("-", 80))
	
	for _, product := range products {
		fmt.Printf("%-10s %-20s %-15s $%-9.2f %-8d %s\n",
			product.ID, product.Name, product.Category, product.Price, product.Stock, product.Description)
	}
}

// Función para buscar un producto
func searchProduct(inventory *Inventory) {
	fmt.Println("\n--- BUSCAR PRODUCTO ---")
	fmt.Println("1. Buscar por ID")
	fmt.Println("2. Buscar por nombre")
	
	option := getUserInput("Seleccione opción: ")
	
	switch option {
	case "1":
		id := getUserInput("Ingrese ID: ")
		if product, exists := inventory.GetProductByID(id); exists {
			displayProduct(product)
		} else {
			fmt.Println("Producto no encontrado.")
		}
	case "2":
		name := getUserInput("Ingrese nombre (o parte del nombre): ")
		products := inventory.SearchByName(name)
		if len(products) == 0 {
			fmt.Println("No se encontraron productos.")
		} else {
			fmt.Printf("Se encontraron %d productos:\n", len(products))
			for _, product := range products {
				displayProduct(product)
				fmt.Println()
			}
		}
	default:
		fmt.Println("Opción no válida.")
	}
}

// Función para actualizar un producto
func updateProduct(inventory *Inventory) {
	fmt.Println("\n--- ACTUALIZAR PRODUCTO ---")
	
	id := getUserInput("ID del producto a actualizar: ")
	product, exists := inventory.GetProductByID(id)
	if !exists {
		fmt.Println("Producto no encontrado.")
		return
	}
	
	fmt.Println("Producto actual:")
	displayProduct(product)
	
	fmt.Println("\nIngrese los nuevos valores (Enter para mantener el actual):")
	
	// Actualizar campos
	name := getUserInputOrKeep("Nombre", product.Name)
	category := getUserInputOrKeep("Categoría", product.Category)
	description := getUserInputOrKeep("Descripción", product.Description)
	
	priceStr := getUserInputOrKeep("Precio", fmt.Sprintf("%.2f", product.Price))
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Error: Precio inválido")
		return
	}
	
	stockStr := getUserInputOrKeep("Stock", fmt.Sprintf("%d", product.Stock))
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		fmt.Println("Error: Stock inválido")
		return
	}
	
	// Crear producto actualizado
	updatedProduct := Product{
		ID:          id,
		Name:        name,
		Category:    category,
		Price:       price,
		Stock:       stock,
		Description: description,
	}
	
	inventory.UpdateProduct(updatedProduct)
	fmt.Println("Producto actualizado exitosamente.")
}

// Función para eliminar un producto
func deleteProduct(inventory *Inventory) {
	fmt.Println("\n--- ELIMINAR PRODUCTO ---")
	
	id := getUserInput("ID del producto a eliminar: ")
	product, exists := inventory.GetProductByID(id)
	if !exists {
		fmt.Println("Producto no encontrado.")
		return
	}
	
	fmt.Println("Producto a eliminar:")
	displayProduct(product)
	
	confirm := getUserInput("¿Está seguro? (s/n): ")
	if strings.ToLower(confirm) == "s" {
		inventory.DeleteProduct(id)
		fmt.Println("Producto eliminado exitosamente.")
	} else {
		fmt.Println("Operación cancelada.")
	}
}

// Función para mostrar reportes
func showReports(inventory *Inventory) {
	fmt.Println("\n--- REPORTES ---")
	fmt.Println("1. Resumen del inventario")
	fmt.Println("2. Productos por categoría")
	fmt.Println("3. Productos con stock bajo")
	fmt.Println("4. Productos más caros")
	fmt.Println("5. Valor total del inventario")
	
	option := getUserInput("Seleccione opción: ")
	
	switch option {
	case "1":
		showInventorySummary(inventory)
	case "2":
		showProductsByCategory(inventory)
	case "3":
		showLowStockProducts(inventory)
	case "4":
		showMostExpensiveProducts(inventory)
	case "5":
		showTotalValue(inventory)
	default:
		fmt.Println("Opción no válida.")
	}
}

// Función para mostrar menú de filtros
func showFilterMenu(inventory *Inventory) {
	fmt.Println("\n--- FILTROS AVANZADOS ---")
	fmt.Println("1. Filtrar por categoría")
	fmt.Println("2. Filtrar por rango de precios")
	fmt.Println("3. Filtrar por stock mínimo")
	
	option := getUserInput("Seleccione opción: ")
	
	switch option {
	case "1":
		category := getUserInput("Categoría: ")
		products := inventory.FilterByCategory(category)
		showFilteredProducts(products, fmt.Sprintf("Productos en categoría '%s'", category))
	case "2":
		minStr := getUserInput("Precio mínimo: ")
		maxStr := getUserInput("Precio máximo: ")
		
		min, err1 := strconv.ParseFloat(minStr, 64)
		max, err2 := strconv.ParseFloat(maxStr, 64)
		
		if err1 != nil || err2 != nil {
			fmt.Println("Error: Precios inválidos")
			return
		}
		
		products := inventory.FilterByPriceRange(min, max)
		showFilteredProducts(products, fmt.Sprintf("Productos entre $%.2f y $%.2f", min, max))
	case "3":
		minStr := getUserInput("Stock mínimo: ")
		minStock, err := strconv.Atoi(minStr)
		if err != nil {
			fmt.Println("Error: Stock inválido")
			return
		}
		
		products := inventory.FilterByMinStock(minStock)
		showFilteredProducts(products, fmt.Sprintf("Productos con stock >= %d", minStock))
	default:
		fmt.Println("Opción no válida.")
	}
}

// Función para mostrar productos filtrados
func showFilteredProducts(products []Product, title string) {
	fmt.Printf("\n--- %s ---\n", title)
	if len(products) == 0 {
		fmt.Println("No se encontraron productos.")
		return
	}
	
	for _, product := range products {
		displayProduct(product)
		fmt.Println()
	}
}

// Función para mostrar resumen del inventario
func showInventorySummary(inventory *Inventory) {
	fmt.Println("\n--- RESUMEN DEL INVENTARIO ---")
	stats := inventory.GetStatistics()
	
	fmt.Printf("Total de productos: %d\n", stats.TotalProducts)
	fmt.Printf("Total de items en stock: %d\n", stats.TotalItems)
	fmt.Printf("Valor total: $%.2f\n", stats.TotalValue)
	fmt.Printf("Precio promedio: $%.2f\n", stats.AveragePrice)
	fmt.Printf("Categorías: %d\n", stats.TotalCategories)
	
	if len(stats.Categories) > 0 {
		fmt.Println("Listado de categorías:")
		for category, count := range stats.Categories {
			fmt.Printf("  - %s: %d productos\n", category, count)
		}
	}
}

// Función para mostrar productos por categoría
func showProductsByCategory(inventory *Inventory) {
	fmt.Println("\n--- PRODUCTOS POR CATEGORÍA ---")
	categories := inventory.GetCategories()
	
	if len(categories) == 0 {
		fmt.Println("No hay categorías disponibles.")
		return
	}
	
	for _, category := range categories {
		products := inventory.FilterByCategory(category)
		fmt.Printf("\n%s (%d productos):\n", category, len(products))
		for _, product := range products {
			fmt.Printf("  - %s (ID: %s) - $%.2f\n", product.Name, product.ID, product.Price)
		}
	}
}

// Función para mostrar productos con stock bajo
func showLowStockProducts(inventory *Inventory) {
	fmt.Println("\n--- PRODUCTOS CON STOCK BAJO ---")
	
	limitStr := getUserInput("Límite de stock bajo (default: 5): ")
	limit := 5
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}
	
	products := inventory.GetLowStockProducts(limit)
	if len(products) == 0 {
		fmt.Printf("No hay productos con stock menor a %d.\n", limit)
		return
	}
	
	fmt.Printf("Productos con stock menor a %d:\n", limit)
	for _, product := range products {
		fmt.Printf("  - %s (ID: %s) - Stock: %d\n", product.Name, product.ID, product.Stock)
	}
}

// Función para mostrar productos más caros
func showMostExpensiveProducts(inventory *Inventory) {
	fmt.Println("\n--- PRODUCTOS MÁS CAROS ---")
	
	limitStr := getUserInput("Número de productos a mostrar (default: 5): ")
	limit := 5
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}
	
	products := inventory.GetMostExpensiveProducts(limit)
	if len(products) == 0 {
		fmt.Println("No hay productos disponibles.")
		return
	}
	
	fmt.Printf("Top %d productos más caros:\n", len(products))
	for i, product := range products {
		fmt.Printf("%d. %s (ID: %s) - $%.2f\n", i+1, product.Name, product.ID, product.Price)
	}
}

// Función para mostrar valor total del inventario
func showTotalValue(inventory *Inventory) {
	fmt.Println("\n--- VALOR TOTAL DEL INVENTARIO ---")
	totalValue := inventory.GetTotalValue()
	fmt.Printf("Valor total del inventario: $%.2f\n", totalValue)
}

// Función para mostrar un producto
func displayProduct(product Product) {
	fmt.Printf("ID: %s\n", product.ID)
	fmt.Printf("Nombre: %s\n", product.Name)
	fmt.Printf("Categoría: %s\n", product.Category)
	fmt.Printf("Precio: $%.2f\n", product.Price)
	fmt.Printf("Stock: %d\n", product.Stock)
	fmt.Printf("Descripción: %s\n", product.Description)
}

// Función para obtener entrada del usuario
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

// Función para obtener entrada del usuario o mantener valor actual
func getUserInputOrKeep(field, current string) string {
	fmt.Printf("%s [%s]: ", field, current)
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return current
	}
	return input
}

// Función para cargar datos de ejemplo
func loadSampleData(inventory *Inventory) {
	sampleProducts := []Product{
		{ID: "001", Name: "Laptop Dell", Category: "Electrónicos", Price: 899.99, Stock: 5, Description: "Laptop para uso profesional"},
		{ID: "002", Name: "Mouse Logitech", Category: "Electrónicos", Price: 25.50, Stock: 20, Description: "Mouse inalámbrico"},
		{ID: "003", Name: "Teclado Mecánico", Category: "Electrónicos", Price: 120.00, Stock: 8, Description: "Teclado mecánico RGB"},
		{ID: "004", Name: "Monitor 24\"", Category: "Electrónicos", Price: 300.00, Stock: 3, Description: "Monitor Full HD"},
		{ID: "005", Name: "Silla Ergonómica", Category: "Muebles", Price: 250.00, Stock: 12, Description: "Silla de oficina ergonómica"},
		{ID: "006", Name: "Escritorio", Category: "Muebles", Price: 180.00, Stock: 6, Description: "Escritorio de madera"},
		{ID: "007", Name: "Cuaderno A4", Category: "Papelería", Price: 3.50, Stock: 50, Description: "Cuaderno universitario"},
		{ID: "008", Name: "Bolígrafos", Category: "Papelería", Price: 1.20, Stock: 100, Description: "Pack de 10 bolígrafos"},
		{ID: "009", Name: "Calculadora", Category: "Electrónicos", Price: 15.00, Stock: 25, Description: "Calculadora científica"},
		{ID: "010", Name: "Lámpara LED", Category: "Iluminación", Price: 35.00, Stock: 15, Description: "Lámpara LED regulable"},
	}
	
	for _, product := range sampleProducts {
		inventory.AddProduct(product)
	}
}