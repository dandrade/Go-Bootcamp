package main

import (
	"sort"
	"strings"
)

// Estructura para representar un producto
type Product struct {
	ID          string
	Name        string
	Category    string
	Price       float64
	Stock       int
	Description string
}

// Estructura para representar el inventario
type Inventory struct {
	products map[string]Product
}

// Estructura para estadísticas del inventario
type InventoryStats struct {
	TotalProducts    int
	TotalItems       int
	TotalValue       float64
	AveragePrice     float64
	TotalCategories  int
	Categories       map[string]int
}

// Constructor para crear un nuevo inventario
func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]Product),
	}
}

// Método para agregar un producto al inventario
func (inv *Inventory) AddProduct(product Product) {
	inv.products[product.ID] = product
}

// Método para verificar si un producto existe
func (inv *Inventory) ProductExists(id string) bool {
	_, exists := inv.products[id]
	return exists
}

// Método para obtener un producto por ID
func (inv *Inventory) GetProductByID(id string) (Product, bool) {
	product, exists := inv.products[id]
	return product, exists
}

// Método para obtener todos los productos
func (inv *Inventory) GetAllProducts() []Product {
	var products []Product
	for _, product := range inv.products {
		products = append(products, product)
	}
	
	// Ordenar por ID para consistencia
	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})
	
	return products
}

// Método para actualizar un producto
func (inv *Inventory) UpdateProduct(product Product) {
	inv.products[product.ID] = product
}

// Método para eliminar un producto
func (inv *Inventory) DeleteProduct(id string) {
	delete(inv.products, id)
}

// Método para buscar productos por nombre
func (inv *Inventory) SearchByName(name string) []Product {
	var results []Product
	searchTerm := strings.ToLower(name)
	
	for _, product := range inv.products {
		if strings.Contains(strings.ToLower(product.Name), searchTerm) {
			results = append(results, product)
		}
	}
	
	return results
}

// Método para filtrar productos por categoría
func (inv *Inventory) FilterByCategory(category string) []Product {
	var results []Product
	categoryLower := strings.ToLower(category)
	
	for _, product := range inv.products {
		if strings.ToLower(product.Category) == categoryLower {
			results = append(results, product)
		}
	}
	
	return results
}

// Método para filtrar productos por rango de precios
func (inv *Inventory) FilterByPriceRange(min, max float64) []Product {
	var results []Product
	
	for _, product := range inv.products {
		if product.Price >= min && product.Price <= max {
			results = append(results, product)
		}
	}
	
	return results
}

// Método para filtrar productos por stock mínimo
func (inv *Inventory) FilterByMinStock(minStock int) []Product {
	var results []Product
	
	for _, product := range inv.products {
		if product.Stock >= minStock {
			results = append(results, product)
		}
	}
	
	return results
}

// Método para obtener productos con stock bajo
func (inv *Inventory) GetLowStockProducts(threshold int) []Product {
	var results []Product
	
	for _, product := range inv.products {
		if product.Stock < threshold {
			results = append(results, product)
		}
	}
	
	// Ordenar por stock (ascendente)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Stock < results[j].Stock
	})
	
	return results
}

// Método para obtener productos más caros
func (inv *Inventory) GetMostExpensiveProducts(limit int) []Product {
	products := inv.GetAllProducts()
	
	// Ordenar por precio (descendente)
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	
	// Limitar resultados
	if limit > len(products) {
		limit = len(products)
	}
	
	return products[:limit]
}

// Método para obtener el valor total del inventario
func (inv *Inventory) GetTotalValue() float64 {
	total := 0.0
	for _, product := range inv.products {
		total += product.Price * float64(product.Stock)
	}
	return total
}

// Método para obtener estadísticas del inventario
func (inv *Inventory) GetStatistics() InventoryStats {
	stats := InventoryStats{
		Categories: make(map[string]int),
	}
	
	totalPrice := 0.0
	
	for _, product := range inv.products {
		stats.TotalProducts++
		stats.TotalItems += product.Stock
		totalPrice += product.Price
		
		// Contar categorías
		if _, exists := stats.Categories[product.Category]; !exists {
			stats.Categories[product.Category] = 0
		}
		stats.Categories[product.Category]++
	}
	
	stats.TotalValue = inv.GetTotalValue()
	stats.TotalCategories = len(stats.Categories)
	
	if stats.TotalProducts > 0 {
		stats.AveragePrice = totalPrice / float64(stats.TotalProducts)
	}
	
	return stats
}

// Método para obtener lista de categorías únicas
func (inv *Inventory) GetCategories() []string {
	categoryMap := make(map[string]bool)
	
	for _, product := range inv.products {
		categoryMap[product.Category] = true
	}
	
	var categories []string
	for category := range categoryMap {
		categories = append(categories, category)
	}
	
	// Ordenar alfabéticamente
	sort.Strings(categories)
	
	return categories
}

// Método para obtener productos por múltiples criterios
func (inv *Inventory) SearchProducts(criteria SearchCriteria) []Product {
	var results []Product
	
	for _, product := range inv.products {
		if inv.matchesCriteria(product, criteria) {
			results = append(results, product)
		}
	}
	
	return results
}

// Estructura para criterios de búsqueda
type SearchCriteria struct {
	Name         string
	Category     string
	MinPrice     float64
	MaxPrice     float64
	MinStock     int
	MaxStock     int
	HasStock     bool
}

// Método auxiliar para verificar si un producto cumple los criterios
func (inv *Inventory) matchesCriteria(product Product, criteria SearchCriteria) bool {
	// Verificar nombre
	if criteria.Name != "" {
		if !strings.Contains(strings.ToLower(product.Name), strings.ToLower(criteria.Name)) {
			return false
		}
	}
	
	// Verificar categoría
	if criteria.Category != "" {
		if strings.ToLower(product.Category) != strings.ToLower(criteria.Category) {
			return false
		}
	}
	
	// Verificar precio mínimo
	if criteria.MinPrice > 0 {
		if product.Price < criteria.MinPrice {
			return false
		}
	}
	
	// Verificar precio máximo
	if criteria.MaxPrice > 0 {
		if product.Price > criteria.MaxPrice {
			return false
		}
	}
	
	// Verificar stock mínimo
	if criteria.MinStock > 0 {
		if product.Stock < criteria.MinStock {
			return false
		}
	}
	
	// Verificar stock máximo
	if criteria.MaxStock > 0 {
		if product.Stock > criteria.MaxStock {
			return false
		}
	}
	
	// Verificar si tiene stock
	if criteria.HasStock {
		if product.Stock <= 0 {
			return false
		}
	}
	
	return true
}

// Método para obtener productos similares (por categoría)
func (inv *Inventory) GetSimilarProducts(productID string) []Product {
	product, exists := inv.GetProductByID(productID)
	if !exists {
		return []Product{}
	}
	
	var similar []Product
	for _, p := range inv.products {
		if p.ID != productID && strings.ToLower(p.Category) == strings.ToLower(product.Category) {
			similar = append(similar, p)
		}
	}
	
	return similar
}

// Método para obtener un resumen rápido del inventario
func (inv *Inventory) GetQuickSummary() map[string]interface{} {
	summary := make(map[string]interface{})
	
	summary["total_products"] = len(inv.products)
	summary["total_value"] = inv.GetTotalValue()
	summary["categories"] = len(inv.GetCategories())
	summary["low_stock"] = len(inv.GetLowStockProducts(5))
	
	return summary
}