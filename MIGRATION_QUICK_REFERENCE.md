# Migration Quick Reference

## Running Migrations

### Automatic (Default)
Migrations run automatically when starting the application:
```bash
go run main.go
```

### With Makefile
```bash
make run              # Start app (with migrations)
make migrate-status   # Check migration status
make migrate-seed     # Seed sample data
```

## Creating New Migrations

### Step 1: Create Migration Files

Option A - Using script:
```bash
./migrate.sh create add_column_name
```

Option B - Manual:
```bash
# Create numbered migration files
touch migrations/000002_add_column_name.up.sql
touch migrations/000002_add_column_name.down.sql
```

### Step 2: Write SQL

**Up migration** (`000002_add_column_name.up.sql`):
```sql
ALTER TABLE todos ADD COLUMN new_column VARCHAR(255);
CREATE INDEX idx_new_column ON todos(new_column);
```

**Down migration** (`000002_add_column_name.down.sql`):
```sql
DROP INDEX idx_new_column ON todos;
ALTER TABLE todos DROP COLUMN new_column;
```

### Step 3: Run Application
```bash
go run main.go
# Migrations run automatically on startup
```

## Common Migration Tasks

### Add Column
```sql
-- Up
ALTER TABLE todos ADD COLUMN status VARCHAR(50) DEFAULT 'pending';

-- Down
ALTER TABLE todos DROP COLUMN status;
```

### Add Index
```sql
-- Up
CREATE INDEX idx_column_name ON todos(column_name);

-- Down
DROP INDEX idx_column_name ON todos;
```

### Rename Column
```sql
-- Up
ALTER TABLE todos CHANGE COLUMN old_name new_name VARCHAR(255);

-- Down
ALTER TABLE todos CHANGE COLUMN new_name old_name VARCHAR(255);
```

### Modify Column
```sql
-- Up
ALTER TABLE todos MODIFY COLUMN description TEXT NOT NULL;

-- Down
ALTER TABLE todos MODIFY COLUMN description TEXT;
```

### Add Foreign Key
```sql
-- Up
ALTER TABLE todos ADD CONSTRAINT fk_user_id 
FOREIGN KEY (user_id) REFERENCES users(id);

-- Down
ALTER TABLE todos DROP FOREIGN KEY fk_user_id;
```

## Checking Status

```bash
# Via Makefile
make migrate-status

# Via command line
go run main.go migrate -status

# Expected output:
# GORM AutoMigration enabled - migrations run automatically on startup
```

## Seeding Data

```bash
# Via Makefile
make migrate-seed

# Via command line
go run main.go migrate -seed
```

This creates sample todos automatically.

## Current Schema

```
TABLE: todos
├── id (VARCHAR 36) - PRIMARY KEY
├── title (VARCHAR 255) - NOT NULL
├── description (TEXT)
├── status (VARCHAR 50) - DEFAULT 'pending'
├── created_at (BIGINT) - NOT NULL
├── updated_at (BIGINT) - NOT NULL
├── INDEX: idx_status
└── INDEX: idx_created_at
```

## Migration Files Format

```
migrations/
├── 000001_create_todos_table.up.sql
├── 000001_create_todos_table.down.sql
├── 000002_add_priority_field.up.sql.example  (example, not executed)
├── 000002_add_priority_field.down.sql.example (example, not executed)
└── ...
```

**Rules:**
- Format: `{number}_{description}.{direction}.sql`
- Numbers: `000001`, `000002`, etc. (must be sequential)
- Direction: `.up.sql` (forward) or `.down.sql` (backward)
- Only `.up.sql` and `.down.sql` files are executed
- Files ending in `.example` are not executed

## Troubleshooting

### Database Connection Failed
```bash
# Check if MySQL is running
docker ps | grep chi-recap-mysql

# Start if needed
make db-up
```

### Migrations Not Running
```bash
# Check logs
make db-logs

# Restart and check
make db-down
make db-up
make run
```

### Reset Database
```bash
# Stop and remove database
make db-down
docker volume rm chi-recap_mysql_data

# Restart
make db-up
make run
```

## Best Practices

✅ DO:
- Create one migration per change
- Use descriptive names
- Test migrations locally first
- Keep migrations small and focused
- Write both up and down migrations
- Add comments to complex SQL

❌ DON'T:
- Modify existing migrations
- Mix unrelated changes
- Use AUTO_INCREMENT (use UUID)
- Add non-nullable columns without defaults
- Run complex transactions
