package persistence

import (
	"time"

	"chi-recap/internal/domain/shared"
	"chi-recap/internal/domain/todo"

	"gorm.io/gorm"
)

// TodoModel is the GORM model for Todo
type TodoModel struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string
	CreatedAt   int64
	UpdatedAt   int64
}

// TableName specifies the table name
func (TodoModel) TableName() string {
	return "todos"
}

// TodoRepository implements the todo.Repository interface
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new TodoRepository
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// Save persists a new todo
func (r *TodoRepository) Save(t *todo.Todo) error {
	model := &TodoModel{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
	return r.db.Create(model).Error
}

// GetByID retrieves a todo by its ID
func (r *TodoRepository) GetByID(id *shared.ID) (*todo.Todo, error) {
	var model TodoModel
	if err := r.db.Where("id = ?", id.String()).First(&model).Error; err != nil {
		return nil, err
	}
	return r.modelToDomain(&model), nil
}

// GetAll retrieves all todos
func (r *TodoRepository) GetAll() ([]*todo.Todo, error) {
	var models []TodoModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	var todos []*todo.Todo
	for _, model := range models {
		todos = append(todos, r.modelToDomain(&model))
	}
	return todos, nil
}

// Delete removes a todo
func (r *TodoRepository) Delete(id *shared.ID) error {
	return r.db.Where("id = ?", id.String()).Delete(&TodoModel{}).Error
}

// Update updates an existing todo
func (r *TodoRepository) Update(t *todo.Todo) error {
	model := &TodoModel{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
	return r.db.Model(&TodoModel{}).Where("id = ?", t.ID.String()).Updates(model).Error
}

// modelToDomain converts a TodoModel to a domain Todo
func (r *TodoRepository) modelToDomain(model *TodoModel) *todo.Todo {
	return &todo.Todo{
		ID:          shared.NewIDFromString(model.ID),
		Title:       model.Title,
		Description: model.Description,
		Status:      todo.Status(model.Status),
		CreatedAt:   time.Unix(model.CreatedAt, 0),
		UpdatedAt:   time.Unix(model.UpdatedAt, 0),
	}
}
