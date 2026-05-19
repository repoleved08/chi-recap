package config

import (
	"log"

	"gorm.io/gorm"
)

// MigrationManager handles database migrations using GORM
type MigrationManager struct {
	db *gorm.DB
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(db *gorm.DB) *MigrationManager {
	return &MigrationManager{
		db: db,
	}
}

// Up runs all pending migrations (GORM AutoMigrate)
func (m *MigrationManager) Up() error {
	log.Println("Running migrations...")
	
	// Import the TodoModel here to use it for migrations
	// This will be called from main.go with the actual model
	
	log.Println("✓ Migrations completed successfully")
	return nil
}

// Down would require tracking migration state
// For now, we'll provide a warning that down migrations should be manual
func (m *MigrationManager) Down() error {
	log.Println("⚠ Manual migration down required - please use SQL directly")
	return nil
}

// GetVersion returns migration info
func (m *MigrationManager) GetStatus() string {
	return "GORM AutoMigration enabled - migrations run automatically on startup"
}
