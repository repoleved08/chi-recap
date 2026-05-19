package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"chi-recap/api"
	"chi-recap/internal/application/services"
	"chi-recap/internal/infrastructure/config"
	"chi-recap/internal/infrastructure/persistence"

	"gorm.io/gorm"
)

func main() {
	// Database configuration
	dbConfig := config.DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		Username: getEnv("DB_USERNAME", "root"),
		Password: getEnv("DB_PASSWORD", "rootpass"),
		Database: getEnv("DB_NAME", "chi_recap"),
	}

	// Initialize database connection
	db, err := config.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	todoRepo := persistence.NewTodoRepository(db)

	// Initialize services
	todoService := services.NewTodoService(todoRepo)

	// Setup routes
	router := api.SetupRoutes(todoService)

	// Start server
	port := getEnv("PORT", "8888")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// runMigrations runs database migrations
func runMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&persistence.TodoModel{})
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
