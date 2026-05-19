# Quick Start Guide

## Step 1: Start MySQL Database

```bash
docker-compose up -d
```

Wait for the database to be ready (check with):
```bash
docker-compose logs mysql | grep "ready for connections"
```

## Step 2: Run the Application

```bash
go run main.go
```

You should see:
```
Starting server on 0.0.0.0:8888
```

## Step 3: Test the API

### Create a Todo
```bash
curl -X POST http://localhost:8888/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn DDD",
    "description": "Study Domain-Driven Design patterns"
  }'
```

### List All Todos
```bash
curl http://localhost:8888/todos
```

### Get Specific Todo
```bash
curl http://localhost:8888/todos/{id}
```

### Update a Todo
```bash
curl -X PUT http://localhost:8888/todos/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Master DDD",
    "description": "Become proficient in Domain-Driven Design",
    "status": "done"
  }'
```

### Mark Todo as Done
```bash
curl -X POST http://localhost:8888/todos/{id}/done
```

### Delete a Todo
```bash
curl -X DELETE http://localhost:8888/todos/{id}
```

## Step 4: Stop Everything

```bash
docker-compose down
```

## Useful Commands

### View logs
```bash
docker-compose logs -f mysql
```

### Connect to MySQL directly
```bash
docker exec -it chi-recap-mysql mysql -u root -p
# Password: rootpass

# Inside MySQL:
USE chi_recap;
SELECT * FROM todos;
```

### Rebuild the application
```bash
go build -o chi-recap
./chi-recap
```

## Project Layout

```
├── api/                    # HTTP handlers & routing
├── internal/domain/        # Business logic & entities  
├── internal/application/   # Use cases & services
├── internal/infrastructure # Database & external deps
└── main.go                 # Entry point
```

## Architecture Pattern

This project uses **Domain-Driven Design** to separate concerns:
- **Domain**: Pure business logic (todo.go)
- **Application**: Use cases (todo_service.go)  
- **Infrastructure**: Database setup (todo_repository.go)
- **API**: HTTP handling (todo_handler.go)

This makes the code:
- ✅ Testable - business logic has no dependencies
- ✅ Maintainable - clear structure and separation
- ✅ Flexible - easy to swap implementations
- ✅ Scalable - organized for growth
