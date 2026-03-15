# Frontend — Next.js 16

A modular architecture using **App Router**, **feature-based modules**, and a **centralized design system**.

## Directory Structure

```
src/
├── app/                          # Next.js App Router (routing only)
│   ├── (auth)/                   # Route group — URL without /auth prefix
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
│   ├── ui/                       # Primitives: Button, Input, Modal
│   └── layout/                   # Structural: Navbar, Sidebar, Footer
│
├── features/                     # Feature modules (domain logic)
│   └── auth/
│       ├── index.ts              # Barrel export
│       ├── actions.ts            # Server actions
│       ├── types.ts              # TypeScript types
│       └── hooks/
│           └── use-auth.ts       # Feature-specific hooks
│
├── lib/                          # Core utilities & configuration
│   ├── api.ts                    # Typed API client for Go backend
│   └── utils.ts                  # Helper functions (cn, formatDate)
│
├── styles/                       # Design system & global CSS
│   ├── globals.css               # Entry point — imports all CSS
│   ├── variables.css             # Design tokens (colors, spacing, etc.)
│   └── components.css            # Base component classes (btn, card, etc.)
│
└── types/                        # Global shared TypeScript types
    └── api.d.ts                  # ApiResponse, PaginatedResponse
```

## Layer Explanations

### `app/` — Routing Only

Contains only `page.tsx`, `layout.tsx`, and `loading.tsx`. **No business logic** should be placed here.

- Use **route groups** `(auth)`, `(dashboard)` to group layouts without adding URL segments.
- Example: `app/(auth)/login/page.tsx` → URL is `/login`, not `/auth/login`.

### `components/` — Shared UI

Components that are **used in multiple places** and are **not tied** to a specific feature.

| Sub-folder | Content | Example |
|------------|---------|---------|
| `ui/` | Pure primitives, no business logic | `Button`, `Input`, `Modal` |
| `layout/` | Structural page components | `Navbar`, `Sidebar`, `Footer` |

### `features/` — Feature Modules

The core of project scalability. Each feature has its **own folder** containing all related logic:

```
features/products/          # New feature example
├── index.ts                # Barrel export
├── actions.ts              # Server actions
├── types.ts                # Types
├── hooks/
│   └── use-products.ts     # Feature hooks
└── components/
    └── product-card.tsx    # Feature-specific components
```

**Advantage**: Deleting 1 feature = deleting 1 folder. It doesn't break other features.

### `lib/` — Core Utilities

| File | Function |
|------|----------|
| `api.ts` | Typed fetch wrapper for the Go backend, with base URL from env |
| `utils.ts` | Helpers: `cn()` (class merge), `formatDate()`, etc. |

### `styles/` — Design System

CSS is organized into 3 files imported sequentially:

| File | Content |
|------|---------|
| `variables.css` | Design tokens: colors, typography, spacing, shadows, z-index |
| `globals.css` | Entry point: imports tokens → Tailwind → components → base styles |
| `components.css` | Reusable CSS classes: `.btn`, `.card`, `.input`, `.badge` |

Imported from `layout.tsx`:
```tsx
import '@/styles/globals.css';
```

### `types/` — Global Types

TypeScript types used across features (not specific to a single feature).

```ts
import type { ApiResponse, PaginatedResponse } from '@/types/api';
```

## Conventions

### Path Alias

Uses `@/` mapped to `./src/*` in `tsconfig.json`:

```ts
import { Button } from '@/components/ui/button';
import { useAuth } from '@/features/auth';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';
```

### Naming

| Type | Format | Example |
|------|--------|---------|
| Files & folders | `kebab-case` | `user-card.tsx`, `use-auth.ts` |
| React Components | `PascalCase` | `UserCard`, `ProductList` |
| Hooks & utils | `camelCase` | `useAuth`, `formatDate` |
| Types | `PascalCase` | `AuthUser`, `ApiResponse` |

### Server vs Client Components

By default, all components are **Server Components**. Add `"use client"` at the first line if interactivity is needed (`useState`, `useEffect`, event handlers).

```tsx
"use client";  // ← Required for interactive components

import { useState } from 'react';
```

### Middleware

`middleware.ts` is located in the `frontend/` root (parallel to `src/`), acting as a global auth guard.

```ts
// frontend/middleware.ts
export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
```
