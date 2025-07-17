package models

import (
	"fmt"
	"strings"
	"time"
)

// Priority representa la prioridad de una tarea
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

// String implementa fmt.Stringer para Priority
func (p Priority) String() string {
	switch p {
	case Low:
		return "Baja"
	case Medium:
		return "Media"
	case High:
		return "Alta"
	default:
		return "Desconocida"
	}
}

// Status representa el estado de una tarea
type Status int

const (
	Pending Status = iota
	InProgress
	Completed
)

// String implementa fmt.Stringer para Status
func (s Status) String() string {
	switch s {
	case Pending:
		return "Pendiente"
	case InProgress:
		return "En Progreso"
	case Completed:
		return "Completada"
	default:
		return "Desconocido"
	}
}

// Task representa una tarea del sistema
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    Priority   `json:"priority"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}

// NewTask crea una nueva tarea con validación
func NewTask(title, description string, priority Priority) (*Task, error) {
	// Generar ID único
	id := generateID()
	now := time.Now()
	
	task := &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      Pending,
		CreatedAt:   now,
		UpdatedAt:   now,
		Tags:        make([]string, 0),
	}
	
	// Validar la tarea antes de retornarla
	if err := task.Validate(); err != nil {
		return nil, err
	}
	
	return task, nil
}

// Validate valida los campos de la tarea
func (t *Task) Validate() error {
	errors := NewMultiValidationError()
	
	// Validar ID
	if strings.TrimSpace(t.ID) == "" {
		errors.AddError(FieldID, t.ID, "no puede estar vacío")
	}
	
	// Validar título
	if strings.TrimSpace(t.Title) == "" {
		errors.AddError(FieldTitle, t.Title, "no puede estar vacío")
	} else if len(t.Title) > 100 {
		errors.AddError(FieldTitle, t.Title, "no puede tener más de 100 caracteres")
	}
	
	// Validar descripción (opcional pero con límite)
	if len(t.Description) > 500 {
		errors.AddError(FieldDescription, t.Description, "no puede tener más de 500 caracteres")
	}
	
	// Validar prioridad
	if t.Priority < Low || t.Priority > High {
		errors.AddError("priority", t.Priority, "debe ser Low, Medium o High")
	}
	
	// Validar estado
	if t.Status < Pending || t.Status > Completed {
		errors.AddError("status", t.Status, "debe ser Pending, InProgress o Completed")
	}
	
	// Validar fecha límite (debe ser futura si está presente)
	if t.DueDate != nil && t.DueDate.Before(time.Now().Truncate(24*time.Hour)) {
		errors.AddError("due_date", t.DueDate, "no puede ser en el pasado")
	}
	
	// Validar tags
	for i, tag := range t.Tags {
		if strings.TrimSpace(tag) == "" {
			errors.AddError(fmt.Sprintf("tags[%d]", i), tag, "no puede estar vacío")
		}
	}
	
	if errors.HasErrors() {
		return *errors
	}
	
	return nil
}

// String implementa fmt.Stringer
func (t Task) String() string {
	return fmt.Sprintf("%s: %s (%s, %s)", t.ID, t.Title, t.Status, t.Priority)
}

// IsOverdue verifica si la tarea está vencida
func (t Task) IsOverdue() bool {
	return t.DueDate != nil && t.DueDate.Before(time.Now()) && t.Status != Completed
}

// MarkComplete marca la tarea como completada
func (t *Task) MarkComplete() {
	t.Status = Completed
	t.UpdatedAt = time.Now()
}

// SetStatus cambia el estado de la tarea
func (t *Task) SetStatus(status Status) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// SetPriority cambia la prioridad de la tarea
func (t *Task) SetPriority(priority Priority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

// SetDueDate establece la fecha límite
func (t *Task) SetDueDate(dueDate time.Time) {
	t.DueDate = &dueDate
	t.UpdatedAt = time.Now()
}

// ClearDueDate quita la fecha límite
func (t *Task) ClearDueDate() {
	t.DueDate = nil
	t.UpdatedAt = time.Now()
}

// AddTag agrega una etiqueta única
func (t *Task) AddTag(tag string) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return
	}
	
	// Verificar que no exista
	for _, existingTag := range t.Tags {
		if existingTag == tag {
			return
		}
	}
	
	t.Tags = append(t.Tags, tag)
	t.UpdatedAt = time.Now()
}

// RemoveTag remueve una etiqueta
func (t *Task) RemoveTag(tag string) bool {
	for i, existingTag := range t.Tags {
		if existingTag == tag {
			t.Tags = append(t.Tags[:i], t.Tags[i+1:]...)
			t.UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

// HasTag verifica si tiene una etiqueta
func (t Task) HasTag(tag string) bool {
	for _, existingTag := range t.Tags {
		if existingTag == tag {
			return true
		}
	}
	return false
}

// GetFormattedDueDate retorna la fecha límite formateada
func (t Task) GetFormattedDueDate() string {
	if t.DueDate == nil {
		return "Sin fecha límite"
	}
	return t.DueDate.Format("2006-01-02")
}

// GetTagsString retorna las etiquetas como string
func (t Task) GetTagsString() string {
	if len(t.Tags) == 0 {
		return "Sin etiquetas"
	}
	return strings.Join(t.Tags, ", ")
}

// DaysUntilDue calcula días hasta la fecha límite
func (t Task) DaysUntilDue() int {
	if t.DueDate == nil {
		return 0
	}
	
	now := time.Now().Truncate(24 * time.Hour)
	due := t.DueDate.Truncate(24 * time.Hour)
	
	if due.Before(now) {
		return 0
	}
	
	return int(due.Sub(now).Hours() / 24)
}

// Update actualiza título y descripción con validación
func (t *Task) Update(title, description string) error {
	oldTitle := t.Title
	oldDescription := t.Description
	
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
	
	if err := t.Validate(); err != nil {
		// Revertir cambios si la validación falla
		t.Title = oldTitle
		t.Description = oldDescription
		return err
	}
	
	return nil
}

// Clone crea una copia de la tarea
func (t Task) Clone() Task {
	clone := t
	
	// Copiar slice de tags
	if len(t.Tags) > 0 {
		clone.Tags = make([]string, len(t.Tags))
		copy(clone.Tags, t.Tags)
	}
	
	// Copiar DueDate si existe
	if t.DueDate != nil {
		dueDate := *t.DueDate
		clone.DueDate = &dueDate
	}
	
	return clone
}

// generateID genera un ID único basado en timestamp
func generateID() string {
	return fmt.Sprintf("T%d", time.Now().UnixNano())
}

// ParsePriority parsea string a Priority
func ParsePriority(input string) (Priority, error) {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "1", "baja", "low", "l":
		return Low, nil
	case "2", "media", "medium", "m":
		return Medium, nil
	case "3", "alta", "high", "h":
		return High, nil
	default:
		return Low, NewValidationError("priority", input, "debe ser 'baja', 'media' o 'alta'")
	}
}

// ParseStatus parsea string a Status
func ParseStatus(input string) (Status, error) {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "1", "pendiente", "pending", "p":
		return Pending, nil
	case "2", "progreso", "progress", "inprogress", "i":
		return InProgress, nil
	case "3", "completada", "completed", "done", "c":
		return Completed, nil
	default:
		return Pending, NewValidationError("status", input, "debe ser 'pendiente', 'progreso' o 'completada'")
	}
}