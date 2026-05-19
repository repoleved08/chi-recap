package api

import (
	"net/http"

	"chi-recap/api/handlers"
	"chi-recap/internal/application/services"

	"github.com/go-chi/chi/v5"
)

// SetupRoutes sets up all the API routes
func SetupRoutes(todoService *services.TodoService) chi.Router {
	router := chi.NewRouter()

	todoHandler := handlers.NewTodoHandler(todoService)

	// Todo routes
	router.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.CreateTodo)
		r.Get("/", todoHandler.ListTodos)
		
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", todoHandler.GetTodo)
			r.Put("/", todoHandler.UpdateTodo)
			r.Delete("/", todoHandler.DeleteTodo)
			r.Post("/done", todoHandler.MarkAsDone)
		})
	})

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return router
}
