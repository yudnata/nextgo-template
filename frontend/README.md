# Frontend — Next.js 16

Arsitektur modular dengan **App Router**, **feature-based modules**, dan **centralized design system**.

## Struktur Direktori

```
src/
├── app/                          # Next.js App Router (routing only)
│   ├── (auth)/                   # Route group — URL tanpa /auth prefix
│   │   ├── layout.tsx            # Centered auth layout
│   │   ├── login/page.tsx        # → /login
│   │   └── register/page.tsx     # → /register
│   ├── dashboard/
│   │   ├── layout.tsx            # → /dashboard layout
│   │   └── page.tsx              # → /dashboard
│   ├── layout.tsx                # Root layout
│   └── page.tsx                  # Home (/)
│
├── components/                   # Shared reusable components
│   ├── ui/                       # Primitif: Button, Input, Modal
│   └── layout/                   # Struktural: Navbar, Sidebar, Footer
│
├── features/                     # Feature modules (domain logic)
│   └── auth/
│       ├── index.ts              # Barrel export
│       ├── actions.ts            # Server actions
│       ├── types.ts              # TypeScript types
│       └── hooks/
│           └── use-auth.ts       # Feature-specific hooks
│
├── lib/                          # Core utilities & konfigurasi
│   ├── api.ts                    # Typed API client ke Go backend
│   └── utils.ts                  # Helper functions (cn, formatDate)
│
├── styles/                       # Design system & global CSS
│   ├── globals.css               # Entry point — import semua CSS
│   ├── variables.css             # Design tokens (colors, spacing, etc.)
│   └── components.css            # Base component classes (btn, card, etc.)
│
└── types/                        # Global shared TypeScript types
    └── api.d.ts                  # ApiResponse, PaginatedResponse
```

## Penjelasan Tiap Layer

### `app/` — Routing Only

Hanya berisi `page.tsx`, `layout.tsx`, dan `loading.tsx`. **Tidak boleh** ada business logic di sini.

- Gunakan **route groups** `(auth)`, `(dashboard)` untuk mengelompokkan layout tanpa menambah segment URL.
- Contoh: `app/(auth)/login/page.tsx` → URL-nya `/login`, bukan `/auth/login`.

### `components/` — Shared UI

Komponen yang **dipakai di banyak tempat** dan **tidak terikat** ke fitur tertentu.

| Sub-folder | Isi | Contoh |
|------------|-----|--------|
| `ui/` | Primitif murni, tanpa business logic | `Button`, `Input`, `Modal` |
| `layout/` | Komponen struktural halaman | `Navbar`, `Sidebar`, `Footer` |

### `features/` — Feature Modules

Inti skalabilitas proyek. Setiap fitur punya **folder sendiri** yang berisi semua logic terkait:

```
features/products/          # Contoh fitur baru
├── index.ts                # Barrel export
├── actions.ts              # Server actions
├── types.ts                # Types
├── hooks/
│   └── use-products.ts     # Feature hooks
└── components/
    └── product-card.tsx    # Feature-specific components
```

**Keuntungan**: Hapus 1 fitur = hapus 1 folder. Tidak merusak fitur lain.

### `lib/` — Core Utilities

| File | Fungsi |
|------|--------|
| `api.ts` | Typed fetch wrapper ke Go backend, dengan base URL dari env |
| `utils.ts` | Helper: `cn()` (class merge), `formatDate()`, dll. |

### `styles/` — Design System

CSS diorganisir dalam 3 file yang di-import secara berurutan:

| File | Isi |
|------|-----|
| `variables.css` | Design tokens: colors, typography, spacing, shadows, z-index |
| `globals.css` | Entry point: import tokens → Tailwind → components → base styles |
| `components.css` | Reusable CSS classes: `.btn`, `.card`, `.input`, `.badge` |

Import dari `layout.tsx`:
```tsx
import '@/styles/globals.css';
```

### `types/` — Global Types

Tipe TypeScript yang dipakai lintas feature (bukan spesifik ke 1 fitur).

```ts
import type { ApiResponse, PaginatedResponse } from '@/types/api';
```

## Konvensi

### Path Alias

Menggunakan `@/` yang di-mapping ke `./src/*` di `tsconfig.json`:

```ts
import { Button } from '@/components/ui/button';
import { useAuth } from '@/features/auth';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';
```

### Penamaan

| Jenis | Format | Contoh |
|-------|--------|--------|
| File & folder | `kebab-case` | `user-card.tsx`, `use-auth.ts` |
| Komponen React | `PascalCase` | `UserCard`, `ProductList` |
| Hooks & utils | `camelCase` | `useAuth`, `formatDate` |
| Types | `PascalCase` | `AuthUser`, `ApiResponse` |

### Server vs Client Components

Secara default, semua komponen adalah **Server Components**. Tambahkan `"use client"` di baris pertama jika butuh interaktivitas (`useState`, `useEffect`, event handler).

```tsx
"use client";  // ← wajib untuk komponen interaktif

import { useState } from 'react';
```

### Middleware

`middleware.ts` terletak di root `frontend/` (sejajar `src/`), berfungsi sebagai auth guard global.

```ts
// frontend/middleware.ts
export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
```
