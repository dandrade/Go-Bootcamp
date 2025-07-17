package repository

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/taskmanager/interfaces"
	"github.com/taskmanager/models"
)

// TODO 1: Define MemoryRepository struct que implemente TaskRepository
type MemoryRepository struct {
	// TODO: Agregar fields necesarios
	// tasks map[string]models.Task
	// mutex sync.RWMutex
}

// TODO 2: Implementa NewMemoryRepository constructor
func NewMemoryRepository() interfaces.TaskRepository {
	// TODO: Crear nuevo repositorio en memoria
	return nil
}

// TODO 3: Implementa Save para crear nueva tarea
func (r *MemoryRepository) Save(task models.Task) error {
	// TODO: Guardar tarea en memoria
	// Verificar que no exista una tarea con el mismo ID
	// Thread-safe usando mutex
	return nil
}

// TODO 4: Implementa GetByID para buscar por ID
func (r *MemoryRepository) GetByID(id string) (*models.Task, error) {
	// TODO: Buscar tarea por ID
	// Retornar TaskNotFoundError si no existe
	// Thread-safe para lectura
	return nil, nil
}

// TODO 5: Implementa GetAll para obtener todas las tareas
func (r *MemoryRepository) GetAll() ([]models.Task, error) {
	// TODO: Retornar todas las tareas como slice
	// Hacer copia para evitar modificaciones concurrentes
	return nil, nil
}

// TODO 6: Implementa Update para actualizar tarea existente
func (r *MemoryRepository) Update(task models.Task) error {
	// TODO: Actualizar tarea existente
	// Verificar que la tarea exista
	// Thread-safe usando mutex
	return nil
}

// TODO 7: Implementa Delete para eliminar tarea
func (r *MemoryRepository) Delete(id string) error {
	// TODO: Eliminar tarea por ID
	// Retornar TaskNotFoundError si no existe
	// Thread-safe usando mutex
	return nil
}

// TODO 8: Implementa FindByStatus para filtrar por estado
func (r *MemoryRepository) FindByStatus(status models.Status) ([]models.Task, error) {
	// TODO: Filtrar tareas por estado
	// Retornar slice de tareas que coincidan
	return nil, nil
}

// TODO 9: Implementa FindByPriority para filtrar por prioridad
func (r *MemoryRepository) FindByPriority(priority models.Priority) ([]models.Task, error) {
	// TODO: Filtrar tareas por prioridad
	return nil, nil
}

// TODO 10: Implementa Search para búsqueda de texto
func (r *MemoryRepository) Search(query string) ([]models.Task, error) {
	// TODO: Buscar en título y descripción (case-insensitive)
	// Retornar tareas que contengan el query
	return nil, nil
}

// TODO 11: Implementa FindByTag para filtrar por etiqueta
func (r *MemoryRepository) FindByTag(tag string) ([]models.Task, error) {
	// TODO: Filtrar tareas que contengan la etiqueta
	return nil, nil
}

// TODO 12: Implementa GetOverdue para tareas vencidas
func (r *MemoryRepository) GetOverdue() ([]models.Task, error) {
	// TODO: Retornar tareas vencidas usando IsOverdue()
	return nil, nil
}

// TODO 13: Implementa GetByDateRange para filtrar por rango de fechas
func (r *MemoryRepository) GetByDateRange(start, end time.Time) ([]models.Task, error) {
	// TODO: Filtrar tareas creadas entre start y end
	return nil, nil
}

// TODO 14: Implementa Count para contar total de tareas
func (r *MemoryRepository) Count() (int, error) {
	// TODO: Retornar número total de tareas
	return 0, nil
}

// TODO 15: Implementa CountByStatus para contar por estado
func (r *MemoryRepository) CountByStatus() (map[models.Status]int, error) {
	// TODO: Contar tareas por cada estado
	return nil, nil
}

// Métodos helper privados

// TODO 16: Implementa método privado copyTask para evitar modificaciones
func (r *MemoryRepository) copyTask(task models.Task) models.Task {
	// TODO: Crear copia profunda de la tarea
	return task.Clone()
}

// TODO 17: Implementa método privado copyTasks para slice
func (r *MemoryRepository) copyTasks(tasks []models.Task) []models.Task {
	// TODO: Crear copia del slice de tareas
	return nil
}

// TODO 18: Implementa método privado matchesQuery para búsqueda
func (r *MemoryRepository) matchesQuery(task models.Task, query string) bool {
	// TODO: Verificar si la tarea coincide con el query
	// Buscar en título, descripción y tags (case-insensitive)
	return false
}

// TODO 19: Implementa método privado isInDateRange
func (r *MemoryRepository) isInDateRange(taskDate, start, end time.Time) bool {
	// TODO: Verificar si la fecha está en el rango
	return false
}

// TODO 20: Implementa método privado validateTask
func (r *MemoryRepository) validateTask(task models.Task) error {
	// TODO: Validar tarea antes de operaciones
	return task.Validate()
}

// Implementación de interfaces adicionales

// TODO 21: Implementa Clear para limpiar todas las tareas (útil para testing)
func (r *MemoryRepository) Clear() error {
	// TODO: Limpiar todas las tareas
	return nil
}

// TODO 22: Implementa Size para obtener tamaño actual
func (r *MemoryRepository) Size() int {
	// TODO: Retornar número de tareas
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.tasks)
}

// TODO 23: Implementa Exists para verificar si existe una tarea
func (r *MemoryRepository) Exists(id string) bool {
	// TODO: Verificar si existe tarea con el ID
	return false
}

// TODO 24: Implementa GetIDs para obtener todos los IDs
func (r *MemoryRepository) GetIDs() []string {
	// TODO: Retornar slice con todos los IDs
	return nil
}

// TODO 25: Implementa FindByMultipleTags para múltiples etiquetas
func (r *MemoryRepository) FindByMultipleTags(tags []string, matchAll bool) ([]models.Task, error) {
	// TODO: Filtrar por múltiples tags
	// Si matchAll es true, la tarea debe tener TODAS las tags
	// Si es false, debe tener AL MENOS UNA tag
	return nil, nil
}

// TODO 26: Implementa GetRecentlyUpdated para tareas actualizadas recientemente
func (r *MemoryRepository) GetRecentlyUpdated(since time.Time) ([]models.Task, error) {
	// TODO: Retornar tareas actualizadas desde 'since'
	return nil, nil
}

// TODO 27: Implementa GetByPriorityAndStatus para filtro combinado
func (r *MemoryRepository) GetByPriorityAndStatus(priority models.Priority, status models.Status) ([]models.Task, error) {
	// TODO: Filtrar por prioridad Y estado
	return nil, nil
}

// TODO 28: Implementa GetDueSoon para tareas que vencen pronto
func (r *MemoryRepository) GetDueSoon(days int) ([]models.Task, error) {
	// TODO: Retornar tareas que vencen en los próximos 'days' días
	return nil, nil
}

// TODO 29: Implementa GetStatistics para estadísticas básicas
func (r *MemoryRepository) GetStatistics() (interfaces.TaskStatistics, error) {
	// TODO: Calcular y retornar estadísticas
	return interfaces.TaskStatistics{}, nil
}

// TODO 30: Implementa método de debugging String()
func (r *MemoryRepository) String() string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return fmt.Sprintf("MemoryRepository{tasks: %d}", len(r.tasks))
}