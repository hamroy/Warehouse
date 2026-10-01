# Warehouse Management System

Sistem manajemen warehouse untuk mengelola item, lokasi, dan stok.

## Struktur Repository

```text
warehouse-management/
├── backend/          # Go REST API
├── frontend/         # React + TypeScript + Vite
├── docker-compose.yml
├── README.md
└── .gitignore
```

## Fitur

- **Item Management** — CRUD item dengan soft delete, search, filter kategori, dan pagination
- **Location** — Master data lokasi warehouse
- **Stock Management** — Receive stock dengan transaksi database dan transaction log
- **API** — REST API dengan response envelope konsisten, validation, dan error handling

## Tech Stack

| Layer    | Stack                        |
| -------- | ---------------------------- |
| Backend  | Go + PostgreSQL              |
| Frontend | React + TypeScript + Vite    |
| Database | PostgreSQL 16                |
| Infra    | Docker + Docker Compose      |

## Quick Start

### Docker (recommended)

```bash
docker compose up --build
```

### Lokal

1. Salin environment file:

   ```bash
   cp backend/.env.example backend/.env
   cp frontend/.env.example frontend/.env
   ```

2. Jalankan PostgreSQL (docker):

   ```bash
   docker compose up postgres
   ```

3. Jalankan backend:

   ```bash
   cd backend
   go run ./cmd/server
   ```

4. Jalankan frontend:

   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## Environment Variables

Lihat `backend/.env.example` dan `frontend/.env.example`.

## Lisensi

Internal — Technical Test
