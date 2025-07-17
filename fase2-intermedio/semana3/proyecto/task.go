package main

import (
	"fmt"
	"strings"
	"time"
)

// TODO 1: Define el enum Priority
// Valores: Low, Medium, High
type Priority int

const (
	// TODO: Definir constantes para Low, Medium, High
)

// TODO 2: Implementa String() para Priority
func (p Priority) String() string {
	// TODO: Retornar "Baja", "Media", "Alta"
	return ""
}

// TODO 3: Define el enum Status
// Valores: Pending, InProgress, Completed
type Status int

const (
	// TODO: Definir constantes para Pending, InProgress, Completed
)

// TODO 4: Implementa String() para Status
func (s Status) String() string {
	// TODO: Retornar "Pendiente", "En Progreso", "Completada"
	return ""
}

// TODO 5: Define el struct Task con todos los campos especificados
type Task struct {
	// TODO: Implementar todos los campos
	// ID, Title, Description, Priority, Status, CreatedAt, UpdatedAt, DueDate, Tags
}

// TODO 6: Implementa un constructor NewTask
// Debe generar un ID único, establecer CreatedAt y UpdatedAt, y Status como Pending
func NewTask(title, description string, priority Priority) *Task {
	// TODO: Implementar constructor
	// Pista: usa time.Now() para las fechas y generateID() para el ID
	return nil
}

// TODO 7: Implementa String() para Task que muestre la información formateada
func (t Task) String() string {
	// TODO: Crear representación legible de la tarea
	// Incluir ID, título, estado, prioridad
	return ""
}

// TODO 8: Implementa IsOverdue() que verifique si la tarea está vencida
func (t Task) IsOverdue() bool {
	// TODO: Verificar si DueDate existe y ya pasó
	return false
}

// TODO 9: Implementa MarkComplete() que marque la tarea como completada
func (t *Task) MarkComplete() {
	// TODO: Cambiar Status a Completed y actualizar UpdatedAt
}

// TODO 10: Implementa SetStatus() para cambiar el estado
func (t *Task) SetStatus(status Status) {
	// TODO: Cambiar status y actualizar UpdatedAt
}

// TODO 11: Implementa SetPriority() para cambiar la prioridad
func (t *Task) SetPriority(priority Priority) {
	// TODO: Cambiar prioridad y actualizar UpdatedAt
}

// TODO 12: Implementa SetDueDate() para establecer fecha límite
func (t *Task) SetDueDate(dueDate time.Time) {
	// TODO: Establecer DueDate y actualizar UpdatedAt
}

// TODO 13: Implementa ClearDueDate() para quitar fecha límite
func (t *Task) ClearDueDate() {
	// TODO: Poner DueDate en nil y actualizar UpdatedAt
}

// TODO 14: Implementa AddTag() para agregar etiqueta
func (t *Task) AddTag(tag string) {
	// TODO: Verificar que el tag no exista ya, agregarlo si es único
	// Actualizar UpdatedAt
}

// TODO 15: Implementa RemoveTag() para remover etiqueta
func (t *Task) RemoveTag(tag string) bool {
	// TODO: Buscar y remover el tag si existe
	// Actualizar UpdatedAt y retornar true si se removió
	return false
}

// TODO 16: Implementa HasTag() para verificar si tiene una etiqueta
func (t Task) HasTag(tag string) bool {
	// TODO: Verificar si el tag existe en la lista
	return false
}

// TODO 17: Implementa GetFormattedDueDate() que retorne la fecha formateada
func (t Task) GetFormattedDueDate() string {
	// TODO: Retornar fecha formateada o "Sin fecha límite"
	return ""
}

// TODO 18: Implementa GetTagsString() que retorne tags como string separado por comas
func (t Task) GetTagsString() string {
	// TODO: Unir tags con comas, o retornar "Sin etiquetas"
	return ""
}

// TODO 19: Implementa DaysUntilDue() que calcule días hasta la fecha límite
func (t Task) DaysUntilDue() int {
	// TODO: Calcular días hasta DueDate
	// Retornar 0 si no hay fecha o si ya pasó
	return 0
}

// TODO 20: Implementa Update() para actualizar título y descripción
func (t *Task) Update(title, description string) {
	// TODO: Actualizar campos y UpdatedAt
}

// generateID genera un ID único basado en timestamp
func generateID() string {
	return fmt.Sprintf("T%d", time.Now().UnixNano())
}

// Funciones helper para parsing (ya implementadas)
func ParsePriority(input string) (Priority, error) {
	switch strings.ToLower(input) {
	case "1", "baja", "low":
		return Low, nil
	case "2", "media", "medium":
		return Medium, nil
	case "3", "alta", "high":
		return High, nil
	default:
		return Low, fmt.Errorf("prioridad inválida: %s", input)
	}
}

func ParseStatus(input string) (Status, error) {
	switch strings.ToLower(input) {
	case "1", "pendiente", "pending":
		return Pending, nil
	case "2", "progreso", "progress", "inprogress":
		return InProgress, nil
	case "3", "completada", "completed", "done":
		return Completed, nil
	default:
		return Pending, fmt.Errorf("estado inválido: %s", input)
	}
}