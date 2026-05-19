package main

import (
	"flag"
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
	// Parse command line flags
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	migrateStatus := migrateCmd.Bool("status", false, "Show current migration status")
	migrateSeed := migrateCmd.Bool("seed", false, "Seed the database with sample data")

	// Database configuration
	dbConfig := config.DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		Username: getEnv("DB_USERNAME", "root"),
		Password: getEnv("DB_PASSWORD", "rootpass"),
		Database: getEnv("DB_NAME", "chi_recap"),
	}

	// Check if migration commands are provided
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrateCmd.Parse(os.Args[2:])

		// Initialize database connection
		db, err := config.NewDatabaseConnection(dbConfig)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if *migrateStatus {
			migrationManager := config.NewMigrationManager(db)
			log.Println(migrationManager.GetStatus())
		} else if *migrateSeed {
			// Seed with sample data
			seedDatabase(db)
		} else {
			migrateCmd.Usage()
		}
		return
	}

	// Initialize database connection
	db, err := config.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations automatically on startup
	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	todoRepo := persistence.NewTodoRepository(db)

	// Initialize services
	todoService := services.NewTodoService(todoRepo)

	// Setup Huma API with documentation
	router, _ := api.SetupHumaAPI(todoService)

	// Start server
	port := getEnv("PORT", "8888")
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	log.Printf("✓ Database connected and migrated successfully")
	log.Printf("✓ Huma API configured")
	log.Printf("Starting server on %s", addr)
	log.Printf("📚 API Documentation available at http://%s/docs", addr)
	log.Printf("📖 OpenAPI Schema available at http://%s/openapi.json", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// runMigrations runs database migrations using GORM AutoMigrate
func runMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")
	return db.AutoMigrate(&persistence.TodoModel{})
}

// seedDatabase seeds the database with sample data
func seedDatabase(db *gorm.DB) error {
	log.Println("Seeding database with sample data...")

	// First run migrations
	if err := runMigrations(db); err != nil {
		return err
	}

	// Sample todos
	todos := []persistence.TodoModel{
		{
			ID:          "1",
			Title:       "Learn Go",
			Description: "Master Go programming language",
			Status:      "done",
		},
		{
			ID:          "2",
			Title:       "Build REST API",
			Description: "Create a RESTful API with Chi and GORM",
			Status:      "in_progress",
		},
		{
			ID:          "3",
			Title:       "Deploy to production",
			Description: "Deploy the application to a production server",
			Status:      "pending",
		},
	}

	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			log.Printf("Warning: Failed to seed todo %s: %v", todo.ID, err)
		}
	}

	log.Println("✓ Database seeded successfully")
	return nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
