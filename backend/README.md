# Backend API (Go Fiber v3)

A high-performance backend API built with Go Fiber v3 using a **Modular Feature-First** architecture.

## Folder Structure

```bash
backend/
├── cmd/
│   └── main.go                  # Main entry point
├── internal/
│   ├── config/                  # App configuration & env loader
│   ├── database/                # DB Connection & Auto-migrations
│   ├── middleware/              # Shared middleware (Auth, Logger, CORS)
│   ├── modules/                 # Feature-based modules (Domains)
│   │   ├── auth/                # Authentication Feature
│   │   │   ├── handler.go       # Controller layer
│   │   │   ├── service.go       # Business Logic layer
│   │   │   ├── repository.go    # Data Access (Raw SQL)
│   │   │   ├── routes.go        # Module-specific routes
│   │   │   └── types.go         # DTOs & Domain entities
│   │   └── feature2/            # Example Feature Module
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── repository.go
│   │       ├── routes.go
│   │       └── types.go
│   └── router/                  # Central router orchestration
├── pkg/                         # Shared utilities
│   ├── response/                # JSON response helper
│   └── validator/               # Input validation helper
├── .env                         # Local configuration
└── Dockerfile                   # Multi-stage build definition
```

## Architectural Workflow

This project follows a **Feature-First / Vertical Slice Architecture**:
Each feature is self-contained within its own folder under `internal/modules/`.

Data flow:
`HTTP Request` → `feature/handler` → `feature/service` → `feature/repository` → `PostgreSQL`

- **model/types**: Defines domain entities and feature-specific DTOs.
- **handler**: Manages input (parsing JSON/params) and output (status codes, standard responses).
- **service**: Contains core business logic (validation, computation, repository coordination).
- **repository**: Pure database operations using Raw SQL with `pgx`.

## Routing & Integration

Every feature has a `routes.go` file to define its internal endpoints. These modules are then registered in `internal/router/router.go`:

```go
// internal/modules/feature2/routes.go
func RegisterRoutes(router fiber.Router, h *Handler) {
    f2 := router.Group("/feature2")
    f2.Post("/", h.Create) // Full Path: /api/feature2/
}

// internal/router/router.go
func Setup(app *fiber.App, authH *auth.Handler, f2H *feature2.Handler) {
    api := app.Group("/api")
    auth.RegisterRoutes(api, authH)
    feature2.RegisterRoutes(api, f2H)
}
```

## Local Setup

1. Copy `.env.example` to `.env` and configure your credentials.
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the development server:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

- **Auth**
  - `POST /api/auth/register` - Create a new user
  - `POST /api/auth/login` - Authenticate and get JWT
- **Feature2**
  - `POST /api/feature2/` - Example boilerplate endpoint
- **Health**
  - `GET /api/health` - Check system status
