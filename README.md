# NextGo Template

A modern, production-ready full-stack boilerplate using **Next.js** for the frontend and **Go (Fiber v3)** for the backend. Built with a focus on modularity, scalability, and performance.

## Tech Stack

### Frontend
- **Framework**: Next.js (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: React 19 Server Actions & Hooks

### Backend
- **Framework**: Go Fiber v3
- **Database**: PostgreSQL (Raw SQL with pgx/v5)
- **Auth**: JWT (JSON Web Tokens)
- **Architecture**: Modular Feature-First / Domain-Driven

### Infrastructure
- **Containerization**: Docker & Docker Compose

## 📂 Project Structure

```bash
├── frontend/                 # Next.js Application
│   ├── src/
│   │   ├── app/              # Routes & Layouts
│   │   ├── components/       # Shared UI components
│   │   ├── features/         # Feature-based modules (Auth, etc.)
│   │   ├── lib/              # Utilities & API Client
│   │   └── types/            # Global TypeScript types
│   └── package.json
│
├── backend/                  # Go Fiber API (Modular Feature-First)
│   ├── cmd/main.go           # Entry point
│   ├── internal/
│   │   ├── config/           # App configuration
│   │   ├── database/         # Connection & Migrations
│   │   ├── middleware/       # Shared middleware
│   │   ├── modules/          # Feature Modules (Auth, feature2)
│   │   │   ├── auth/         # Auth Feature
│   │   │   └── feature2/     # Example Feature
│   │   └── router/           # Global router wiring
│   ├── pkg/
│   │   └── response/         # Generic API response helpers
│   └── go.mod
│
└── docker-compose.yml        # Orchestration (App + Postgres)
```

## 🛠️ Getting Started

### Prerequisites
- Node.js (v20+)
- Go (v1.23+)
- Docker & Docker Compose

### 1. Database Setup
Run the PostgreSQL service using Docker:
```bash
docker-compose up -d postgres
```

### 2. Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Copy `.env.example` to `.env` and adjust the values.
3. Install dependencies and run:
   ```bash
   go mod tidy
   go run cmd/main.go
   ```

### 3. Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ```

## 🔐 Authentication Flow
1. **Frontend**: Calls Server Actions (`loginAction` / `registerAction`).
2. **Backend**: Validates credentials and returns a JWT token.
3. **Storage**: The token is stored in **Secure HTTP-Only Cookies** and `localStorage` via a universal API client.
4. **Protection**: `middleware.ts` (Next.js) and `AuthRequired` (Go Middleware) protect restricted routes.

## 📄 License
MIT
