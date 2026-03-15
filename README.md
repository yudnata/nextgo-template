# NextGo Stack Template

Boilerplate modern untuk membangun web app dengan **Next.js 16** (Frontend) dan **Go Fiber** (Backend). Dioptimasi untuk performa, skalabilitas, dan developer experience.

## Tech Stack

| Layer | Teknologi | Versi |
|-------|-----------|-------|
| Frontend | Next.js (App Router) | 16.1.6 |
| UI Framework | React | 19.2.3 |
| Styling | Tailwind CSS | 4 |
| Language | TypeScript | 5 |
| Backend | Go + Fiber v3 | 1.25.0 |
| Database | PostgreSQL | 16 |
| Auth | JWT + Bcrypt | - |
| SQL Driver | Raw SQL (pgx/v5) | - |
| Container | Docker Compose | 3.8 |

## Struktur Proyek

```
nextgo-template/
├── frontend/                 # Next.js Application
│   ├── src/
│   │   ├── app/              # App Router (routing & pages)
│   │   ├── components/       # Shared reusable components
│   │   ├── features/         # Feature modules (domain logic)
│   │   ├── lib/              # Core utilities & API client
│   │   ├── styles/           # Design system & global CSS
│   │   └── types/            # Global shared TypeScript types
│   ├── middleware.ts          # Auth guard & request middleware
│   ├── next.config.ts
│   └── package.json
│
├── backend/                  # Go Fiber API (Modular Feature-First)
│   ├── cmd/main.go           # Entry point
│   ├── internal/
│   │   ├── config/           # App configuration
│   │   ├── database/         # DB Connection & Migrations
│   │   ├── middleware/       # Shared middleware
│   │   ├── modules/          # Feature Modules
│   │   │   ├── auth/         # Auth Feature
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── routes.go
│   │   │   │   └── types.go
│   │   │   └── feature2/     # Feature 2 (Boilerplate)
│   │   │       ├── handler.go
│   │   │       ├── service.go
│   │   │       ├── repository.go
│   │   │       ├── routes.go
│   │   │       └── types.go
│   │   └── router/           # Global router wiring
│   ├── pkg/
│   │   ├── response/         # JSON response helper
│   │   └── validator/        # Input validation
│   └── go.mod
│
└── docker-compose.yml        # Orchestration (App + Postgres)
```

> Lihat [frontend/README.md](frontend/README.md) untuk detail arsitektur frontend.

## Kenapa Next.js + Go?

**Next.js** — SSR/SSG untuk SEO dan performa, file-based routing yang clean, dan React Server Components untuk efisiensi rendering.

**Go + Fiber v3** — Compiled langsung ke machine code, goroutines untuk concurrency murah, dan Fiber v3 lebih modern & performant untuk handling API.

**Raw SQL (pgx)** — Kontrol penuh atas query, performa maksimal, dan tanpa overhead dari ORM yang kompleks.

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Node.js ≥ 18 (rekomendasi 20+)
- Go ≥ 1.25 (untuk fitur terbaru)

### 1. Clone & Konfigurasi Environment

```bash
# Backend
cd backend
cp .env.example .env

# Frontend
cd ../frontend
cp .env.example .env.local
```

### 2. Jalankan dengan Docker (Paling Mudah)

```bash
docker-compose up --build
```

| Service | URL |
|---------|-----|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8080/api/health |
| Postgres | localhost:5432 |

### 3. Jalankan Secara Lokal (Development)

**Backend:**
1. Pastikan PostgreSQL berjalan (bisa pakai docker: `docker run -d --name pg -e POSTGRES_PASSWORD=nextgo_secret -p 5432:5432 postgres:16-alpine`)
2. `cd backend`
3. `go mod tidy`
4. `go run cmd/api/main.go`
# → Server di port 8080, otomatis migrate table `users`

**Frontend:**
1. `cd frontend`
2. `npm install`
3. `npm run dev`
# → App di port 3000

## Environment Variables

### Frontend (`frontend/.env.local`)
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
NODE_ENV=development
```

### Backend (`backend/.env`)
```env
PORT=8080
DATABASE_URL=postgres://nextgo:nextgo_secret@localhost:5432/nextgo_db?sslmode=disable
JWT_SECRET=super_secret_jwt_key
MODE=development
```

## Scripts

| Command | Deskripsi |
|---------|-----------|
| `npm run dev` | Jalankan dev server |
| `npm run build` | Build production |
| `npm run start` | Jalankan production server |
| `npm run lint` | Jalankan ESLint |
