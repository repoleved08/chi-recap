# Database Migration Guide

This document explains how to manage database migrations for the chi-recap project.

## Overview

The project uses two migration approaches:

1. **Automatic Migrations (GORM AutoMigrate)** - Runs automatically on application startup
2. **Manual SQL Migrations** - SQL files in the `migrations/` directory for complex changes

## Automatic Migrations (Recommended for Development)

GORM's `AutoMigrate` is enabled by default and will:
- Create tables if they don't exist
- Add missing columns
- Create indices defined in model tags
- Run automatically when the application starts

### Running the Application

```bash
go run main.go
```

This will:
1. Connect to the database
2. Automatically run migrations
3. Start the server

**Output:**
```
Running database migrations...
✓ Database connected and migrated successfully
Starting server on 0.0.0.0:8888
```

## Manual SQL Migrations (For Production)

For production environments, use SQL migration files for explicit control.

### Migration File Structure

Migration files are located in the `migrations/` directory with naming convention:

```
migrations/
├── 000001_create_todos_table.up.sql
├── 000001_create_todos_table.down.sql
├── 000002_add_priority_column.up.sql
└── 000002_add_priority_column.down.sql
```

- **Sequence**: `000001`, `000002`, etc.
- **Suffix**: `.up.sql` (forward) and `.down.sql` (backward)
- **Format**: `{number}_{description}.{direction}.sql`

### Create a New Migration

#### Option 1: Using the Migration Script

```bash
./migrate.sh create add_priority_field
```

This creates:
- `migrations/000002_add_priority_field.up.sql`
- `migrations/000002_add_priority_field.down.sql`

#### Option 2: Manual Creation

```bash
# Create migration files
touch migrations/000002_add_priority_field.up.sql
touch migrations/000002_add_priority_field.down.sql
```

### Writing Migration SQL

**Up Migration** (`000002_add_priority_field.up.sql`):
```sql
-- Add priority field to todos table
ALTER TABLE todos ADD COLUMN priority INT DEFAULT 0 AFTER status;
CREATE INDEX idx_priority ON todos(priority);
```

**Down Migration** (`000002_add_priority_field.down.sql`):
```sql
-- Remove priority field from todos table
DROP INDEX idx_priority ON todos;
ALTER TABLE todos DROP COLUMN priority;
```

## Checking Migration Status

```bash
go run main.go migrate -status
```

Output:
```
GORM AutoMigration enabled - migrations run automatically on startup
```

## Seeding the Database

### With Sample Data

```bash
go run main.go migrate -seed
```

This will:
1. Run all pending migrations
2. Insert sample todo data

Sample data created:
- "Learn Go" - Status: done
- "Build REST API" - Status: in_progress
- "Deploy to production" - Status: pending

## Current Schema

### Todos Table

```sql
CREATE TABLE todos (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

**Columns:**
- `id` - UUID (primary key)
- `title` - Todo title
- `description` - Detailed description
- `status` - `pending` or `done`
- `created_at` - Unix timestamp
- `updated_at` - Unix timestamp

**Indices:**
- `idx_status` - For filtering by status
- `idx_created_at` - For sorting by creation date

## Migration Best Practices

### ✅ DO:

- ✓ Create one migration per schema change
- ✓ Use descriptive migration names
- ✓ Test migrations on a development database first
- ✓ Include both `.up.sql` and `.down.sql` files
- ✓ Keep migrations small and focused
- ✓ Add indices for frequently queried columns
- ✓ Include comments in SQL files

### ❌ DON'T:

- ✗ Modify existing migration files (create new ones instead)
- ✗ Use AUTO_INCREMENT for IDs (use UUID instead)
- ✗ Add non-nullable columns without defaults
- ✗ Run migrations manually in MySQL - use the application

## Troubleshooting

### Migrations Not Running

```bash
# Check if database is running
docker ps | grep chi-recap-mysql

# If not running:
docker-compose up -d
```

### Database Connection Error

```bash
# Verify database credentials in .env
cat .env | grep DB_

# Test connection manually
docker exec -it chi-recap-mysql mysql -u root -p -e "USE chi_recap; SHOW TABLES;"
```

### Need to Rollback Everything

```bash
# Using migration script
./migrate.sh fresh

# Or manually
docker exec chi-recap-mysql mysql -u root -proot -e "DROP DATABASE chi_recap; CREATE DATABASE chi_recap;"
```

### Manual SQL Execution

If needed, execute SQL directly in the database:

```bash
# Connect to MySQL
docker exec -it chi-recap-mysql mysql -u root -proot chi_recap

# Execute migration SQL
source migrations/000001_create_todos_table.up.sql
```

## Integration with CI/CD

For automated deployments:

```bash
#!/bin/bash
# deploy.sh

# Set environment variables
export DB_HOST=prod-db.example.com
export DB_PORT=3306
export DB_USERNAME=prod_user
export DB_PASSWORD=$DB_PASSWORD  # From secrets

# Run migrations
go run main.go

# If using external migration tool:
# migrate -path ./migrations -database "mysql://..." up
```

## Future Enhancements

The project can be extended to use advanced migration tools:

1. **golang-migrate/migrate** - Standalone migration tool
2. **Flyway** - Language-agnostic migration tool
3. **Liquibase** - Version control for databases

These would provide:
- Explicit migration versioning
- Better rollback capabilities
- Multi-database support
- Transaction management

## Environment Variables

Migrations use these variables from `.env`:

```
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=rootpass
DB_NAME=chi_recap
```

Modify these before running migrations for different environments.
