package todo

import (
	"time"

	"chi-recap/internal/domain/shared"
)

// Status represents the status of a todo
type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)

// Todo is the aggregate root for the Todo domain
type Todo struct {
	ID        *shared.ID
	Title     string
	Description string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTodo creates a new Todo
func NewTodo(title, description string) *Todo {
	return &Todo{
		ID:          shared.NewID(),
		Title:       title,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// MarkAsDone marks the todo as done
func (t *Todo) MarkAsDone() {
	t.Status = StatusDone
	t.UpdatedAt = time.Now()
}

// MarkAsPending marks the todo as pending
func (t *Todo) MarkAsPending() {
	t.Status = StatusPending
	t.UpdatedAt = time.Now()
}

// UpdateTitle updates the title of the todo
func (t *Todo) UpdateTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

// UpdateDescription updates the description of the todo
func (t *Todo) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}
