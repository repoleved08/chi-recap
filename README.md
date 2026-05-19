# Chi-Recap: Todo API with DDD Architecture

A clean, production-ready Todo API built with Go using Domain-Driven Design (DDD) patterns, Chi router, GORM, MySQL, and **Huma for automatic API documentation**.

### 🎯 Key Features

- ✅ **Domain-Driven Design** - Clean separation of concerns
- ✅ **Huma Documentation** - Auto-generated OpenAPI/Swagger UI
- ✅ **GORM ORM** - Type-safe database access
- ✅ **MySQL** - Production-ready database
- ✅ **Chi Router** - Fast HTTP router
- ✅ **Database Migrations** - Automated schema management
- ✅ **Docker Support** - Easy local development

### 📚 Interactive API Documentation

Once running, access the API documentation at:

| Resource | URL |
|----------|-----|
| **Swagger UI** | `http://localhost:8888/docs` |
| **OpenAPI Schema** | `http://localhost:8888/openapi.json` |
| **ReDoc** | `http://localhost:8888/docs/redoc` |

See [HUMA_DOCUMENTATION.md](./HUMA_DOCUMENTATION.md) for details.

## Project Structure

Following Domain-Driven Design principles:

```
chi-recap/
├── api/                          # API layer (HTTP handlers & routing)
│   ├── handlers/
│   │   └── todo_handler.go      # HTTP request handlers for todos
│   └── router.go                # Route definitions
├── internal/
│   ├── domain/                  # Domain layer (business logic & entities)
│   │   ├── shared/
│   │   │   └── id.go            # Value object for IDs
│   │   └── todo/
│   │       ├── todo.go          # Todo aggregate root
│   │       └── repository.go    # Repository interface
│   ├── application/             # Application layer (use cases)
│   │   └── services/
│   │       └── todo_service.go  # Todo business logic & services
│   └── infrastructure/          # Infrastructure layer (external dependencies)
│       ├── config/
│       │   └── database.go      # Database configuration
│       └── persistence/
│           └── todo_repository.go  # GORM repository implementation
├── main.go                      # Application entry point
├── docker-compose.yml           # MySQL database setup
├── go.mod                       # Go dependencies
└── .env                        # Environment configuration
```

## Architecture Layers

### Domain Layer
- **Entities**: `Todo` aggregate root with business logic
- **Value Objects**: `ID` for unique identifiers
- **Repository Interface**: Abstraction for data persistence

### Application Layer
- **Services**: `TodoService` with use cases
- **DTOs**: Request/Response models for API

### Infrastructure Layer
- **Database Config**: GORM MySQL connection setup
- **Repository Implementation**: Concrete GORM implementation

### API Layer
- **Handlers**: HTTP endpoints implementation
- **Router**: Route definitions using Chi

## Prerequisites

- Go 1.26.2+
- Docker & Docker Compose
- MySQL 8.0 (via Docker)

## Getting Started

### 1. Start MySQL Database

```bash
docker-compose up -d
```

This will start MySQL on `localhost:3306` with:
- User: `root`
- Password: `rootpass`
- Database: `chi_recap`

### 2. Install Dependencies

```bash
go mod download
```

### 3. Run the Application

```bash
go run main.go
```

The API will be available at `http://localhost:8888`

### 4. Build for Production

```bash
go build -o chi-recap
./chi-recap
```

## API Endpoints

### Health Check
- `GET /health` - Check API health status

### Todos

- `POST /todos` - Create a new todo
  ```json
  {
    "title": "Buy groceries",
    "description": "Milk, eggs, bread"
  }
  ```

- `GET /todos` - List all todos

- `GET /todos/{id}` - Get a specific todo

- `PUT /todos/{id}` - Update a todo
  ```json
  {
    "title": "Buy groceries",
    "description": "Milk, eggs, bread, butter",
    "status": "done"
  }
  ```

- `DELETE /todos/{id}` - Delete a todo

- `POST /todos/{id}/done` - Mark todo as done

## Environment Variables

Configure via `.env` file:

```
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=rootpass
DB_NAME=chi_recap
PORT=8888
```

## Development

### Code Structure Benefits

- **Separation of Concerns**: Each layer has a clear responsibility
- **Testability**: Services can be tested without HTTP concerns
- **Flexibility**: Easy to swap implementations (e.g., different databases)
- **Scalability**: Domain logic is independent of frameworks
- **Maintainability**: Clear structure makes code easier to understand

### Key DDD Concepts

- **Aggregate**: `Todo` acts as an aggregate root
- **Repository Pattern**: Data access abstraction
- **Value Objects**: `ID` encapsulates identifier logic
- **Domain Services**: Business logic in `TodoService`
- **Dependency Injection**: Services receive dependencies

## Common Tasks

### Stop Database
```bash
docker-compose down
```

### View Database Logs
```bash
docker-compose logs mysql
```

### Access MySQL CLI
```bash
docker exec -it chi-recap-mysql mysql -u root -p
```

## Future Enhancements

- [ ] Add authentication & authorization
- [ ] Implement pagination for list endpoints
- [ ] Add request validation
- [ ] Implement error handling middleware
- [ ] Add logging
- [ ] Create integration tests
- [ ] Add database migrations with flyway/migrate
- [ ] Implement CQRS pattern for complex queries
- [ ] Add WebSocket support for real-time updates
- [ ] Containerize the application

## License

MIT
