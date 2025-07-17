package interfaces

import "github.com/taskmanager/models"

// TODO 1: Define TaskRepository interface con todos los métodos necesarios
type TaskRepository interface {
	// TODO: Operaciones CRUD básicas
	// Save(task models.Task) error
	// GetByID(id string) (*models.Task, error)
	// GetAll() ([]models.Task, error)
	// Update(task models.Task) error
	// Delete(id string) error
	
	// TODO: Operaciones de búsqueda y filtrado
	// FindByStatus(status models.Status) ([]models.Task, error)
	// FindByPriority(priority models.Priority) ([]models.Task, error)
	// Search(query string) ([]models.Task, error)
	// FindByTag(tag string) ([]models.Task, error)
	
	// TODO: Operaciones avanzadas
	// GetOverdue() ([]models.Task, error)
	// GetByDateRange(start, end time.Time) ([]models.Task, error)
	// Count() (int, error)
	// CountByStatus() (map[models.Status]int, error)
}

// TODO 2: Define TaskStorage interface para persistencia
type TaskStorage interface {
	// TODO: Operaciones de almacenamiento
	// Load() ([]models.Task, error)
	// Save(tasks []models.Task) error
	// Backup() error
	// Restore() error
	// Clear() error
}

// TODO 3: Define TaskValidator interface para validación
type TaskValidator interface {
	// TODO: Validaciones específicas
	// ValidateTask(task models.Task) error
	// ValidateTitle(title string) error
	// ValidateDescription(description string) error
	// ValidateDueDate(dueDate *time.Time) error
}

// TODO 4: Define TaskSearcher interface para búsquedas complejas
type TaskSearcher interface {
	// TODO: Búsquedas avanzadas
	// SearchByMultipleCriteria(criteria SearchCriteria) ([]models.Task, error)
	// FindSimilar(task models.Task) ([]models.Task, error)
	// GetRecommendations(userID string) ([]models.Task, error)
}

// TODO 5: Define SearchCriteria struct para búsquedas complejas
type SearchCriteria struct {
	// TODO: Criterios de búsqueda
	// Query string
	// Status []models.Status
	// Priority []models.Priority
	// Tags []string
	// DateFrom *time.Time
	// DateTo *time.Time
	// DueDateFrom *time.Time
	// DueDateTo *time.Time
	// IncludeCompleted bool
	// SortBy string
	// SortOrder string
	// Limit int
	// Offset int
}

// TODO 6: Define TaskMetrics interface para estadísticas
type TaskMetrics interface {
	// TODO: Métricas y estadísticas
	// GetTaskStats() (TaskStatistics, error)
	// GetProductivityStats() (ProductivityStats, error)
	// GetTrendAnalysis() (TrendData, error)
	// GetCompletionRate() (float64, error)
}

// TODO 7: Define TaskStatistics struct
type TaskStatistics struct {
	// TODO: Estadísticas básicas
	// TotalTasks int
	// CompletedTasks int
	// PendingTasks int
	// InProgressTasks int
	// OverdueTasks int
	// TasksByPriority map[models.Priority]int
	// AverageCompletionTime time.Duration
	// MostUsedTags []string
}

// TODO 8: Define ProductivityStats struct
type ProductivityStats struct {
	// TODO: Estadísticas de productividad
	// TasksCreatedToday int
	// TasksCompletedToday int
	// WeeklyCompletion int
	// MonthlyCompletion int
	// ProductivityScore float64
	// BestDay string
	// WorstDay string
}

// TODO 9: Define TrendData struct
type TrendData struct {
	// TODO: Datos de tendencias
	// DailyTasks map[string]int
	// WeeklyTasks map[string]int
	// MonthlyTasks map[string]int
	// CompletionTrend []float64
	// PriorityTrend map[models.Priority][]int
}

// TODO 10: Define TaskExporter interface para import/export
type TaskExporter interface {
	// TODO: Operaciones de exportación
	// ExportToJSON(tasks []models.Task, filename string) error
	// ExportToCSV(tasks []models.Task, filename string) error
	// ImportFromJSON(filename string) ([]models.Task, error)
	// ImportFromCSV(filename string) ([]models.Task, error)
}

// TODO 11: Define TaskNotifier interface para notificaciones
type TaskNotifier interface {
	// TODO: Sistema de notificaciones
	// NotifyTaskDue(task models.Task) error
	// NotifyTaskOverdue(task models.Task) error
	// NotifyTaskCompleted(task models.Task) error
	// SetupReminders(tasks []models.Task) error
}

// TODO 12: Define TaskBackup interface para respaldos
type TaskBackup interface {
	// TODO: Sistema de respaldos
	// CreateBackup(filename string) error
	// RestoreFromBackup(filename string) error
	// ListBackups() ([]string, error)
	// DeleteBackup(filename string) error
	// AutoBackup(interval time.Duration) error
}

// TODO 13: Define TaskConfig interface para configuración
type TaskConfig interface {
	// TODO: Configuración del sistema
	// GetConfig() (Config, error)
	// SetConfig(config Config) error
	// ResetToDefaults() error
	// ValidateConfig(config Config) error
}

// TODO 14: Define Config struct
type Config struct {
	// TODO: Configuraciones del sistema
	// DataFile string
	// BackupEnabled bool
	// BackupInterval time.Duration
	// AutoSave bool
	// DefaultPriority models.Priority
	// DateFormat string
	// MaxTasks int
	// NotificationsEnabled bool
}

// TODO 15: Define TaskLogger interface para logging
type TaskLogger interface {
	// TODO: Sistema de logging
	// LogTaskCreated(task models.Task) error
	// LogTaskUpdated(oldTask, newTask models.Task) error
	// LogTaskDeleted(task models.Task) error
	// LogError(operation string, err error) error
	// GetLogs(limit int) ([]LogEntry, error)
}

// TODO 16: Define LogEntry struct
type LogEntry struct {
	// TODO: Entrada de log
	// Timestamp time.Time
	// Level string
	// Operation string
	// TaskID string
	// Message string
	// Error string
}

// TODO 17: Define TaskCache interface para caché
type TaskCache interface {
	// TODO: Sistema de caché
	// Get(key string) (*models.Task, bool)
	// Set(key string, task *models.Task, ttl time.Duration)
	// Delete(key string)
	// Clear()
	// Size() int
}

// TODO 18: Define TaskQuery interface para queries complejas
type TaskQuery interface {
	// TODO: Query builder
	// Where(field string, operator string, value interface{}) TaskQuery
	// OrderBy(field string, direction string) TaskQuery
	// Limit(limit int) TaskQuery
	// Offset(offset int) TaskQuery
	// Execute() ([]models.Task, error)
	// Count() (int, error)
}

// TODO 19: Define TaskService interface (para la capa de servicio)
type TaskService interface {
	// TODO: Métodos de servicio de alto nivel
	// CreateTask(title, description string, priority models.Priority) (*models.Task, error)
	// UpdateTask(id, title, description string) error
	// DeleteTask(id string) error
	// GetTask(id string) (*models.Task, error)
	// ListTasks() ([]models.Task, error)
	// SearchTasks(query string) ([]models.Task, error)
	// GetTasksByStatus(status models.Status) ([]models.Task, error)
	// MarkTaskComplete(id string) error
	// SetTaskPriority(id string, priority models.Priority) error
	// AddTaskTag(id, tag string) error
	// RemoveTaskTag(id, tag string) error
	// SetTaskDueDate(id string, dueDate time.Time) error
	// GetStatistics() (TaskStatistics, error)
}

// TODO 20: Define TaskManager interface (para el manager principal)
type TaskManager interface {
	// TODO: Interface principal que combine todo
	TaskService
	TaskMetrics
	TaskExporter
	TaskBackup
	TaskConfig
	
	// TODO: Métodos específicos del manager
	// Initialize() error
	// Shutdown() error
	// Health() error
}