package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// TODO 1: Define el struct TaskManager
type TaskManager struct {
	// TODO: Implementar campos
	// tasks []Task
	// nextID int (para IDs secuenciales si prefieres)
}

// TODO 2: Implementa NewTaskManager como constructor
func NewTaskManager() *TaskManager {
	// TODO: Inicializar TaskManager con slice vacío
	return nil
}

// TODO 3: Implementa AddTask que cree y agregue una nueva tarea
func (tm *TaskManager) AddTask(title, description string, priority Priority) *Task {
	// TODO: Crear nueva tarea y agregarla al slice
	// Retornar puntero a la tarea creada
	return nil
}

// TODO 4: Implementa GetTask que busque una tarea por ID
func (tm *TaskManager) GetTask(id string) *Task {
	// TODO: Buscar tarea por ID y retornar puntero o nil
	return nil
}

// TODO 5: Implementa GetAllTasks que retorne copia de todas las tareas
func (tm *TaskManager) GetAllTasks() []Task {
	// TODO: Retornar copia del slice de tareas
	return nil
}

// TODO 6: Implementa GetTasksByStatus que filtre por estado
func (tm *TaskManager) GetTasksByStatus(status Status) []Task {
	// TODO: Filtrar tareas por estado y retornar slice
	return nil
}

// TODO 7: Implementa GetTasksByPriority que filtre por prioridad
func (tm *TaskManager) GetTasksByPriority(priority Priority) []Task {
	// TODO: Filtrar tareas por prioridad
	return nil
}

// TODO 8: Implementa DeleteTask que elimine una tarea por ID
func (tm *TaskManager) DeleteTask(id string) bool {
	// TODO: Buscar y eliminar tarea, retornar true si se eliminó
	return false
}

// TODO 9: Implementa SearchTasks que busque en título y descripción
func (tm *TaskManager) SearchTasks(query string) []Task {
	// TODO: Buscar en título y descripción (case-insensitive)
	return nil
}

// TODO 10: Implementa GetTasksByTag que filtre por etiqueta
func (tm *TaskManager) GetTasksByTag(tag string) []Task {
	// TODO: Filtrar tareas que contengan la etiqueta
	return nil
}

// TODO 11: Implementa GetOverdueTasks que retorne tareas vencidas
func (tm *TaskManager) GetOverdueTasks() []Task {
	// TODO: Filtrar tareas vencidas usando IsOverdue()
	return nil
}

// TODO 12: Implementa GetTasksDueSoon que retorne tareas por vencer (próximos N días)
func (tm *TaskManager) GetTasksDueSoon(days int) []Task {
	// TODO: Filtrar tareas que vencen en los próximos N días
	return nil
}

// TODO 13: Implementa GetTaskCount que retorne número total de tareas
func (tm *TaskManager) GetTaskCount() int {
	// TODO: Retornar cantidad de tareas
	return 0
}

// TODO 14: Implementa GetTaskCountByStatus que cuente por estado
func (tm *TaskManager) GetTaskCountByStatus() map[Status]int {
	// TODO: Contar tareas por cada estado
	return nil
}

// TODO 15: Implementa GetTaskCountByPriority que cuente por prioridad
func (tm *TaskManager) GetTaskCountByPriority() map[Priority]int {
	// TODO: Contar tareas por cada prioridad
	return nil
}

// TODO 16: Implementa SortTasksByPriority que ordene por prioridad (High -> Low)
func (tm *TaskManager) SortTasksByPriority() []Task {
	// TODO: Crear copia, ordenar por prioridad y retornar
	return nil
}

// TODO 17: Implementa SortTasksByDueDate que ordene por fecha límite
func (tm *TaskManager) SortTasksByDueDate() []Task {
	// TODO: Ordenar por fecha límite (sin fecha al final)
	return nil
}

// TODO 18: Implementa SortTasksByCreatedDate que ordene por fecha de creación
func (tm *TaskManager) SortTasksByCreatedDate() []Task {
	// TODO: Ordenar por CreatedAt (más recientes primero)
	return nil
}

// TODO 19: Implementa MarkTaskComplete que marque una tarea como completada
func (tm *TaskManager) MarkTaskComplete(id string) bool {
	// TODO: Buscar tarea y marcarla como completada
	return false
}

// TODO 20: Implementa UpdateTaskStatus que actualice el estado de una tarea
func (tm *TaskManager) UpdateTaskStatus(id string, status Status) bool {
	// TODO: Buscar tarea y actualizar estado
	return false
}

// TODO 21: Implementa UpdateTaskPriority que actualice la prioridad
func (tm *TaskManager) UpdateTaskPriority(id string, priority Priority) bool {
	// TODO: Buscar tarea y actualizar prioridad
	return false
}

// TODO 22: Implementa SetTaskDueDate que establezca fecha límite
func (tm *TaskManager) SetTaskDueDate(id string, dueDate time.Time) bool {
	// TODO: Buscar tarea y establecer fecha límite
	return false
}

// TODO 23: Implementa AddTaskTag que agregue etiqueta a una tarea
func (tm *TaskManager) AddTaskTag(id string, tag string) bool {
	// TODO: Buscar tarea y agregar etiqueta
	return false
}

// TODO 24: Implementa RemoveTaskTag que remueva etiqueta de una tarea
func (tm *TaskManager) RemoveTaskTag(id string, tag string) bool {
	// TODO: Buscar tarea y remover etiqueta
	return false
}

// TODO 25: Implementa GetAllTags que retorne todas las etiquetas únicas
func (tm *TaskManager) GetAllTags() []string {
	// TODO: Recopilar todas las etiquetas únicas de todas las tareas
	return nil
}

// TODO 26: Implementa GetStatistics que retorne estadísticas generales
type TaskStatistics struct {
	TotalTasks      int
	CompletedTasks  int
	PendingTasks    int
	InProgressTasks int
	OverdueTasks    int
	CompletionRate  float64 // Porcentaje de tareas completadas
}

func (tm *TaskManager) GetStatistics() TaskStatistics {
	// TODO: Calcular y retornar estadísticas
	return TaskStatistics{}
}

// TODO 27: Implementa ClearCompletedTasks que elimine todas las tareas completadas
func (tm *TaskManager) ClearCompletedTasks() int {
	// TODO: Eliminar tareas completadas y retornar cantidad eliminada
	return 0
}

// TODO 28: Implementa UpdateTask que actualice título y descripción
func (tm *TaskManager) UpdateTask(id, title, description string) bool {
	// TODO: Buscar tarea y actualizar campos
	return false
}

// Funciones helper para ordenamiento (ejemplo de implementación)
func sortTasksByPriority(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority // High(2) > Medium(1) > Low(0)
	})
}

func sortTasksByDueDate(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		// Tareas sin fecha van al final
		if tasks[i].DueDate == nil && tasks[j].DueDate == nil {
			return false
		}
		if tasks[i].DueDate == nil {
			return false
		}
		if tasks[j].DueDate == nil {
			return true
		}
		return tasks[i].DueDate.Before(*tasks[j].DueDate)
	})
}

func sortTasksByCreatedDate(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt.After(tasks[j].CreatedAt) // Más recientes primero
	})
}