package api

import (
	"chi-recap/internal/application/services"

	"github.com/go-chi/chi/v5"
)

// SetupRoutes sets up all the API routes (deprecated, use SetupHumaAPI instead)
func SetupRoutes(todoService *services.TodoService) chi.Router {
	// This function is kept for backward compatibility
	// Use SetupHumaAPI instead for full documentation support
	router, _ := SetupHumaAPI(todoService)
	return router
}
