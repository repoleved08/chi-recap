# Huma API Documentation Setup

## Overview

This project now includes **Huma**, a REST API framework for Go that automatically generates interactive API documentation.

### What is Huma?

Huma is a modern Go web framework that:
- ✅ Automatically generates OpenAPI 3.0 documentation
- ✅ Provides interactive Swagger UI for testing endpoints
- ✅ Type-safe request/response handling
- ✅ Built-in validation and error handling
- ✅ Zero boilerplate code

### Documentation Access

Once the server is running, access the documentation at:

| Resource | URL |
|----------|-----|
| **Interactive Docs** | `http://localhost:8888/docs` |
| **Swagger UI** | `http://localhost:8888/docs/ui` |
| **OpenAPI Schema** | `http://localhost:8888/openapi.json` |
| **ReDoc** | `http://localhost:8888/docs/redoc` |

## Quick Start

### 1. Start the Application

```bash
# Start database
docker-compose up -d

# Start the server
go run main.go
```

### 2. Open API Documentation

Visit: **http://localhost:8888/docs**

You'll see the interactive Swagger UI with:
- ✅ All endpoints listed
- ✅ Request/response examples
- ✅ Parameter descriptions
- ✅ "Try it out" button to test endpoints

### 3. Test an Endpoint

1. Navigate to "POST /todos"
2. Click "Try it out"
3. Enter sample data:
   ```json
   {
     "title": "Learn Huma",
     "description": "Master the Huma REST framework"
   }
   ```
4. Click "Execute"
5. See the response

## API Endpoints with Huma Documentation

### Health Check

```
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

### Create Todo

```
POST /todos
```

**Request Body:**
```json
{
  "title": "Learn Go",
  "description": "Master Go programming language"
}
```

**Response (201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Learn Go",
  "description": "Master Go programming language",
  "status": "pending",
  "created_at": 1716193200,
  "updated_at": 1716193200
}
```

### List All Todos

```
GET /todos
```

**Response:**
```json
{
  "todos": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Learn Go",
      "description": "Master Go programming language",
      "status": "pending",
      "created_at": 1716193200,
      "updated_at": 1716193200
    }
  ]
}
```

### Get Single Todo

```
GET /todos/{id}
```

**Path Parameters:**
- `id` - Todo ID (UUID)

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Learn Go",
  "description": "Master Go programming language",
  "status": "pending",
  "created_at": 1716193200,
  "updated_at": 1716193200
}
```

### Update Todo

```
PUT /todos/{id}
```

**Path Parameters:**
- `id` - Todo ID (UUID)

**Request Body:**
```json
{
  "title": "Master Go",
  "description": "Become proficient in Go programming",
  "status": "done"
}
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Master Go",
  "description": "Become proficient in Go programming",
  "status": "done",
  "created_at": 1716193200,
  "updated_at": 1716193210
}
```

### Delete Todo

```
DELETE /todos/{id}
```

**Path Parameters:**
- `id` - Todo ID (UUID)

**Response:** 204 No Content

### Mark Todo as Done

```
POST /todos/{id}/done
```

**Path Parameters:**
- `id` - Todo ID (UUID)

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Learn Go",
  "description": "Master Go programming language",
  "status": "done",
  "created_at": 1716193200,
  "updated_at": 1716193210
}
```

## Using the Interactive Documentation

### Testing with Swagger UI

1. **Open** http://localhost:8888/docs
2. **Find endpoint** - Endpoints are grouped by resource
3. **Click endpoint** - Expands to show details
4. **Try it out** - Click the button to test
5. **Enter values** - Fill in required parameters
6. **Execute** - Send the request
7. **View response** - See status, headers, and body

### Generating OpenAPI Schema

The OpenAPI schema is automatically generated and available at:

```bash
curl http://localhost:8888/openapi.json | jq .
```

This JSON can be used with tools like:
- Swagger Editor
- Postman
- API documentation generators

## How Huma Documentation Works

### Request/Response Types

Huma automatically documents request and response types based on struct tags:

```go
type TodoCreateRequest struct {
    Title       string `json:"title" doc:"Todo title" maxLength:"255"`
    Description string `json:"description" doc:"Todo description"`
}

type TodoResponse struct {
    ID        string `json:"id" doc:"Unique identifier (UUID)"`
    Title     string `json:"title" doc:"Todo title"`
    Status    string `json:"status" doc:"Status" enum:"pending,done"`
    CreatedAt int64  `json:"created_at" doc:"Creation timestamp (Unix)"`
}
```

### Struct Tag Documentation

| Tag | Purpose | Example |
|-----|---------|---------|
| `json` | JSON field name | `json:"title"` |
| `doc` | Field description | `doc:"Todo title"` |
| `example` | Example value | `example:"Learn Go"` |
| `maxLength` | Max string length | `maxLength:"255"` |
| `enum` | Allowed values | `enum:"pending,done"` |
| `path` | Path parameter | `path:"id"` |

### Endpoint Handler Pattern

```go
huma.Post(api, "/todos", func(ctx context.Context, input *struct {
    Body TodoCreateRequest
}) (*TodoResponse, error) {
    // Implementation
    return response, nil
})
```

**Components:**
- `huma.Post` - HTTP method and router
- `/todos` - URL path
- `input` - Struct with request data
- `TodoResponse` - Response type
- `error` - Error handling

## API Metadata

The API metadata is configured in `api/huma.go`:

```go
config := huma.DefaultConfig("Todo API", "1.0.0")
config.Info.Description = "A complete REST API for managing todos..."
config.Info.Contact = &huma.Contact{
    Name:  "Chi-Recap Project",
    Email: "support@chi-recap.local",
    URL:   "https://github.com/repoleved08/chi-recap",
}
config.Info.License = &huma.License{
    Name: "MIT",
    URL:  "https://opensource.org/licenses/MIT",
}
```

This metadata appears in the documentation:
- API title and version
- Description
- Contact information
- License

## Common Tasks

### Adding a New Endpoint with Documentation

1. **Define request type:**
   ```go
   type CreateUserRequest struct {
       Name  string `json:"name" doc:"User name" maxLength:"255"`
       Email string `json:"email" doc:"User email" format:"email"`
   }
   ```

2. **Define response type:**
   ```go
   type UserResponse struct {
       ID    string `json:"id" doc:"User ID (UUID)"`
       Name  string `json:"name" doc:"User name"`
       Email string `json:"email" doc:"User email"`
   }
   ```

3. **Register handler:**
   ```go
   huma.Post(api, "/users", func(ctx context.Context, input *struct {
       Body CreateUserRequest
   }) (*UserResponse, error) {
       // Implementation
       return response, nil
   })
   ```

4. **Documentation auto-generated!**

### Viewing Generated OpenAPI

```bash
# Pretty print OpenAPI schema
curl http://localhost:8888/openapi.json | jq '.'

# Export to file
curl http://localhost:8888/openapi.json > openapi.json

# Validate with online tools
# Visit: https://editor.swagger.io
# Paste content of openapi.json
```

## Troubleshooting

### Documentation Not Showing

```bash
# Check if server is running
curl http://localhost:8888/health

# Check Huma docs endpoint
curl http://localhost:8888/docs
```

### OpenAPI Schema Not Generating

- Ensure handler functions follow the pattern: `func(context.Context, *Input) (*Output, error)`
- All request/response types must have proper struct tags
- Check for build errors: `go build`

### Cannot Access Documentation UI

```bash
# Verify port is correct
lsof -i :8888

# Check firewall
sudo ufw allow 8888

# Try different browser or clear cache
```

## Advanced Features

### Custom Error Responses

```go
return nil, &huma.ErrorModel{
    Status: http.StatusBadRequest,
    Title:  "Validation Error",
    Detail: "Invalid todo data",
}
```

### Status Codes

```go
// Framework handles these automatically
// 200 - OK (GET, POST, PUT)
// 201 - Created (POST creating resource)
// 204 - No Content (DELETE, PUT with no response)
// 400 - Bad Request (validation errors)
// 404 - Not Found (resource not found)
// 500 - Internal Server Error (exceptions)
```

### Query Parameters

```go
huma.Get(api, "/todos", func(ctx context.Context, input *struct {
    Status string `query:"status" doc:"Filter by status" enum:"pending,done"`
    Limit  int    `query:"limit" doc:"Max results" example:"10"`
}) (*TodoListResponse, error) {
    // Implementation
})
```

### Path Parameters

```go
huma.Get(api, "/todos/{id}", func(ctx context.Context, input *struct {
    ID string `path:"id" doc:"Todo ID"`
}) (*TodoResponse, error) {
    // Implementation
})
```

## Documentation Features

### Swagger UI Features
- ✅ Interactive endpoint testing
- ✅ Request/response examples
- ✅ Schema validation
- ✅ Authentication support
- ✅ Model expansion
- ✅ Try it out functionality

### ReDoc Features
- ✅ Clean, easy-to-read documentation
- ✅ Night mode
- ✅ Search functionality
- ✅ Good for long documentation

## Integration with Other Tools

### Postman

1. Open Postman
2. Click "Import"
3. Select "Link"
4. Enter: `http://localhost:8888/openapi.json`
5. Collections imported automatically

### Insomnia

1. Open Insomnia
2. Create new API client
3. Paste OpenAPI URL: `http://localhost:8888/openapi.json`
4. All endpoints available

### VS Code REST Client

Create `.http` file:
```
### Create Todo
POST http://localhost:8888/todos
Content-Type: application/json

{
  "title": "Learn Huma",
  "description": "Master REST API documentation"
}

### Get All Todos
GET http://localhost:8888/todos
```

## Best Practices

✅ **DO:**
- Document all fields with `doc` tags
- Provide meaningful `example` values
- Use `enum` for constrained values
- Add `description` to struct types (comments)
- Keep descriptions concise
- Use proper HTTP status codes
- Validate input data

❌ **DON'T:**
- Leave fields undocumented
- Use vague field names
- Skip example values
- Mix multiple concerns in one endpoint
- Ignore error cases

## References

- **Huma Documentation**: https://huma.rocks/
- **OpenAPI 3.0 Spec**: https://spec.openapis.org/oas/v3.0.3
- **Swagger UI**: https://swagger.io/tools/swagger-ui/

## Next Steps

1. ✅ Start the server: `go run main.go`
2. ✅ Open docs: http://localhost:8888/docs
3. ✅ Test endpoints in Swagger UI
4. ✅ Export schema for CI/CD pipelines
5. ✅ Integrate with API management tools

Your API is now fully documented and interactive! 🚀
