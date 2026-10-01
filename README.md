# Warehouse Management System

Sistem manajemen warehouse untuk mengelola item, lokasi, dan stok.

## Overview

Warehouse Management System (WMS) adalah aplikasi fullstack untuk mengelola
katalog item warehouse, lokasi penyimpanan, dan stok. Menyediakan REST API
(Go) dan dashboard web (React) dengan fitur CRUD item, pencarian, filter
kategori, pagination, soft delete, dan penerimaan stok dengan pencatatan
transaksi.

## Struktur Repository

```text
warehouse-management/
├── backend/                # Go REST API
│   ├── cmd/
│   │   ├── server/         # entrypoint server (auto-migrate)
│   │   ├── migrate/        # CLI migrasi (up|down)
│   │   └── seed/           # CLI seed data (idempotent)
│   ├── internal/
│   │   ├── api/            # handler + response envelope + error handler
│   │   ├── config/         # loader environment variable
│   │   ├── db/             # migrator + seed + migrations (embedded)
│   │   ├── handler/        # health handler
│   │   ├── item/           # domain, repository, service
│   │   ├── location/       # domain, repository
│   │   ├── server/         # router, middleware, request logger
│   │   └── stock/          # domain, repository, service (transactional)
│   ├── docs/swagger.yaml   # OpenAPI 3.0
│   └── Dockerfile
├── frontend/               # React + TypeScript + Vite
│   ├── src/
│   │   ├── hooks/          # useItems, useDebounce
│   │   ├── pages/items/    # ItemListPage, ItemFormPage
│   │   ├── services/       # itemApi client
│   │   └── lib/api.ts      # apiFetch helper
│   ├── nginx.conf          # SPA fallback
│   └── Dockerfile
├── Docs/PRD/               # dokumen requirement
├── docker-compose.yml
├── README.md
└── .gitignore
```

## Fitur

- **Item Management** — CRUD item dengan soft delete, search (SKU/name),
  filter kategori, dan pagination
- **Location** — Master data lokasi warehouse (disediakan via seed)
- **Stock Management** — Receive stock dalam transaksi database
  (create-or-increase) dengan transaction log
- **API Infrastructure** — Response envelope konsisten, request validation,
  centralized error handler, request logger

## Tech Stack

| Layer    | Stack                                   |
| -------- | --------------------------------------- |
| Backend  | Go 1.23 + chi router + lib/pq         |
| Frontend | React 19 + TypeScript + Vite + TanStack Query |
| Database | PostgreSQL 16                           |
| Infra    | Docker + Docker Compose                 |

## Architecture

Backend mengikuti layered architecture:

```text
HTTP Request
    ↓
Handler (internal/api)      — parse request, panggil service
    ↓
Service (internal/*/service) — business logic, validasi, transaksi
    ↓
Repository (internal/*/repository) — data access, parameterized query
    ↓
PostgreSQL
```

- **Handler**: tidak berisi business logic
- **Service**: orkestrasi aturan bisnis (misal: duplicate SKU → 409,
  receive stock dalam satu transaksi)
- **Repository**: hanya SQL, tidak berisi aturan bisnis

## API

Base URL: `http://localhost:8080`

Semua endpoint menggunakan response envelope:

```json
{ "success": true, "message": "...", "data": {} }
```

Error:

```json
{ "success": false, "message": "Validation failed", "errors": ["..."] }
```

| Method | Endpoint | Deskripsi |
| ------ | -------- | --------- |
| GET | `/health` | Health check |
| POST | `/api/v1/items` | Buat item (201/400/409) |
| GET | `/api/v1/items` | List item (search, category, page, limit) |
| GET | `/api/v1/items/:id` | Detail item (404 jika tidak ada/soft-deleted) |
| PUT | `/api/v1/items/:id` | Update item |
| DELETE | `/api/v1/items/:id` | Soft delete item |
| POST | `/api/v1/stock/receive` | Terima stok (400 qty≤0, 404 item/location) |
| GET | `/api/v1/stock/:item_id` | Stok item per lokasi + total |

Dokumentasi lengkap: [backend/docs/swagger.yaml](backend/docs/swagger.yaml)

Contoh receive stock:

```bash
curl -X POST http://localhost:8080/api/v1/stock/receive \
  -H "Content-Type: application/json" \
  -d '{"item_id":"<uuid>","location_id":"<uuid>","qty":100}'
```

## Database

Tiga tabel utama (lihat `backend/internal/db/migrations/`):

- **items** — `id UUID PK`, `sku UNIQUE`, `name`, `category`, `unit`,
  timestamps, `deleted_at` (soft delete)
- **locations** — `id UUID PK`, `code UNIQUE`, `zone`, `type`
- **stocks** — `id UUID PK`, `item_id FK`, `location_id FK`, `qty`,
  `UNIQUE (item_id, location_id)`, `CHECK (qty >= 0)`
- **stock_logs** — log transaksi: `type`, `item_id`, `location_id`, `qty`

Migrasi di-embed dalam binary dan dijalankan otomatis saat server start.
Migrasi manual:

```bash
cd backend
go run ./cmd/migrate up    # apply
go run ./cmd/migrate down  # rollback last
```

Seed data (idempoten, `ON CONFLICT DO NOTHING`):

```bash
cd backend
go run ./cmd/seed
```

Seed: locations `A-01-01`, `A-01-02`, `B-01-01`; items `SKU-001` Laptop Dell,
`SKU-002` Mouse, `SKU-003` Office Chair.

## Setup

### Docker (recommended)

```bash
docker compose up --build
```

Kemudian buka:

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- OpenAPI: lihat `backend/docs/swagger.yaml`

### Lokal (tanpa Docker)

Prerequisite: Go 1.23+, Node 18+, PostgreSQL 16.

```bash
# 1. environment
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env

# 2. database (docker)
docker compose up postgres

# 3. backend — migrasi otomatis saat start
cd backend
go run ./cmd/server

# 4. seed (opsional, idempoten)
go run ./cmd/seed

# 5. frontend
cd frontend
npm install
npm run dev
```

## Environment Variables

### Backend (`backend/.env.example`)

| Variable | Default | Deskripsi |
| -------- | ------- | --------- |
| `PORT` | `8080` | Port server |
| `DB_HOST` | `localhost` | Host PostgreSQL |
| `DB_PORT` | `5432` | Port PostgreSQL |
| `DB_USER` | `warehouse` | User database |
| `DB_PASSWORD` | `warehouse` | Password database |
| `DB_NAME` | `warehouse` | Nama database |
| `DB_SSLMODE` | `disable` | SSL mode koneksi |

### Frontend (`frontend/.env.example`)

| Variable | Default | Deskripsi |
| -------- | ------- | --------- |
| `VITE_API_URL` | `http://localhost:8080` | URL backend API |

**Tidak ada credential yang di-hardcode** — semua nilai dibaca dari
environment variable. File `.env*` (kecuali `.env.example`) di-ignore oleh git.

## Docker

```bash
# build semua image
docker compose build

# jalankan semua service
docker compose up -d

# lihat log
docker compose logs -f backend

# hentikan + hapus volume
docker compose down -v
```

Service:

| Service | Port | Deskripsi |
| ------- | ---- | --------- |
| `postgres` | 5432 | Database (healthcheck) |
| `backend` | 8080 | Go API (auto-migrate) |
| `frontend` | 5173 | React via nginx |

## Testing

Backend (Go):

```bash
cd backend
go build ./...
go vet ./...
go test ./...
```

Frontend:

```bash
cd frontend
npm run build
```

Manual QA checklist ada di `Docs/PRD/warehouse-management-task-breakdown.md`
(TASK-1301..1303).

## Author

**Ilham Royani**

- GitHub: https://github.com/hamroy
- Repository: https://github.com/hamroy/Warehouse
