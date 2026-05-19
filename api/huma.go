package api

import (
	"context"
	"net/http"

	"chi-recap/internal/application/services"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// TodoCreateRequest is the request to create a todo
type TodoCreateRequest struct {
	Title       string `json:"title" doc:"Todo title" maxLength:"255" example:"Learn Go"`
	Description string `json:"description" doc:"Todo description" example:"Master the Go programming language"`
}

// TodoResponse is the todo response
type TodoResponse struct {
	ID          string `json:"id" doc:"Unique identifier (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title       string `json:"title" doc:"Todo title" example:"Learn Go"`
	Description string `json:"description" doc:"Todo description"`
	Status      string `json:"status" doc:"Todo status (pending or done)" enum:"pending,done" example:"pending"`
	CreatedAt   int64  `json:"created_at" doc:"Creation timestamp (Unix)" example:"1716193200"`
	UpdatedAt   int64  `json:"updated_at" doc:"Last update timestamp (Unix)" example:"1716193200"`
}

// TodoUpdateRequest is the request to update a todo
type TodoUpdateRequest struct {
	Title       string `json:"title" doc:"Todo title" maxLength:"255" example:"Learn Go"`
	Description string `json:"description" doc:"Todo description"`
	Status      string `json:"status" doc:"Todo status (pending or done)" enum:"pending,done" example:"done"`
}

// TodoListResponse is the list of todos
type TodoListResponse struct {
	Todos []TodoResponse `json:"todos" doc:"List of todos"`
}

// HealthResponse is the health check response
type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// SetupHumaAPI configures Huma API with documentation
func SetupHumaAPI(todoService *services.TodoService) (chi.Router, huma.API) {
	router := chi.NewMux()
	config := huma.DefaultConfig("Todo API", "1.0.0")
	config.Info.Description = "A complete REST API for managing todos using Domain-Driven Design patterns"
	config.Info.Contact = &huma.Contact{
		Name:  "Chi-Recap Project",
		Email: "support@chi-recap.local",
		URL:   "https://github.com/repoleved08/chi-recap",
	}
	config.Info.License = &huma.License{
		Name: "MIT",
		URL:  "https://opensource.org/licenses/MIT",
	}

	api := humachi.New(router, config)

	// Health check endpoint
	huma.Get(api, "/health", func(ctx context.Context, input *struct{}) (*HealthResponse, error) {
		return &HealthResponse{Status: "ok"}, nil
	})

	// Create todo
	huma.Post(api, "/todos", func(ctx context.Context, input *struct {
		Body TodoCreateRequest
	}) (*TodoResponse, error) {
		resp, err := todoService.CreateTodo(services.CreateTodoRequest{
			Title:       input.Body.Title,
			Description: input.Body.Description,
		})
		if err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusInternalServerError,
				Title:  "Failed to create todo",
				Detail: err.Error(),
			}
		}

		return &TodoResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			Status:      resp.Status,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}, nil
	})

	// List all todos
	huma.Get(api, "/todos", func(ctx context.Context, input *struct{}) (*TodoListResponse, error) {
		todos, err := todoService.GetAllTodos()
		if err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusInternalServerError,
				Title:  "Failed to list todos",
				Detail: err.Error(),
			}
		}

		var todoResponses []TodoResponse
		for _, todo := range todos {
			todoResponses = append(todoResponses, TodoResponse{
				ID:          todo.ID,
				Title:       todo.Title,
				Description: todo.Description,
				Status:      todo.Status,
				CreatedAt:   todo.CreatedAt,
				UpdatedAt:   todo.UpdatedAt,
			})
		}

		return &TodoListResponse{Todos: todoResponses}, nil
	})

	// Get single todo
	huma.Get(api, "/todos/{id}", func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"Todo ID (UUID)"`
	}) (*TodoResponse, error) {
		resp, err := todoService.GetTodoByID(input.ID)
		if err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusNotFound,
				Title:  "Todo not found",
				Detail: err.Error(),
			}
		}

		return &TodoResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			Status:      resp.Status,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}, nil
	})

	// Update todo
	huma.Put(api, "/todos/{id}", func(ctx context.Context, input *struct {
		ID   string `path:"id" doc:"Todo ID (UUID)"`
		Body TodoUpdateRequest
	}) (*TodoResponse, error) {
		resp, err := todoService.UpdateTodo(input.ID, services.UpdateTodoRequest{
			Title:       input.Body.Title,
			Description: input.Body.Description,
			Status:      input.Body.Status,
		})
		if err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusInternalServerError,
				Title:  "Failed to update todo",
				Detail: err.Error(),
			}
		}

		return &TodoResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			Status:      resp.Status,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}, nil
	})

	// Delete todo
	huma.Delete(api, "/todos/{id}", func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"Todo ID (UUID)"`
	}) (*struct{}, error) {
		if err := todoService.DeleteTodo(input.ID); err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusInternalServerError,
				Title:  "Failed to delete todo",
				Detail: err.Error(),
			}
		}
		return &struct{}{}, nil
	})

	// Mark todo as done
	huma.Post(api, "/todos/{id}/done", func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"Todo ID (UUID)"`
	}) (*TodoResponse, error) {
		resp, err := todoService.MarkTodoAsDone(input.ID)
		if err != nil {
			return nil, &huma.ErrorModel{
				Status: http.StatusInternalServerError,
				Title:  "Failed to mark todo as done",
				Detail: err.Error(),
			}
		}

		return &TodoResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			Status:      resp.Status,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}, nil
	})

	return router, api
}
