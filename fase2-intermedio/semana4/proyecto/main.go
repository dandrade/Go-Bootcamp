package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/taskmanager/interfaces"
	"github.com/taskmanager/models"
	"github.com/taskmanager/repository"
	"github.com/taskmanager/service"
)

var (
	taskService interfaces.TaskService
	scanner     *bufio.Scanner
)

func main() {
	fmt.Println("=== Task Manager v2.0 - Interfaces & Error Handling ===")
	
	// TODO 1: Inicializar el sistema
	if err := initializeSystem(); err != nil {
		fmt.Printf("Error inicializando sistema: %v\n", err)
		return
	}
	
	scanner = bufio.NewScanner(os.Stdin)
	
	// TODO 2: Crear algunas tareas de ejemplo si es la primera vez
	createSampleTasksIfNeeded()
	
	fmt.Println("¡Bienvenido al Sistema de Gestión de Tareas v2.0!")
	
	// TODO 3: Loop principal del menú
	for {
		showMainMenu()
		choice := getInput("Selecciona una opción: ")
		
		if !handleMenuChoice(choice) {
			break
		}
		
		pauseForUser()
		clearScreen()
	}
	
	fmt.Println("¡Gracias por usar Task Manager v2.0! ¡Hasta luego!")
}

// TODO 4: Implementa initializeSystem
func initializeSystem() error {
	// TODO: Configurar repositorio y service
	// Permitir elegir entre memoria y JSON
	fmt.Println("Configurando sistema...")
	fmt.Println("1. Usar almacenamiento en memoria (temporal)")
	fmt.Println("2. Usar almacenamiento en archivo JSON (persistente)")
	
	choice := getInput("Selecciona tipo de almacenamiento (1-2): ")
	
	var repo interfaces.TaskRepository
	var err error
	
	switch choice {
	case "1":
		// TODO: Crear repositorio en memoria
		repo = repository.NewMemoryRepository()
		fmt.Println("✓ Usando almacenamiento en memoria")
	case "2", "":
		// TODO: Crear repositorio JSON
		repo = repository.NewJSONRepository("tasks.json")
		fmt.Println("✓ Usando almacenamiento en archivo JSON")
	default:
		// TODO: Default a JSON
		repo = repository.NewJSONRepository("tasks.json")
		fmt.Println("✓ Opción inválida, usando almacenamiento JSON por defecto")
	}
	
	// TODO: Crear service
	taskService = service.NewTaskService(repo)
	
	return err
}

func showMainMenu() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("                    TASK MANAGER v2.0")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("📝 TAREAS")
	fmt.Println("  1.  Crear nueva tarea")
	fmt.Println("  2.  Listar todas las tareas")
	fmt.Println("  3.  Ver tarea específica")
	fmt.Println("  4.  Actualizar tarea")
	fmt.Println("  5.  Cambiar estado de tarea")
	fmt.Println("  6.  Eliminar tarea")
	fmt.Println("\n🔍 BÚSQUEDA Y FILTROS")
	fmt.Println("  7.  Buscar tareas")
	fmt.Println("  8.  Filtrar por estado")
	fmt.Println("  9.  Filtrar por prioridad")
	fmt.Println("  10. Tareas vencidas")
	fmt.Println("  11. Tareas por vencer")
	fmt.Println("\n🏷️  GESTIÓN")
	fmt.Println("  12. Gestionar etiquetas")
	fmt.Println("  13. Establecer fechas límite")
	fmt.Println("  14. Operaciones en lote")
	fmt.Println("\n📊 INFORMACIÓN")
	fmt.Println("  15. Ver estadísticas")
	fmt.Println("  16. Estadísticas de productividad")
	fmt.Println("\n⚙️  SISTEMA")
	fmt.Println("  17. Archivar tareas completadas")
	fmt.Println("  18. Verificar salud del sistema")
	fmt.Println("\n  0.  Salir")
	fmt.Println(strings.Repeat("=", 60))
}

// TODO 5: Implementa handleMenuChoice
func handleMenuChoice(choice string) bool {
	switch choice {
	case "1":
		createNewTask()
	case "2":
		listAllTasks()
	case "3":
		viewSpecificTask()
	case "4":
		updateTask()
	case "5":
		changeTaskStatus()
	case "6":
		deleteTask()
	case "7":
		searchTasks()
	case "8":
		filterByStatus()
	case "9":
		filterByPriority()
	case "10":
		showOverdueTasks()
	case "11":
		showTasksDueSoon()
	case "12":
		manageTaskTags()
	case "13":
		manageDueDates()
	case "14":
		bulkOperations()
	case "15":
		showStatistics()
	case "16":
		showProductivityStats()
	case "17":
		archiveCompletedTasks()
	case "18":
		checkSystemHealth()
	case "0", "q", "quit", "salir":
		return false
	default:
		fmt.Println("❌ Opción inválida. Por favor intenta de nuevo.")
	}
	return true
}

// TODO 6: Implementa createNewTask con manejo de errores
func createNewTask() {
	fmt.Println("\n--- Crear Nueva Tarea ---")
	
	title := getInput("Título: ")
	if title == "" {
		fmt.Println("❌ El título es requerido")
		return
	}
	
	description := getInput("Descripción (opcional): ")
	
	fmt.Println("Prioridad:")
	fmt.Println("  1. Baja")
	fmt.Println("  2. Media")
	fmt.Println("  3. Alta")
	priorityStr := getInput("Selecciona prioridad (1-3): ")
	
	priority, err := models.ParsePriority(priorityStr)
	if err != nil {
		handleError("Error de validación", err)
		return
	}
	
	task, err := taskService.CreateTask(title, description, priority)
	if err != nil {
		handleError("Error creando tarea", err)
		return
	}
	
	fmt.Printf("✅ Tarea creada exitosamente: %s\n", task.ID)
	fmt.Printf("📋 %s\n", task.String())
}

// TODO 7: Implementa listAllTasks con formato mejorado
func listAllTasks() {
	fmt.Println("\n--- Todas las Tareas ---")
	
	tasks, err := taskService.ListTasks()
	if err != nil {
		handleError("Error obteniendo tareas", err)
		return
	}
	
	if len(tasks) == 0 {
		fmt.Println("📋 No hay tareas registradas.")
		return
	}
	
	printTasksTable(tasks)
}

// TODO 8: Implementa viewSpecificTask con detalles completos
func viewSpecificTask() {
	fmt.Println("\n--- Ver Tarea Específica ---")
	
	id := getInput("ID de la tarea: ")
	if id == "" {
		fmt.Println("❌ ID es requerido")
		return
	}
	
	task, err := taskService.GetTask(id)
	if err != nil {
		handleError("Error obteniendo tarea", err)
		return
	}
	
	printTaskDetails(*task)
}

// TODO 9: Implementa updateTask
func updateTask() {
	fmt.Println("\n--- Actualizar Tarea ---")
	
	id := getInput("ID de la tarea: ")
	if id == "" {
		fmt.Println("❌ ID es requerido")
		return
	}
	
	// Mostrar tarea actual
	currentTask, err := taskService.GetTask(id)
	if err != nil {
		handleError("Error obteniendo tarea", err)
		return
	}
	
	fmt.Println("\nTarea actual:")
	printTaskDetails(*currentTask)
	
	fmt.Println("\nNuevos valores (Enter para mantener actual):")
	newTitle := getInputWithDefault("Título", currentTask.Title)
	newDescription := getInputWithDefault("Descripción", currentTask.Description)
	
	err = taskService.UpdateTask(id, newTitle, newDescription)
	if err != nil {
		handleError("Error actualizando tarea", err)
		return
	}
	
	fmt.Println("✅ Tarea actualizada exitosamente")
}

// TODO 10: Implementa changeTaskStatus
func changeTaskStatus() {
	fmt.Println("\n--- Cambiar Estado de Tarea ---")
	
	id := getInput("ID de la tarea: ")
	if id == "" {
		fmt.Println("❌ ID es requerido")
		return
	}
	
	fmt.Println("Estados disponibles:")
	fmt.Println("  1. Pendiente")
	fmt.Println("  2. En Progreso") 
	fmt.Println("  3. Completada")
	
	statusStr := getInput("Selecciona nuevo estado (1-3): ")
	
	if statusStr == "3" {
		// Usar método específico para completar
		err := taskService.MarkTaskComplete(id)
		if err != nil {
			handleError("Error completando tarea", err)
			return
		}
		fmt.Println("✅ Tarea marcada como completada")
	} else {
		status, err := models.ParseStatus(statusStr)
		if err != nil {
			handleError("Error de validación", err)
			return
		}
		
		err = taskService.SetTaskPriority(id, models.Priority(status)) // TODO: Esto necesita arreglo
		if err != nil {
			handleError("Error cambiando estado", err)
			return
		}
		fmt.Println("✅ Estado actualizado exitosamente")
	}
}

// TODO 11: Implementa deleteTask con confirmación
func deleteTask() {
	fmt.Println("\n--- Eliminar Tarea ---")
	
	id := getInput("ID de la tarea: ")
	if id == "" {
		fmt.Println("❌ ID es requerido")
		return
	}
	
	// Mostrar tarea antes de eliminar
	task, err := taskService.GetTask(id)
	if err != nil {
		handleError("Error obteniendo tarea", err)
		return
	}
	
	fmt.Println("\nTarea a eliminar:")
	printTaskDetails(*task)
	
	confirmation := getInput("\n¿Estás seguro de eliminar esta tarea? (s/N): ")
	if strings.ToLower(confirmation) != "s" && strings.ToLower(confirmation) != "si" {
		fmt.Println("❌ Eliminación cancelada")
		return
	}
	
	err = taskService.DeleteTask(id)
	if err != nil {
		handleError("Error eliminando tarea", err)
		return
	}
	
	fmt.Println("✅ Tarea eliminada exitosamente")
}

// TODO 12: Implementa searchTasks
func searchTasks() {
	fmt.Println("\n--- Buscar Tareas ---")
	
	query := getInput("Término de búsqueda: ")
	if query == "" {
		fmt.Println("❌ Debes proporcionar un término de búsqueda")
		return
	}
	
	tasks, err := taskService.SearchTasks(query)
	if err != nil {
		handleError("Error en búsqueda", err)
		return
	}
	
	if len(tasks) == 0 {
		fmt.Printf("🔍 No se encontraron tareas que contengan '%s'\n", query)
		return
	}
	
	fmt.Printf("🔍 Encontradas %d tarea(s):\n", len(tasks))
	printTasksTable(tasks)
}

// TODO 13: Implementa filterByStatus
func filterByStatus() {
	fmt.Println("\n--- Filtrar por Estado ---")
	fmt.Println("  1. Pendientes")
	fmt.Println("  2. En Progreso")
	fmt.Println("  3. Completadas")
	
	statusStr := getInput("Selecciona estado (1-3): ")
	status, err := models.ParseStatus(statusStr)
	if err != nil {
		handleError("Error de validación", err)
		return
	}
	
	tasks, err := taskService.GetTasksByStatus(status)
	if err != nil {
		handleError("Error filtrando tareas", err)
		return
	}
	
	if len(tasks) == 0 {
		fmt.Printf("📋 No hay tareas con estado '%s'\n", status)
		return
	}
	
	fmt.Printf("📋 Tareas con estado '%s':\n", status)
	printTasksTable(tasks)
}

// TODO 14-25: Implementa resto de funciones del menú
func filterByPriority() {
	// TODO: Implementar filtro por prioridad
	fmt.Println("🚧 Función en desarrollo...")
}

func showOverdueTasks() {
	// TODO: Mostrar tareas vencidas
	fmt.Println("🚧 Función en desarrollo...")
}

func showTasksDueSoon() {
	// TODO: Mostrar tareas que vencen pronto
	fmt.Println("🚧 Función en desarrollo...")
}

func manageTaskTags() {
	// TODO: Gestión de etiquetas
	fmt.Println("🚧 Función en desarrollo...")
}

func manageDueDates() {
	// TODO: Gestión de fechas límite
	fmt.Println("🚧 Función en desarrollo...")
}

func bulkOperations() {
	// TODO: Operaciones en lote
	fmt.Println("🚧 Función en desarrollo...")
}

func showStatistics() {
	fmt.Println("\n--- Estadísticas del Sistema ---")
	
	stats, err := taskService.GetStatistics()
	if err != nil {
		handleError("Error obteniendo estadísticas", err)
		return
	}
	
	fmt.Printf("📊 Total de tareas: %d\n", stats.TotalTasks)
	fmt.Printf("✅ Completadas: %d\n", stats.CompletedTasks) 
	fmt.Printf("⏳ Pendientes: %d\n", stats.PendingTasks)
	fmt.Printf("🔄 En progreso: %d\n", stats.InProgressTasks)
	fmt.Printf("⚠️  Vencidas: %d\n", stats.OverdueTasks)
	
	if stats.TotalTasks > 0 {
		completionRate := float64(stats.CompletedTasks) / float64(stats.TotalTasks) * 100
		fmt.Printf("📈 Tasa de completitud: %.1f%%\n", completionRate)
	}
}

func showProductivityStats() {
	// TODO: Estadísticas de productividad
	fmt.Println("🚧 Función en desarrollo...")
}

func archiveCompletedTasks() {
	// TODO: Archivar tareas completadas
	fmt.Println("🚧 Función en desarrollo...")
}

func checkSystemHealth() {
	// TODO: Verificar salud del sistema
	fmt.Println("🚧 Función en desarrollo...")
}

// Funciones helper

func handleError(context string, err error) {
	fmt.Printf("❌ %s: ", context)
	
	// TODO: Manejo específico por tipo de error
	switch e := err.(type) {
	case models.TaskNotFoundError:
		fmt.Printf("Tarea '%s' no encontrada\n", e.ID)
	case models.ValidationError:
		fmt.Printf("Error de validación en campo '%s': %s\n", e.Field, e.Message)
	case models.MultiValidationError:
		fmt.Println("Múltiples errores de validación:")
		for _, valErr := range e.GetErrors() {
			fmt.Printf("  - %s: %s\n", valErr.Field, valErr.Message)
		}
	case models.PersistenceError:
		fmt.Printf("Error de almacenamiento (%s): %v\n", e.Operation, e.Cause)
	default:
		fmt.Printf("%v\n", err)
	}
}

func printTasksTable(tasks []models.Task) {
	if len(tasks) == 0 {
		return
	}
	
	// Header
	fmt.Printf("%-12s | %-30s | %-12s | %-8s | %-12s | %-8s\n", 
		"ID", "Título", "Estado", "Prioridad", "Fecha Límite", "Etiquetas")
	fmt.Println(strings.Repeat("-", 90))
	
	// Tareas
	for _, task := range tasks {
		title := task.Title
		if len(title) > 30 {
			title = title[:27] + "..."
		}
		
		dueDate := task.GetFormattedDueDate()
		if len(dueDate) > 12 {
			dueDate = dueDate[:9] + "..."
		}
		
		tags := task.GetTagsString()
		if len(tags) > 8 {
			tags = tags[:5] + "..."
		}
		
		// Agregar emojis según estado
		status := task.Status.String()
		switch task.Status {
		case models.Pending:
			status = "⏳ " + status
		case models.InProgress:
			status = "🔄 " + status
		case models.Completed:
			status = "✅ " + status
		}
		
		// Agregar emoji si está vencida
		if task.IsOverdue() {
			dueDate = "⚠️ " + dueDate
		}
		
		fmt.Printf("%-12s | %-30s | %-12s | %-8s | %-12s | %-8s\n",
			task.ID, title, status, task.Priority, dueDate, tags)
	}
}

func printTaskDetails(task models.Task) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("📋 ID: %s\n", task.ID)
	fmt.Printf("📝 Título: %s\n", task.Title)
	fmt.Printf("📄 Descripción: %s\n", task.Description)
	
	status := task.Status.String()
	switch task.Status {
	case models.Pending:
		status = "⏳ " + status
	case models.InProgress:
		status = "🔄 " + status
	case models.Completed:
		status = "✅ " + status
	}
	fmt.Printf("📊 Estado: %s\n", status)
	
	priority := task.Priority.String()
	switch task.Priority {
	case models.Low:
		priority = "🟢 " + priority
	case models.Medium:
		priority = "🟡 " + priority
	case models.High:
		priority = "🔴 " + priority
	}
	fmt.Printf("⚡ Prioridad: %s\n", priority)
	
	fmt.Printf("📅 Creada: %s\n", task.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("📝 Actualizada: %s\n", task.UpdatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("⏰ Fecha límite: %s\n", task.GetFormattedDueDate())
	fmt.Printf("🏷️  Etiquetas: %s\n", task.GetTagsString())
	
	if task.IsOverdue() {
		fmt.Println("⚠️  TAREA VENCIDA")
	} else if task.DueDate != nil {
		days := task.DaysUntilDue()
		if days > 0 {
			fmt.Printf("⏱️  Vence en %d día(s)\n", days)
		}
	}
	
	fmt.Println(strings.Repeat("-", 60))
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

func getInputWithDefault(prompt, defaultValue string) string {
	if defaultValue != "" {
		prompt = fmt.Sprintf("%s [%s]: ", prompt, defaultValue)
	} else {
		prompt = fmt.Sprintf("%s: ", prompt)
	}
	
	input := getInput(prompt)
	if input == "" {
		return defaultValue
	}
	return input
}

func pauseForUser() {
	fmt.Print("\nPresiona Enter para continuar...")
	scanner.Scan()
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func createSampleTasksIfNeeded() {
	// TODO: Verificar si hay tareas, si no crear ejemplos
	tasks, err := taskService.ListTasks()
	if err != nil || len(tasks) > 0 {
		return
	}
	
	fmt.Println("🎯 Creando tareas de ejemplo...")
	
	// Crear tareas de ejemplo
	samples := []struct {
		title       string
		description string
		priority    models.Priority
	}{
		{"Estudiar Go interfaces", "Revisar material teórico y completar ejercicios", models.High},
		{"Implementar proyecto final", "Completar sistema de gestión de tareas", models.High},
		{"Revisar documentación", "Leer docs oficiales de Go", models.Medium},
		{"Practicar error handling", "Completar ejercicios de manejo de errores", models.Medium},
	}
	
	for _, sample := range samples {
		task, err := taskService.CreateTask(sample.title, sample.description, sample.priority)
		if err == nil && len(samples) > 2 {
			// Marcar algunas como en progreso
			if sample.priority == models.Medium {
				// TODO: Cambiar estado (necesita implementación completa)
			}
		}
	}
	
	fmt.Println("✅ Tareas de ejemplo creadas")
}