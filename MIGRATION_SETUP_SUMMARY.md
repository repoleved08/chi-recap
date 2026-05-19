# Database Migration System - Complete Setup Summary

## What Was Implemented

Your project now has a complete database migration system with two approaches:

### 1. **Automatic Migrations (GORM AutoMigrate)**
- ✅ Runs automatically on application startup
- ✅ Creates/updates tables automatically
- ✅ Best for development and testing
- ✅ No configuration needed

### 2. **Manual SQL Migrations**
- ✅ SQL files in `migrations/` directory
- ✅ Explicit control over schema changes
- ✅ Better for production environments
- ✅ Full version control of migrations

## Migration Structure

```
migrations/
├── 000001_create_todos_table.up.sql       ✅ Active
├── 000001_create_todos_table.down.sql     ✅ Active
├── 000002_add_priority_field.up.sql.example    (example only)
└── 000002_add_priority_field.down.sql.example  (example only)
```

**Naming Convention:** `{number}_{description}.{direction}.sql`
- Numbers must be sequential: 000001, 000002, 000003, etc.
- Directions: `.up.sql` (forward) or `.down.sql` (backward)
- Only files without `.example` are executed

## Quick Start

### Start Database & Run Migrations
```bash
# Option 1: Using Makefile (recommended)
make run

# Option 2: Using docker-compose + go
docker-compose up -d
go run main.go

# Option 3: Build and run binary
go build -o chi-recap
./chi-recap
```

**What happens:**
1. ✅ Database container starts
2. ✅ Connection established
3. ✅ Automatic migrations run
4. ✅ Server starts on port 8888

### Check Migration Status
```bash
make migrate-status
# Output: GORM AutoMigration enabled - migrations run automatically on startup
```

### Seed Sample Data
```bash
make migrate-seed
# Creates 3 sample todos
```

## Using Migrations

### View Current Database Schema

```bash
# Connect to database
make db-shell

# View todos table structure
mysql> DESCRIBE todos;

# View existing data
mysql> SELECT * FROM todos;
mysql> SELECT COUNT(*) FROM todos;
```

### Create New Migration

**Step 1: Generate files**
```bash
./migrate.sh create add_tags_column
# Creates:
# - migrations/000002_add_tags_column.up.sql
# - migrations/000002_add_tags_column.down.sql
```

**Step 2: Edit migration files**

`migrations/000002_add_tags_column.up.sql`:
```sql
ALTER TABLE todos ADD COLUMN tags VARCHAR(255);
CREATE INDEX idx_tags ON todos(tags);
```

`migrations/000002_add_tags_column.down.sql`:
```sql
DROP INDEX idx_tags ON todos;
ALTER TABLE todos DROP COLUMN tags;
```

**Step 3: Run application**
```bash
go run main.go
# Migration runs automatically
```

## Database Schema

### Current Tables

**todos** table:
```
Column      | Type          | Constraints
------------|---------------|-------------------
id          | VARCHAR(36)   | PRIMARY KEY
title       | VARCHAR(255)  | NOT NULL
description | TEXT          | 
status      | VARCHAR(50)   | DEFAULT 'pending'
created_at  | BIGINT        | NOT NULL
updated_at  | BIGINT        | NOT NULL

Indices:
- idx_status
- idx_created_at
```

## Useful Commands

### Development
```bash
make build              # Build binary
make run                # Run with migrations
make test               # Run tests
make clean              # Remove binary
```

### Database
```bash
make db-up              # Start MySQL container
make db-down            # Stop MySQL container
make db-logs            # View database logs
make db-shell           # Connect to MySQL
```

### Migrations
```bash
make migrate-status     # Show migration status
make migrate-seed       # Seed sample data
./migrate.sh create     # Create new migration
./migrate.sh help       # Migration script help
```

## File Locations

```
chi-recap/
├── main.go                              # Entry point with migration logic
├── migrations/                          # Migration files
│   ├── 000001_create_todos_table.up.sql
│   └── 000001_create_todos_table.down.sql
├── MIGRATIONS.md                        # Detailed migration guide
├── MIGRATION_QUICK_REFERENCE.md         # Quick reference
├── migrate.sh                           # Migration script
├── Makefile                             # Make commands
├── docker-compose.yml                   # MySQL setup
└── internal/infrastructure/config/
    └── migration.go                     # Migration manager
```

## How Migrations Work

### When You Start the Application

```bash
go run main.go
```

**Process:**
1. Parse command-line flags
2. Connect to database
3. Initialize migration manager
4. Run `db.AutoMigrate(&persistence.TodoModel{})` 
5. GORM automatically:
   - Creates table if missing
   - Adds new columns
   - Creates indices
   - Updates schema
6. Server starts
7. Ready to accept requests

### Migration Status During Startup

```
Running database migrations...
✓ Database connected and migrated successfully
Starting server on 0.0.0.0:8888
```

## Managing Migrations

### View All Migrations
```bash
ls -la migrations/
# Shows all .up.sql and .down.sql files (numbered)
```

### Check What's in Database
```bash
make db-shell

# Then in MySQL:
SHOW TABLES;
DESCRIBE todos;
SELECT * FROM todos LIMIT 10;
```

### Reset Everything
```bash
# Remove database volume
make db-down
docker volume rm chi-recap_mysql_data

# Restart fresh
make db-up
make run
```

## Migration Examples

### Example 1: Add a New Column

Create migration:
```bash
./migrate.sh create add_priority_to_todos
```

Edit `000002_add_priority_to_todos.up.sql`:
```sql
ALTER TABLE todos ADD COLUMN priority INT DEFAULT 1;
CREATE INDEX idx_priority ON todos(priority);
```

Edit `000002_add_priority_to_todos.down.sql`:
```sql
DROP INDEX idx_priority ON todos;
ALTER TABLE todos DROP COLUMN priority;
```

### Example 2: Create New Table

Edit `000003_create_users_table.up.sql`:
```sql
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at BIGINT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

Edit `000003_create_users_table.down.sql`:
```sql
DROP TABLE IF EXISTS users;
```

## Best Practices

✅ **DO:**
- One migration per logical change
- Use descriptive names
- Test locally first
- Keep migrations small
- Write both up and down
- Add comments for complex SQL

❌ **DON'T:**
- Modify existing migrations
- Mix unrelated changes
- Skip down migrations
- Use AUTO_INCREMENT
- Add non-nullable without defaults

## Troubleshooting

### "Connection refused"
```bash
# Check if database is running
docker ps | grep chi-recap-mysql

# Start it
make db-up
```

### Migrations won't run
```bash
# Check logs
make db-logs

# Restart everything
make db-down
make db-up
make run
```

### Need to rollback
```bash
# See current status
make migrate-status

# Reset database
docker volume rm chi-recap_mysql_data
make db-up
make run
```

## Next Steps

1. **Create Migrations**: Use `./migrate.sh create` for schema changes
2. **Test Locally**: Always test on development database first
3. **Version Control**: Commit migration files to git
4. **Production**: Document all migrations for deployment

## Additional Resources

- **MIGRATIONS.md** - Comprehensive migration guide
- **MIGRATION_QUICK_REFERENCE.md** - Quick reference card
- **migration.go** - Migration manager implementation
- **main.go** - Application entry point with migration logic

## Support

For help with specific migrations:
```bash
./migrate.sh help
make help
```

For detailed documentation:
- Read MIGRATIONS.md for complete guide
- Check MIGRATION_QUICK_REFERENCE.md for quick answers
- Look at migration examples in migrations/ directory
