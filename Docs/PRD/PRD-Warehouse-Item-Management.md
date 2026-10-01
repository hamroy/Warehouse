TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
TECH STACK

  - Frontend: React 18+, Next.js 14+, TypeScript 5+, Tailwind CSS
  - Backend: Node.js, Express, GraphQL (Apollo), tRPC
  - Database: PostgreSQL, Prisma ORM, Redis caching
  - Testing: Playwright (E2E), Vitest (unit), React Testing Library
  - DevOps: Docker, GitHub Actions, AWS/Vercel

  ARCHITECTURE PRINCIPLES

  1. Clean Architecture: Domain → Use Case → Interface → Framework
  2. Component-Driven: Atomic design, single responsibility
  3. Type Safety: Strict TypeScript, zero any, branded types
  4. API Design: GraphQL internal, REST external, versioning mandatory
  5. Testing Pyramid: 70% unit, 20% integration, 10% E2E

  CODE PATTERNS

  // Frontend: Compound Component Pattern
  const Card = {
    Root: ({ children }) => <div className="rounded-lg border">{children}</div>,
    Header: ({ title }) => <h3 className="font-bold">{title}</h3>,
    Body: ({ children }) => <div className="p-4">{children}</div>
  };

  // Backend: Repository + Service Pattern
  class UserRepository {
    async findById(id: UserId): Promise<User> {
      return this.db.user.findUnique({ where: { id } });
    }
  }

  // Types: Branded types for IDs
  type UserId = string & { __brand: 'UserId' };
  const createUserId = (id: string): UserId => id as UserId;

  WORKFLOW

  1. Analyze requirements → Domain model first
  2. Write tests → Red-Green-Refactor
  3. Implement → Feature slices, not layers
  4. Review → Self-review checklist below
  5. Deploy → Docker → CI/CD → Staging → Prod

  CODE REVIEW CHECKLIST

  - [ ] No business logic in components (use hooks/services)
  - [ ] All functions < 20 lines, pure when possible
  - [ ] Error handling: never swallow, always log
  - [ ] Loading states: skeletons, not spinners
  - [ ] Accessibility: aria-labels, keyboard nav
  - [ ] Performance: React.memo for lists, lazy load routes

  RESPONSE STYLE

  - Direct, no fluff, code-first explanations
  - Always provide complete, runnable examples
  - Suggest 2-3 alternatives with trade-offs
  - Flag security risks immediately
  - Ask clarifying questions before coding

  CONSTRAINTS

  - Never use any or @ts-ignore
  - Never expose database IDs to client (use UUIDs)
  - Never trust client input (validate at boundary)
  - Never commit secrets (use env vars + validation)
# PRD — Warehouse Item Management System

**Project:** Warehouse Item Management System  
**Type:** Fullstack Technical Test  
**Version:** MVP 1.0  
**Target:** Technical Assessment

---

## 1. Product Overview

### 1.1 Background

Sistem digunakan untuk mengelola **master barang gudang dan stok barang berdasarkan lokasi penyimpanan**.

Sistem minimal harus mampu:

1. Membuat item/barang.
2. Melihat daftar item.
3. Mencari dan memfilter item.
4. Mengubah item.
5. Soft delete item.
6. Menerima stok ke lokasi tertentu.
7. Melihat stok sebuah item pada seluruh lokasi.

Konsep utama:

```text
Item
  │
  ├── SKU
  ├── Name
  ├── Category
  └── Unit
       │
       ▼
Stock
       │
       ├── Location A → 100
       ├── Location B → 50
       └── Location C → 25
```

Entitas minimal yang diwajibkan adalah **Item, Location, dan Stock**.

---

## 2. Product Goal

Membangun aplikasi fullstack sederhana yang menunjukkan kemampuan kandidat dalam:

- REST API design
- Layered architecture
- Database design
- Business rule implementation
- Error handling
- API documentation
- React application development
- API integration
- Dockerization
- Technical documentation

Yang dinilai bukan hanya aplikasi berhasil berjalan, tetapi juga:

- Pemahaman terhadap kode.
- Kemampuan menjelaskan keputusan desain.
- Trade-off yang diambil.
- Struktur kode.
- Production-readiness.

---

## 3. Scope MVP

### 3.1 Backend

#### Item Management

- Create Item
- List Item
- Search Item
- Filter Category
- Pagination
- Get Item Detail
- Update Item
- Soft Delete Item

#### Stock Management

- Receive stock
- Stock berdasarkan item
- Stock berdasarkan location

#### Technical

- REST API
- PostgreSQL
- Database migration
- Layered architecture
- Middleware
- Standard response envelope
- Centralized error handling
- OpenAPI 3.0
- Docker + Docker Compose

---

### 3.2 Frontend

#### Dashboard

- Item list
- Search
- Category filter
- Pagination
- Create item
- Edit item

#### Form

- SKU
- Name
- Category
- Unit
- Client-side validation
- Loading state
- Success notification
- Error notification

#### UI

- Responsive desktop
- Responsive tablet
- Loading skeleton/spinner
- Toast notification

---

## 4. Out of Scope

Fitur berikut tidak wajib untuk MVP:

- Stock transfer
- Location management API
- Item detail & stock page
- Global Error Boundary
- Redis Cache
- Authentication
- Authorization
- Deployment

`/stock/transfer` dan `/locations` dapat ditambahkan sebagai nilai tambah.

---

# 5. User Role

Technical test tidak mendefinisikan authentication/authorization.

Untuk MVP digunakan satu implicit role:

### Warehouse Operator

Permissions:

```text
View Items
Create Item
Edit Item
Delete Item
Receive Stock
View Stock
```

---

# 6. Domain Model

## 6.1 Item

Master data barang.

| Field | Type | Required | Rule |
|---|---|---:|---|
| id | UUID | Yes | Primary Key |
| sku | string | Yes | Unique |
| name | string | Yes | Required |
| category | string | Yes | Required |
| unit | string | Yes | Required |
| created_at | timestamp | Yes | Auto |
| updated_at | timestamp | Yes | Auto |
| deleted_at | timestamp | No | Soft Delete |

---

## 6.2 Location

Lokasi penyimpanan barang.

| Field | Type | Required |
|---|---|---:|
| id | UUID | Yes |
| code | string | Yes |
| zone | string | Yes |
| type | string | Yes |

Example:

```text
Location
--------------------
id: loc-001
code: A-01-01
zone: A
type: RACK
```

---

## 6.3 Stock

Kuantitas item pada lokasi tertentu.

| Field | Type | Required |
|---|---|---:|
| id | UUID | Yes |
| item_id | UUID | Yes |
| location_id | UUID | Yes |
| qty | integer/decimal | Yes |
| updated_at | timestamp | Yes |

Relationship:

```text
Item 1 ─────── N Stock
                  │
                  N
                  │
                  1
               Location
```

Satu item dapat mempunyai stok di banyak lokasi.

---

# 7. Database ERD

```text
┌──────────────────────┐
│        items         │
├──────────────────────┤
│ id PK                │
│ sku UNIQUE           │
│ name                 │
│ category             │
│ unit                 │
│ created_at           │
│ updated_at           │
│ deleted_at           │
└──────────┬───────────┘
           │
           │ 1
           │
           │ N
┌──────────▼───────────┐
│        stocks        │
├──────────────────────┤
│ id PK                │
│ item_id FK           │
│ location_id FK       │
│ qty                  │
│ updated_at            │
└──────────┬───────────┘
           │ N
           │
           │ 1
┌──────────▼───────────┐
│      locations       │
├──────────────────────┤
│ id PK                │
│ code                 │
│ zone                 │
│ type                 │
└──────────────────────┘
```

### Recommended Constraint

```text
UNIQUE(item_id, location_id)
```

Tujuannya agar satu item pada satu lokasi hanya mempunyai satu record stock.

---

# 8. Business Rules

## BR-001 — SKU Unique

SKU item harus unik.

Jika SKU sudah digunakan:

```http
409 Conflict
```

Response:

```json
{
  "success": false,
  "message": "SKU already exists",
  "errors": [
    {
      "field": "sku",
      "reason": "Duplicate entry"
    }
  ]
}
```

---

## BR-002 — Stock Tidak Boleh Negatif

Validasi wajib dilakukan pada service layer.

Contoh:

```text
Current Stock = 10
Receive       = 5
Result        = 15
```

Quantity tidak boleh menghasilkan stock negatif.

Validasi:

```text
qty > 0
```

---

## BR-003 — Soft Delete

Delete item tidak melakukan physical delete.

Gunakan:

```sql
UPDATE items
SET deleted_at = NOW()
WHERE id = $1;
```

List item hanya menampilkan:

```sql
WHERE deleted_at IS NULL
```

---

## BR-004 — Stock Transaction Logging

Setiap mutasi stock `receive` harus mencatat log transaksi minimal pada application level.

Contoh:

```text
2026-10-01 13:30:22
RECEIVE_STOCK
item=SKU-001
location=A-01
qty=100
```

---

# 9. API Specification

## 9.1 Create Item

```http
POST /api/v1/items
```

Request:

```json
{
  "sku": "SKU-001",
  "name": "Laptop Dell",
  "category": "Electronics",
  "unit": "PCS"
}
```

Response:

```json
{
  "success": true,
  "message": "Item created successfully",
  "data": {
    "id": "uuid",
    "sku": "SKU-001",
    "name": "Laptop Dell",
    "category": "Electronics",
    "unit": "PCS"
  }
}
```

---

## 9.2 List Items

```http
GET /api/v1/items
```

Query parameters:

```text
?page=1
&limit=10
&category=Electronics
&search=laptop
```

Response:

```json
{
  "success": true,
  "message": "Items retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "sku": "SKU-001",
      "name": "Laptop Dell",
      "category": "Electronics",
      "unit": "PCS"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 42
  }
}
```

---

## 9.3 Get Item

```http
GET /api/v1/items/:id
```

---

## 9.4 Update Item

```http
PUT /api/v1/items/:id
```

Request:

```json
{
  "sku": "SKU-001",
  "name": "Laptop Dell Updated",
  "category": "Electronics",
  "unit": "PCS"
}
```

---

## 9.5 Delete Item

```http
DELETE /api/v1/items/:id
```

Behavior:

```text
deleted_at = NOW()
```

---

## 9.6 Receive Stock

```http
POST /api/v1/stock/receive
```

Request:

```json
{
  "item_id": "uuid",
  "location_id": "uuid",
  "qty": 100
}
```

Flow:

```text
Request
   ↓
Validate Request
   ↓
Find Item
   ↓
Find Location
   ↓
Validate Quantity
   ↓
Find Stock
   ↓
Create / Increase Stock
   ↓
Write Log
   ↓
Commit
```

---

## 9.7 Get Stock by Item

```http
GET /api/v1/stock/:item_id
```

Example response:

```json
{
  "success": true,
  "message": "Stock retrieved successfully",
  "data": [
    {
      "location": {
        "id": "uuid",
        "code": "A-01-01",
        "zone": "A",
        "type": "RACK"
      },
      "qty": 100
    },
    {
      "location": {
        "id": "uuid",
        "code": "B-01-01",
        "zone": "B",
        "type": "RACK"
      },
      "qty": 50
    }
  ]
}
```

---

# 10. API Response Standard

Semua API menggunakan response envelope.

## Success

```json
{
  "success": true,
  "message": "Item retrieved successfully",
  "data": {},
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 42
  }
}
```

## Error

```json
{
  "success": false,
  "message": "SKU already exists",
  "errors": [
    {
      "field": "sku",
      "reason": "Duplicate entry"
    }
  ]
}
```

---

# 11. HTTP Status

| Status | Usage |
|---:|---|
| 200 | Successful GET/UPDATE |
| 201 | Successful CREATE |
| 204 | Successful DELETE, jika digunakan |
| 400 | Validation error |
| 404 | Resource not found |
| 409 | Duplicate/conflict |
| 500 | Internal server error |

---

# 12. Backend Architecture

Required architecture:

```text
Handler / Controller
        ↓
     Service
        ↓
   Repository
        ↓
    PostgreSQL
```

Recommended structure:

```text
backend/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   │   ├── item_handler.go
│   │   └── stock_handler.go
│   │
│   ├── service/
│   │   ├── item_service.go
│   │   └── stock_service.go
│   │
│   ├── repository/
│   │   ├── item_repository.go
│   │   ├── stock_repository.go
│   │   └── location_repository.go
│   │
│   ├── model/
│   │   ├── item.go
│   │   ├── stock.go
│   │   └── location.go
│   │
│   ├── middleware/
│   │   ├── logger.go
│   │   └── error.go
│   │
│   └── response/
│       └── response.go
│
├── migrations/
├── docs/
│   └── swagger.yaml
├── Dockerfile
└── .env.example
```

---

# 13. Responsibility per Layer

## Handler

Responsible for:

- HTTP request
- Parse body
- Input validation
- Calling service
- HTTP response

Handler tidak boleh berisi business logic.

---

## Service

Responsible for:

- Business rules
- SKU validation
- Stock validation
- Transaction orchestration
- Calling repository

Contoh:

```text
ReceiveStock()
├── Validate quantity
├── Validate item
├── Validate location
├── Update stock
└── Write transaction log
```

---

## Repository

Responsible for:

- Database query
- Insert
- Update
- Select
- Delete/soft delete

Repository tidak menentukan business rule.

---

# 14. Middleware

Required:

### Request Logger

Log:

```text
HTTP Method
Path
Status
Duration
```

Example:

```text
POST /api/v1/items 201 15ms
```

### Centralized Error Handler

Semua error dikonversi ke response standard.

### Response Envelope

Response konsisten:

```json
{
  "success": true,
  "message": "...",
  "data": {}
}
```

---

# 15. Frontend PRD

## 15.1 Page — Item Dashboard

Route:

```text
/items
```

UI:

```text
┌──────────────────────────────────────────────────────────┐
│ Warehouse Items                          [+ Add Item]     │
├──────────────────────────────────────────────────────────┤
│ Search SKU / Name     Category ▼                         │
├──────────────────────────────────────────────────────────┤
│ SKU        Name          Category       Unit     Action   │
├──────────────────────────────────────────────────────────┤
│ SKU-001    Laptop        Electronics    PCS      Edit     │
│ SKU-002    Mouse         Electronics    PCS      Edit     │
│ SKU-003    Chair         Furniture      PCS      Edit     │
├──────────────────────────────────────────────────────────┤
│                 < 1 2 3 4 5 >                            │
└──────────────────────────────────────────────────────────┘
```

Features:

- Display SKU
- Display name
- Display category
- Display unit
- Search
- Category filter
- Pagination
- Create
- Edit

---

# 16. Search

Search berdasarkan:

- Name
- SKU

Debounce minimal:

```text
300ms
```

Flow:

```text
User Input
    ↓
Wait 300ms
    ↓
Request API
```

---

# 17. Category Filter

Dropdown:

```text
All Categories
Electronics
Furniture
Food
Tools
...
```

Request:

```http
GET /api/v1/items?category=Electronics
```

---

# 18. Pagination

Backend mendukung pagination sehingga frontend sebaiknya menggunakan server-side pagination.

Request:

```http
GET /api/v1/items?page=2&limit=10
```

Response metadata:

```json
{
  "meta": {
    "page": 2,
    "limit": 10,
    "total": 42
  }
}
```

---

# 19. Create/Edit Item

Routes:

```text
/items/new
/items/:id/edit
```

Form:

```text
SKU
[________________]

Name
[________________]

Category
[________________]

Unit
[________________]

        [Cancel] [Save]
```

Required fields:

- SKU
- Name
- Category
- Unit

Submit:

```text
Create
  ↓
POST /api/v1/items

Edit
  ↓
PUT /api/v1/items/:id
```

---

# 20. Frontend Architecture

```text
frontend/
└── src/
    ├── components/
    │   ├── ItemTable.tsx
    │   ├── ItemForm.tsx
    │   ├── SearchInput.tsx
    │   └── CategoryFilter.tsx
    │
    ├── hooks/
    │   ├── useItems.ts
    │   └── useItem.ts
    │
    ├── pages/
    │   ├── ItemListPage.tsx
    │   └── ItemFormPage.tsx
    │
    ├── services/
    │   └── itemApi.ts
    │
    ├── types/
    │   └── item.ts
    │
    ├── App.tsx
    └── main.tsx
```

Logic API dan state management tidak ditempatkan langsung di component.

---

# 21. Frontend Data Flow

```text
Page
 │
 ├── useItems()
 │      │
 │      ▼
 │   TanStack Query
 │      │
 │      ▼
 │   itemApi.ts
 │      │
 │      ▼
 │   REST API
 │      │
 │      ▼
 │   Go Backend
 │
 └── ItemTable
```

---

# 22. Frontend Technical Requirements

| Component | Requirement |
|---|---|
| Framework | React 18/19 + TypeScript |
| Routing | React Router v6 |
| Data Fetching | TanStack Query |
| UI | shadcn/ui, MUI, Ant Design, atau custom |
| Search | Debounce ≥ 300ms |
| Responsive | Desktop + Tablet |

---

# 23. Recommended Tech Stack

## Backend

```text
Go
HTTP Router: Chi atau Gin
PostgreSQL
pgx
golang-migrate
OpenAPI 3.0
Docker
Docker Compose
```

## Frontend

```text
React 19
TypeScript
Vite
React Router
TanStack Query
shadcn/ui
Tailwind CSS
```

---

# 24. Docker Architecture

```text
                 Browser
                    │
                    ▼
              React Frontend
                    │
                    │ HTTP
                    ▼
              Go REST API
                    │
                    ▼
                PostgreSQL
```

Docker Compose:

```text
docker-compose
│
├── backend
├── frontend
└── postgres
```

Target:

```bash
docker compose up --build
```

harus berjalan tanpa error.

---

# 25. Environment Variables

Gunakan `.env`.

Contoh:

```env
APP_ENV=development

DATABASE_HOST=postgres
DATABASE_PORT=5432
DATABASE_NAME=warehouse
DATABASE_USER=postgres
DATABASE_PASSWORD=secret

PORT=8080
```

Commit:

```text
.env.example
```

Jangan commit:

```text
.env
```

---

# 26. Transaction Strategy

Receive stock sebaiknya menggunakan database transaction.

```text
BEGIN
   │
   ├── Validate item
   ├── Validate location
   ├── Insert/Update stock
   └── Write transaction log
   │
COMMIT
```

Jika terjadi error:

```text
ROLLBACK
```

Tujuannya menjaga konsistensi stock dan transaction log.

---

# 27. Recommended Stock Strategy

Gunakan unique combination:

```text
(item_id, location_id)
```

Contoh:

```text
SKU-001 + A01 = 100
SKU-001 + B01 = 50
```

Jika receive:

```text
SKU-001 + A01 = 20
```

hasil:

```text
SKU-001 + A01 = 120
```

Tidak membuat duplicate stock record.

---

# 28. Testing Strategy

## Unit Test

### Item Service

```text
Create Item
├── valid item
├── duplicate SKU
└── invalid input
```

### Stock Service

```text
Receive Stock
├── valid quantity
├── item not found
├── location not found
└── invalid quantity
```

---

## Integration Test

Test minimal:

```text
POST /api/v1/items
GET /api/v1/items
GET /api/v1/items/:id
PUT /api/v1/items/:id
DELETE /api/v1/items/:id
POST /api/v1/stock/receive
GET /api/v1/stock/:item_id
```

---

# 29. MVP User Flow

## Create Item

```text
User
 ↓
/items
 ↓
Add Item
 ↓
Fill Form
 ↓
Client Validation
 ↓
POST /api/v1/items
 ↓
Backend Validation
 ↓
Service
 ↓
Repository
 ↓
PostgreSQL
 ↓
201 Created
 ↓
Toast Success
 ↓
Item List
```

---

## Receive Stock

```text
Receive Stock
      ↓
Validate Item
      ↓
Validate Location
      ↓
Validate Qty > 0
      ↓
Begin Transaction
      ↓
Find Existing Stock
      ↓
Create / Update
      ↓
Write Log
      ↓
Commit
```

---

# 30. Definition of Done

## Backend

- [ ] Go application running
- [ ] PostgreSQL connected
- [ ] Database migration works
- [ ] Create item
- [ ] List item
- [ ] Search
- [ ] Filter category
- [ ] Pagination
- [ ] Get item
- [ ] Update item
- [ ] Soft delete
- [ ] Receive stock
- [ ] Get stock
- [ ] SKU unique
- [ ] Stock validation
- [ ] Transaction logging
- [ ] Error middleware
- [ ] Request logging
- [ ] Response envelope
- [ ] OpenAPI
- [ ] Docker

## Frontend

- [ ] React + TypeScript
- [ ] React Router
- [ ] TanStack Query
- [ ] Item list
- [ ] Search
- [ ] Debounce 300ms
- [ ] Category filter
- [ ] Pagination
- [ ] Create item
- [ ] Edit item
- [ ] Client validation
- [ ] Loading state
- [ ] Toast
- [ ] Responsive desktop
- [ ] Responsive tablet

---

# 31. Repository Structure

```text
warehouse-management/
│
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   ├── middleware/
│   │   └── response/
│   ├── migrations/
│   ├── docs/
│   ├── Dockerfile
│   └── .env.example
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── hooks/
│   │   ├── pages/
│   │   ├── types/
│   │   └── services/
│   └── package.json
│
├── docker-compose.yml
├── README.md
└── .gitignore
```

---

# 32. README Requirements

README harus berisi:

```text
# Warehouse Management System

## Overview

## Architecture

## Tech Stack

## Database Schema

## API Documentation

## Setup

### Docker

docker compose up --build

### Backend

...

### Frontend

...

## Environment Variables

## API Examples

## Technical Decisions

1. Why Go?
2. Why Layered Architecture?
3. Why PostgreSQL Transaction?

## AI Declaration

AI tools used:
- ChatGPT
- ...

AI assistance:
- Boilerplate
- Documentation
- Scaffolding

Written manually:
- Business logic
- Architecture decisions
- Final implementation
```

---

# 33. Technical Decisions & Trade-offs

## Decision 1 — Go vs Java

### Go

Advantages:

- Cepat dikembangkan.
- Simple deployment.
- Lightweight.
- Cocok untuk REST API.
- Boilerplate relatif sedikit.

### Java Spring Boot

Advantages:

- Ecosystem enterprise matang.
- Banyak library.
- Strong convention.
- Cocok untuk perusahaan Java-centric.

Untuk technical test dengan waktu terbatas, Go dapat digunakan untuk menghasilkan REST API yang sederhana dan terstruktur.

---

## Decision 2 — Layered Architecture

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Database
```

Keuntungan:

- Separation of concerns.
- Business logic tidak bercampur dengan HTTP.
- Repository mudah diganti/test.
- Service mudah di-unit-test.

Trade-off:

- Lebih banyak file/interface.
- Untuk aplikasi sangat kecil bisa terasa verbose.

Namun requirement technical test memang meminta layered architecture.

---

## Decision 3 — PostgreSQL Transaction

Receive stock menggunakan transaction:

```text
BEGIN
 ↓
Update Stock
 ↓
Write Log
 ↓
COMMIT
```

Jika salah satu proses gagal:

```text
ROLLBACK
```

Tujuannya menjaga consistency antara perubahan stock dan proses transaksi.

---

# 34. AI Usage Declaration

AI diperbolehkan dalam technical test, tetapi penggunaannya harus dideklarasikan.

Contoh:

```markdown
## AI Usage Declaration

### Tools Used

- ChatGPT
- GitHub Copilot

### AI Assisted

- Boilerplate generation
- Documentation drafting
- Code explanation
- Initial test case suggestions

### Manually Implemented / Reviewed

- Domain model
- Business rules
- Database design
- Architecture decisions
- Business logic
- Final code review
```

Yang penting kandidat memahami dan dapat menjelaskan kode yang digunakan.

---

# 35. Submission Checklist

Sebelum submission:

```text
[ ] Public GitHub repository
[ ] backend/
[ ] frontend/
[ ] docker-compose.yml
[ ] Dockerfile
[ ] migrations/
[ ] OpenAPI / Swagger
[ ] .env.example
[ ] README.md
[ ] AI declaration
[ ] Technical decisions
[ ] docker compose up --build berhasil
[ ] Backend API berjalan
[ ] Frontend berjalan
[ ] PostgreSQL berjalan
[ ] CRUD item berhasil
[ ] Receive stock berhasil
[ ] Stock query berhasil
[ ] Error handling berhasil
[ ] Soft delete berhasil
[ ] Duplicate SKU menghasilkan 409
```

---

# 36. Implementation Order

Urutan pengerjaan yang direkomendasikan:

```text
1. Understand Requirement
        ↓
2. Domain Model
        ↓
3. ERD
        ↓
4. API Contract
        ↓
5. Go Project Setup
        ↓
6. PostgreSQL + Migration
        ↓
7. Repository
        ↓
8. Service
        ↓
9. Handler
        ↓
10. Middleware
        ↓
11. OpenAPI
        ↓
12. Docker
        ↓
13. React Setup
        ↓
14. API Client
        ↓
15. TanStack Query
        ↓
16. Item List
        ↓
17. Create/Edit Form
        ↓
18. Integration
        ↓
19. Testing
        ↓
20. README
        ↓
21. Final Review
```

---

# 37. Final Review Checklist

## Architecture

- [ ] Handler tidak berisi business logic.
- [ ] Business rule berada di service.
- [ ] Database access berada di repository.
- [ ] API version menggunakan `/api/v1`.
- [ ] Response konsisten.

## Security

- [ ] Tidak ada credential hardcoded.
- [ ] Environment variable digunakan.
- [ ] Input divalidasi.
- [ ] SQL menggunakan parameterized query.
- [ ] `.env` tidak di-commit.

## Database

- [ ] Foreign key benar.
- [ ] SKU unique.
- [ ] `(item_id, location_id)` unique.
- [ ] Soft delete.
- [ ] Transaction untuk receive stock.
- [ ] Migration tersedia.

## Frontend

- [ ] TypeScript types tersedia.
- [ ] API logic berada di service.
- [ ] Server state menggunakan TanStack Query.
- [ ] Search menggunakan debounce.
- [ ] Loading state tersedia.
- [ ] Error state tersedia.
- [ ] Toast tersedia.
- [ ] Responsive.

## DevOps

- [ ] Dockerfile berjalan.
- [ ] docker-compose berjalan.
- [ ] PostgreSQL berjalan.
- [ ] Environment variables tersedia.
- [ ] `docker compose up --build` berhasil.

## Documentation

- [ ] README jelas.
- [ ] Setup instructions tersedia.
- [ ] Architecture dijelaskan.
- [ ] API documentation tersedia.
- [ ] Technical decisions dijelaskan.
- [ ] AI usage dideklarasikan.

---

# 38. Key Principle

Untuk technical test ini:

```text
Jangan mengejar banyak fitur.

Prioritas:

Requirement
    ↓
Correctness
    ↓
Architecture
    ↓
Error Handling
    ↓
Docker
    ↓
Documentation
    ↓
UI Polish
    ↓
Bonus Features
```

Lebih baik menyelesaikan seluruh requirement dengan implementasi yang bersih daripada menambahkan banyak fitur yang belum stabil.
