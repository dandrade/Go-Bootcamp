package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// TODO 1: Crear instancia de TaskManager
	tm := NewTaskManager()
	if tm == nil {
		fmt.Println("Error: TaskManager no implementado")
		return
	}

	// Crear algunas tareas de ejemplo para testing
	createSampleTasks(tm)

	// Scanner para leer input del usuario
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Sistema de Gestión de Tareas v1.0 ===")
	fmt.Println("¡Bienvenido al Task Manager!")

	for {
		showMenu()
		fmt.Print("Selecciona una opción: ")
		
		if !scanner.Scan() {
			break
		}
		
		choice := strings.TrimSpace(scanner.Text())
		
		// TODO 2-10: Implementar cada case del menú
		switch choice {
		case "1":
			createTask(tm, scanner)
		case "2":
			listAllTasks(tm)
		case "3":
			viewTask(tm, scanner)
		case "4":
			markTaskComplete(tm, scanner)
		case "5":
			deleteTask(tm, scanner)
		case "6":
			searchTasks(tm, scanner)
		case "7":
			filterByStatus(tm, scanner)
		case "8":
			filterByPriority(tm, scanner)
		case "9":
			manageTaskTags(tm, scanner)
		case "10":
			setTaskDueDate(tm, scanner)
		case "11":
			viewStatistics(tm)
		case "12":
			advancedOptions(tm, scanner)
		case "0", "q", "quit", "salir":
			fmt.Println("¡Gracias por usar Task Manager! ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción inválida. Por favor intenta de nuevo.")
		}
		
		fmt.Println("\nPresiona Enter para continuar...")
		scanner.Scan()
		clearScreen()
	}
}

func showMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("               TASK MANAGER")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("1.  Crear nueva tarea")
	fmt.Println("2.  Listar todas las tareas")
	fmt.Println("3.  Ver tarea específica")
	fmt.Println("4.  Marcar tarea como completada")
	fmt.Println("5.  Eliminar tarea")
	fmt.Println("6.  Buscar tareas")
	fmt.Println("7.  Filtrar por estado")
	fmt.Println("8.  Filtrar por prioridad")
	fmt.Println("9.  Gestionar etiquetas")
	fmt.Println("10. Establecer fecha límite")
	fmt.Println("11. Ver estadísticas")
	fmt.Println("12. Opciones avanzadas")
	fmt.Println("0.  Salir")
	fmt.Println(strings.Repeat("=", 50))
}

// TODO 2: Implementar createTask
func createTask(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Crear Nueva Tarea ---")
	
	// TODO: Solicitar título (requerido)
	fmt.Print("Título: ")
	if !scanner.Scan() {
		return
	}
	title := strings.TrimSpace(scanner.Text())
	if title == "" {
		fmt.Println("Error: El título es requerido")
		return
	}
	
	// TODO: Solicitar descripción (opcional)
	fmt.Print("Descripción: ")
	if !scanner.Scan() {
		return
	}
	description := strings.TrimSpace(scanner.Text())
	
	// TODO: Solicitar prioridad
	fmt.Print("Prioridad (1=Baja, 2=Media, 3=Alta): ")
	if !scanner.Scan() {
		return
	}
	priorityStr := strings.TrimSpace(scanner.Text())
	priority, err := ParsePriority(priorityStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// TODO: Crear la tarea usando TaskManager
	task := tm.AddTask(title, description, priority)
	if task == nil {
		fmt.Println("Error: No se pudo crear la tarea")
		return
	}
	
	fmt.Printf("✅ Tarea creada exitosamente: %s\n", task.ID)
}

// TODO 3: Implementar listAllTasks
func listAllTasks(tm *TaskManager) {
	fmt.Println("\n--- Todas las Tareas ---")
	
	// TODO: Obtener todas las tareas
	tasks := tm.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println("No hay tareas registradas.")
		return
	}
	
	// TODO: Mostrar tareas en formato tabla
	printTasksTable(tasks)
}

// TODO 4: Implementar viewTask
func viewTask(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Ver Tarea Específica ---")
	
	// TODO: Solicitar ID de tarea
	fmt.Print("ID de la tarea: ")
	if !scanner.Scan() {
		return
	}
	id := strings.TrimSpace(scanner.Text())
	
	// TODO: Buscar y mostrar tarea
	task := tm.GetTask(id)
	if task == nil {
		fmt.Printf("Error: No se encontró tarea con ID '%s'\n", id)
		return
	}
	
	printTaskDetails(*task)
}

// TODO 5: Implementar markTaskComplete
func markTaskComplete(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Marcar Tarea como Completada ---")
	
	// TODO: Solicitar ID y marcar como completada
	fmt.Print("ID de la tarea: ")
	if !scanner.Scan() {
		return
	}
	id := strings.TrimSpace(scanner.Text())
	
	if tm.MarkTaskComplete(id) {
		fmt.Println("✅ Tarea marcada como completada")
	} else {
		fmt.Printf("Error: No se encontró tarea con ID '%s'\n", id)
	}
}

// TODO 6: Implementar deleteTask
func deleteTask(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Eliminar Tarea ---")
	
	// TODO: Solicitar ID y confirmar eliminación
	fmt.Print("ID de la tarea: ")
	if !scanner.Scan() {
		return
	}
	id := strings.TrimSpace(scanner.Text())
	
	// Confirmar eliminación
	fmt.Printf("¿Estás seguro de eliminar la tarea '%s'? (s/N): ", id)
	if !scanner.Scan() {
		return
	}
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
	
	if confirm == "s" || confirm == "si" || confirm == "y" || confirm == "yes" {
		if tm.DeleteTask(id) {
			fmt.Println("✅ Tarea eliminada exitosamente")
		} else {
			fmt.Printf("Error: No se encontró tarea con ID '%s'\n", id)
		}
	} else {
		fmt.Println("Eliminación cancelada")
	}
}

// TODO 7: Implementar searchTasks
func searchTasks(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Buscar Tareas ---")
	
	// TODO: Solicitar término de búsqueda
	fmt.Print("Término de búsqueda: ")
	if !scanner.Scan() {
		return
	}
	query := strings.TrimSpace(scanner.Text())
	if query == "" {
		fmt.Println("Error: Debes proporcionar un término de búsqueda")
		return
	}
	
	// TODO: Buscar y mostrar resultados
	tasks := tm.SearchTasks(query)
	if len(tasks) == 0 {
		fmt.Printf("No se encontraron tareas que contengan '%s'\n", query)
		return
	}
	
	fmt.Printf("Encontradas %d tarea(s):\n", len(tasks))
	printTasksTable(tasks)
}

// TODO 8: Implementar filterByStatus
func filterByStatus(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Filtrar por Estado ---")
	fmt.Println("1. Pendientes")
	fmt.Println("2. En Progreso")
	fmt.Println("3. Completadas")
	
	fmt.Print("Selecciona estado: ")
	if !scanner.Scan() {
		return
	}
	statusStr := strings.TrimSpace(scanner.Text())
	
	// TODO: Parsear estado y filtrar
	status, err := ParseStatus(statusStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	tasks := tm.GetTasksByStatus(status)
	if len(tasks) == 0 {
		fmt.Printf("No hay tareas con estado '%s'\n", status)
		return
	}
	
	fmt.Printf("Tareas con estado '%s':\n", status)
	printTasksTable(tasks)
}

// TODO 9: Implementar filterByPriority
func filterByPriority(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Filtrar por Prioridad ---")
	
	// TODO: Similar a filterByStatus pero para prioridad
	// Implementar función similar
}

// TODO 10: Implementar manageTaskTags
func manageTaskTags(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Gestionar Etiquetas ---")
	fmt.Println("1. Agregar etiqueta")
	fmt.Println("2. Remover etiqueta")
	fmt.Println("3. Ver tareas por etiqueta")
	
	// TODO: Implementar submenú para gestión de etiquetas
}

// setTaskDueDate establece fecha límite para una tarea
func setTaskDueDate(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Establecer Fecha Límite ---")
	
	fmt.Print("ID de la tarea: ")
	if !scanner.Scan() {
		return
	}
	id := strings.TrimSpace(scanner.Text())
	
	fmt.Print("Fecha límite (YYYY-MM-DD): ")
	if !scanner.Scan() {
		return
	}
	dateStr := strings.TrimSpace(scanner.Text())
	
	dueDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Error: Formato de fecha inválido: %v\n", err)
		return
	}
	
	if tm.SetTaskDueDate(id, dueDate) {
		fmt.Println("✅ Fecha límite establecida")
	} else {
		fmt.Printf("Error: No se encontró tarea con ID '%s'\n", id)
	}
}

// viewStatistics muestra estadísticas del sistema
func viewStatistics(tm *TaskManager) {
	fmt.Println("\n--- Estadísticas del Sistema ---")
	
	stats := tm.GetStatistics()
	
	fmt.Printf("Total de tareas: %d\n", stats.TotalTasks)
	fmt.Printf("Completadas: %d\n", stats.CompletedTasks)
	fmt.Printf("Pendientes: %d\n", stats.PendingTasks)
	fmt.Printf("En progreso: %d\n", stats.InProgressTasks)
	fmt.Printf("Vencidas: %d\n", stats.OverdueTasks)
	fmt.Printf("Tasa de completitud: %.1f%%\n", stats.CompletionRate)
	
	// Mostrar distribución por prioridad
	priorityCount := tm.GetTaskCountByPriority()
	fmt.Println("\nDistribución por prioridad:")
	for priority, count := range priorityCount {
		fmt.Printf("  %s: %d\n", priority, count)
	}
}

// advancedOptions menú de opciones avanzadas
func advancedOptions(tm *TaskManager, scanner *bufio.Scanner) {
	fmt.Println("\n--- Opciones Avanzadas ---")
	fmt.Println("1. Ver tareas vencidas")
	fmt.Println("2. Ver tareas por vencer pronto")
	fmt.Println("3. Ordenar por prioridad")
	fmt.Println("4. Ordenar por fecha límite")
	fmt.Println("5. Limpiar tareas completadas")
	
	// TODO: Implementar submenú avanzado
}

// Funciones helper para mostrar información

func printTasksTable(tasks []Task) {
	if len(tasks) == 0 {
		return
	}
	
	// Header
	fmt.Printf("%-12s | %-25s | %-12s | %-8s | %-12s\n", 
		"ID", "Título", "Estado", "Prioridad", "Fecha Límite")
	fmt.Println(strings.Repeat("-", 80))
	
	// Tareas
	for _, task := range tasks {
		title := task.Title
		if len(title) > 25 {
			title = title[:22] + "..."
		}
		
		dueDate := task.GetFormattedDueDate()
		if len(dueDate) > 12 {
			dueDate = dueDate[:9] + "..."
		}
		
		fmt.Printf("%-12s | %-25s | %-12s | %-8s | %-12s\n",
			task.ID, title, task.Status, task.Priority, dueDate)
	}
}

func printTaskDetails(task Task) {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("ID: %s\n", task.ID)
	fmt.Printf("Título: %s\n", task.Title)
	fmt.Printf("Descripción: %s\n", task.Description)
	fmt.Printf("Estado: %s\n", task.Status)
	fmt.Printf("Prioridad: %s\n", task.Priority)
	fmt.Printf("Creada: %s\n", task.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("Actualizada: %s\n", task.UpdatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("Fecha límite: %s\n", task.GetFormattedDueDate())
	fmt.Printf("Etiquetas: %s\n", task.GetTagsString())
	
	if task.IsOverdue() {
		fmt.Println("⚠️  TAREA VENCIDA")
	}
	fmt.Println(strings.Repeat("-", 50))
}

func clearScreen() {
	// Simple clear para hacer la interfaz más limpia
	fmt.Print("\033[2J\033[H")
}

// createSampleTasks crea tareas de ejemplo para testing
func createSampleTasks(tm *TaskManager) {
	// Solo crear si no hay tareas
	if tm.GetTaskCount() > 0 {
		return
	}
	
	// Crear algunas tareas de ejemplo
	task1 := tm.AddTask("Revisar documentación de Go", "Leer docs oficiales y tutoriales", Medium)
	if task1 != nil {
		task1.AddTag("programación")
		task1.AddTag("go")
		dueDate := time.Now().AddDate(0, 0, 7) // En 7 días
		task1.SetDueDate(dueDate)
	}
	
	task2 := tm.AddTask("Implementar API REST", "Crear endpoints para el proyecto", High)
	if task2 != nil {
		task2.AddTag("desarrollo")
		task2.AddTag("api")
		task2.SetStatus(InProgress)
	}
	
	task3 := tm.AddTask("Escribir tests unitarios", "Crear tests para todas las funciones", Medium)
	if task3 != nil {
		task3.AddTag("testing")
		task3.AddTag("calidad")
	}
	
	fmt.Println("📝 Se han creado algunas tareas de ejemplo para comenzar")
}