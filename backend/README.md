# Backend API (Go Fiber v3)

Aplikasi backend menggunakan Go Fiber v3 dengan arsitektur **Modular Feature-First**.

## Struktur Folder

```bash
backend/
├── cmd/
│   └── main.go                  # Entry point utama
├── internal/
│   ├── config/                  # Konfigurasi aplikasi & env
│   ├── database/                # Koneksi database & migrasi
│   ├── middleware/              # Middleware (Auth, Logger, CORS)
│   ├── modules/                 # Modul fitur (Domain)
│   │   ├── auth/                # Fitur Autentikasi
│   │   │   ├── handler.go       # Controller
│   │   │   ├── service.go       # Business Logic
│   │   │   ├── repository.go    # DB Access (Raw SQL)
│   │   │   ├── routes.go        # Routing internal modul
│   │   │   └── types.go         # Struct Req/Res & Entity
│   │   └── feature2/            # Contoh Fitur Kedua
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── repository.go
│   │       ├── routes.go
│   │       └── types.go
│   └── router/                  # Central router wiring
├── pkg/                         # Shared libraries
│   ├── response/                # Helper JSON response
│   └── validator/               # Helper validasi input
├── .env                         # Konfigurasi env (Lokal)
└── Dockerfile                   # Definisi build docker
```

## Arsitektur & Sinkronisasi Route

Setiap fitur memiliki file `routes.go` sendiri untuk mendefinisikan endpoint internalnya. Semua modul fitur kemudian didaftarkan di `internal/router/router.go`:

```go
// internal/modules/feature2/routes.go
func RegisterRoutes(router fiber.Router, h *Handler) {
    f2 := router.Group("/feature2")
    f2.Post("/", h.Create) // Endpoint: /api/feature2/
}

// internal/router/router.go
func Setup(app *fiber.App, authH *auth.Handler, f2H *feature2.Handler) {
    api := app.Group("/api")
    auth.RegisterRoutes(api, authH)
    feature2.RegisterRoutes(api, f2H)
}
```

## Setup Lokal

1. Salin file `.env.example` menjadi `.env` dan sesuaikan nilainya.
2. Jalankan dependensi:
   ```bash
   go mod tidy
   ```
3. Jalankan server:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

- **Auth**
  - `POST /api/auth/register` - Daftar user baru
  - `POST /api/auth/login` - Login & ambil token
- **Feature2**
  - `POST /api/feature2/` - Contoh hit endpoint baru
- **Lainnya**
  - `GET /api/health` - Cek status server via router (sedang dalam proses pemindahan)
