package services

import (
	"chi-recap/internal/domain/shared"
	"chi-recap/internal/domain/todo"
)

// CreateTodoRequest is the request to create a new todo
type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodoResponse is the response for a todo
type TodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// UpdateTodoRequest is the request to update a todo
type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// TodoService handles todo business logic
type TodoService struct {
	repository todo.Repository
}

// NewTodoService creates a new TodoService
func NewTodoService(repository todo.Repository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

// CreateTodo creates a new todo
func (s *TodoService) CreateTodo(req CreateTodoRequest) (*TodoResponse, error) {
	newTodo := todo.NewTodo(req.Title, req.Description)
	
	if err := s.repository.Save(newTodo); err != nil {
		return nil, err
	}

	return s.domainToResponse(newTodo), nil
}

// GetTodoByID retrieves a todo by ID
func (s *TodoService) GetTodoByID(id string) (*TodoResponse, error) {
	todoID := shared.NewIDFromString(id)
	t, err := s.repository.GetByID(todoID)
	if err != nil {
		return nil, err
	}

	return s.domainToResponse(t), nil
}

// GetAllTodos retrieves all todos
func (s *TodoService) GetAllTodos() ([]*TodoResponse, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]*TodoResponse, len(todos))
	for i, t := range todos {
		responses[i] = s.domainToResponse(t)
	}

	return responses, nil
}

// UpdateTodo updates an existing todo
func (s *TodoService) UpdateTodo(id string, req UpdateTodoRequest) (*TodoResponse, error) {
	todoID := shared.NewIDFromString(id)
	t, err := s.repository.GetByID(todoID)
	if err != nil {
		return nil, err
	}

	t.UpdateTitle(req.Title)
	t.UpdateDescription(req.Description)

	if req.Status == "done" {
		t.MarkAsDone()
	} else if req.Status == "pending" {
		t.MarkAsPending()
	}

	if err := s.repository.Update(t); err != nil {
		return nil, err
	}

	return s.domainToResponse(t), nil
}

// DeleteTodo deletes a todo
func (s *TodoService) DeleteTodo(id string) error {
	todoID := shared.NewIDFromString(id)
	return s.repository.Delete(todoID)
}

// MarkTodoAsDone marks a todo as done
func (s *TodoService) MarkTodoAsDone(id string) (*TodoResponse, error) {
	todoID := shared.NewIDFromString(id)
	t, err := s.repository.GetByID(todoID)
	if err != nil {
		return nil, err
	}

	t.MarkAsDone()

	if err := s.repository.Update(t); err != nil {
		return nil, err
	}

	return s.domainToResponse(t), nil
}

// domainToResponse converts a domain Todo to a response
func (s *TodoService) domainToResponse(t *todo.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
}
