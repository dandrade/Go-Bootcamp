package main

import (
	"fmt"
	"math"
)

// TODO 1: Define la interface Shape que debe tener los métodos:
// - Area() float64
// - Perimeter() float64
// - String() string
type Shape interface {
	// TODO: Definir métodos de la interface
}

// TODO 2: Define la interface Drawable que extienda Shape y agregue:
// - Draw() string
type Drawable interface {
	// TODO: Embeber Shape y agregar Draw()
}

// Circle representa un círculo
type Circle struct {
	Radius float64
}

// TODO 3: Implementa los métodos de Shape para Circle
func (c Circle) Area() float64 {
	// TODO: Calcular área del círculo (π * r²)
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO: Calcular perímetro del círculo (2 * π * r)
	return 0
}

func (c Circle) String() string {
	// TODO: Retornar representación string "Círculo(radio=X)"
	return ""
}

// TODO 4: Implementa Draw() para que Circle implemente Drawable
func (c Circle) Draw() string {
	// TODO: Retornar representación ASCII simple del círculo
	return ""
}

// Rectangle representa un rectángulo
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO 5: Implementa los métodos de Shape para Rectangle
func (r Rectangle) Area() float64 {
	// TODO: Calcular área del rectángulo (width * height)
	return 0
}

func (r Rectangle) Perimeter() float64 {
	// TODO: Calcular perímetro del rectángulo (2 * (width + height))
	return 0
}

func (r Rectangle) String() string {
	// TODO: Retornar "Rectángulo(ancho=X, alto=Y)"
	return ""
}

// TODO 6: Implementa Draw() para Rectangle
func (r Rectangle) Draw() string {
	// TODO: Retornar representación ASCII del rectángulo
	return ""
}

// Triangle representa un triángulo
type Triangle struct {
	Side1, Side2, Side3 float64
}

// TODO 7: Implementa los métodos de Shape para Triangle
func (t Triangle) Area() float64 {
	// TODO: Usar fórmula de Herón para calcular el área
	// s = (a + b + c) / 2
	// area = sqrt(s * (s-a) * (s-b) * (s-c))
	return 0
}

func (t Triangle) Perimeter() float64 {
	// TODO: Calcular perímetro (suma de lados)
	return 0
}

func (t Triangle) String() string {
	// TODO: Retornar "Triángulo(lados=X,Y,Z)"
	return ""
}

// TODO 8: Implementa Draw() para Triangle
func (t Triangle) Draw() string {
	// TODO: Retornar representación ASCII simple
	return ""
}

// TODO 9: Implementa la función CalculateTotalArea que reciba un slice de Shape
// y retorne el área total
func CalculateTotalArea(shapes []Shape) float64 {
	// TODO: Sumar las áreas de todas las formas
	return 0
}

// TODO 10: Implementa la función FindLargestShape que encuentre la forma con mayor área
func FindLargestShape(shapes []Shape) Shape {
	// TODO: Encontrar y retornar la forma con mayor área
	// Retornar nil si el slice está vacío
	return nil
}

// TODO 11: Implementa la función PrintShapeInfo que imprima información de cualquier Shape
func PrintShapeInfo(shape Shape) {
	// TODO: Imprimir información formateada de la forma
	// Ejemplo: "Círculo(radio=5): Área=78.54, Perímetro=31.42"
}

// TODO 12: Implementa la función DrawAll que dibuje todas las formas Drawable
func DrawAll(drawables []Drawable) {
	// TODO: Dibujar todas las formas que implementen Drawable
}

// TODO 13: Define la interface Resizable con método:
// - Resize(factor float64)
type Resizable interface {
	// TODO: Definir método Resize
}

// TODO 14: Implementa Resize() para Circle (cambia el radio)
func (c *Circle) Resize(factor float64) {
	// TODO: Multiplicar el radio por el factor
}

// TODO 15: Implementa Resize() para Rectangle (cambia width y height)
func (r *Rectangle) Resize(factor float64) {
	// TODO: Multiplicar width y height por el factor
}

// TODO 16: Implementa Resize() para Triangle (cambia todos los lados)
func (t *Triangle) Resize(factor float64) {
	// TODO: Multiplicar todos los lados por el factor
}

// TODO 17: Implementa la función ResizeAll que redimensione todas las formas
func ResizeAll(shapes []Resizable, factor float64) {
	// TODO: Redimensionar todas las formas por el factor dado
}

// TODO 18: Define el tipo ShapeType como enum
type ShapeType int

const (
	// TODO: Definir constantes CircleType, RectangleType, TriangleType
)

// TODO 19: Implementa la función CreateShape que cree formas basado en el tipo
func CreateShape(shapeType ShapeType, params ...float64) Shape {
	// TODO: Crear la forma apropiada basada en el tipo y parámetros
	// Circle: 1 parámetro (radio)
	// Rectangle: 2 parámetros (width, height)
	// Triangle: 3 parámetros (side1, side2, side3)
	return nil
}

// TODO 20: Implementa la función que verifique si una interface es de cierto tipo
func IsCircle(shape Shape) bool {
	// TODO: Usar type assertion para verificar si es Circle
	return false
}

func IsRectangle(shape Shape) bool {
	// TODO: Usar type assertion para verificar si es Rectangle
	return false
}

func main() {
	fmt.Println("=== Sistema de Formas Geométricas ===\n")

	// Crear formas
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	triangle := Triangle{Side1: 3.0, Side2: 4.0, Side3: 5.0}

	// Mostrar información individual
	fmt.Println("--- Información de Formas ---")
	PrintShapeInfo(circle)
	PrintShapeInfo(rectangle)
	PrintShapeInfo(triangle)

	// Trabajar con slice de shapes
	shapes := []Shape{circle, rectangle, triangle}

	fmt.Printf("\n--- Análisis de Formas ---\n")
	fmt.Printf("Área total: %.2f\n", CalculateTotalArea(shapes))
	
	largest := FindLargestShape(shapes)
	if largest != nil {
		fmt.Printf("Forma más grande: %s\n", largest)
	}

	// Test de polimorfismo con Drawable
	fmt.Println("\n--- Dibujos ---")
	drawables := []Drawable{circle, rectangle, triangle}
	DrawAll(drawables)

	// Test de redimensionamiento
	fmt.Println("\n--- Redimensionamiento ---")
	resizables := []Resizable{&circle, &rectangle, &triangle}
	
	fmt.Println("Antes del redimensionamiento:")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
	
	ResizeAll(resizables, 1.5)
	
	fmt.Println("\nDespués del redimensionamiento (factor 1.5):")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}

	// Test de factory
	fmt.Println("\n--- Factory Pattern ---")
	newCircle := CreateShape(CircleType, 3.0)
	newRect := CreateShape(RectangleType, 2.0, 8.0)
	newTriangle := CreateShape(TriangleType, 6.0, 8.0, 10.0)
	
	if newCircle != nil {
		PrintShapeInfo(newCircle)
	}
	if newRect != nil {
		PrintShapeInfo(newRect)
	}
	if newTriangle != nil {
		PrintShapeInfo(newTriangle)
	}

	// Test de type assertions
	fmt.Println("\n--- Type Assertions ---")
	for i, shape := range shapes {
		fmt.Printf("Forma %d: ", i+1)
		if IsCircle(shape) {
			fmt.Println("Es un círculo")
		} else if IsRectangle(shape) {
			fmt.Println("Es un rectángulo")
		} else {
			fmt.Println("Es otro tipo de forma")
		}
	}

	// Ejemplo de type switch
	fmt.Println("\n--- Type Switch ---")
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Círculo con radio %.2f\n", s.Radius)
		case Rectangle:
			fmt.Printf("Rectángulo de %.2f x %.2f\n", s.Width, s.Height)
		case Triangle:
			fmt.Printf("Triángulo con lados %.2f, %.2f, %.2f\n", s.Side1, s.Side2, s.Side3)
		default:
			fmt.Printf("Forma desconocida: %T\n", s)
		}
	}
}