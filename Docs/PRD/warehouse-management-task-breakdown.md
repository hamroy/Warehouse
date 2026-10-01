# Technical Test — Warehouse Management System
# Task Breakdown

> Task breakdown berdasarkan requirement technical test.  
> Fokus: implementasi bertahap, feature-slice, dan setiap task memiliki acceptance criteria.

---

## 0. Project Foundation

### TASK-001 — Initialize Repository

**Priority:** P0  
**Module:** Foundation

#### Goal

Menyiapkan struktur repository untuk backend, frontend, Docker, dan documentation.

#### Tasks

- [ ] Buat repository.
- [ ] Buat folder `backend/`.
- [ ] Buat folder `frontend/`.
- [ ] Buat `docker-compose.yml`.
- [ ] Buat `.gitignore`.
- [ ] Buat `README.md`.

#### Expected Structure

```text
warehouse-management/
├── backend/
├── frontend/
├── docker-compose.yml
├── README.md
└── .gitignore
```

#### Acceptance Criteria

- [ ] Repository dapat di-clone.
- [ ] Struktur folder tersedia.
- [ ] Tidak ada secret di repository.
- [ ] Initial commit berhasil.

---

### TASK-002 — Backend Project Setup

**Priority:** P0  
**Module:** Foundation

#### Tasks

- [ ] Initialize Go project.
- [ ] Setup HTTP server.
- [ ] Setup router.
- [ ] Setup configuration.
- [ ] Setup environment variables.
- [ ] Buat health endpoint.

#### API

```http
GET /health
```

#### Acceptance Criteria

- [ ] Backend dapat dijalankan.
- [ ] Port configurable.
- [ ] `/health` mengembalikan `200`.
- [ ] Tidak ada credential hardcoded.

---

### TASK-003 — Frontend Project Setup

**Priority:** P0  
**Module:** Foundation

#### Tasks

- [ ] Setup React.
- [ ] Setup TypeScript.
- [ ] Setup Vite.
- [ ] Setup React Router.
- [ ] Setup TanStack Query.
- [ ] Setup UI library.
- [ ] Setup Tailwind CSS jika digunakan.

#### Acceptance Criteria

- [ ] `npm install` berhasil.
- [ ] `npm run dev` berhasil.
- [ ] React application tampil.
- [ ] Router berjalan.
- [ ] Query provider tersedia.

---

# 1. Database & Domain

### TASK-101 — Design Database

**Priority:** P0  
**Module:** Database

#### Entities

```text
items
locations
stocks
```

#### Tasks

- [ ] Design `items`.
- [ ] Design `locations`.
- [ ] Design `stocks`.
- [ ] Define primary keys.
- [ ] Define foreign keys.
- [ ] Define unique constraints.
- [ ] Define timestamps.
- [ ] Define soft delete pada item.

#### Acceptance Criteria

- [ ] Item memiliki UUID.
- [ ] SKU unique.
- [ ] Stock memiliki `item_id`.
- [ ] Stock memiliki `location_id`.
- [ ] `(item_id, location_id)` unique.
- [ ] Item memiliki `deleted_at`.

---

### TASK-102 — Create Database Migration

**Priority:** P0  
**Module:** Database

#### Migration

```text
001_create_items.sql
002_create_locations.sql
003_create_stocks.sql
```

#### Acceptance Criteria

- [ ] Migration dari database kosong berhasil.
- [ ] Foreign key berhasil dibuat.
- [ ] Unique constraint berhasil dibuat.
- [ ] Rollback dapat dilakukan.

---

### TASK-103 — Create Database Seed

**Priority:** P1  
**Module:** Database

#### Seed Data

Locations:

```text
A-01-01
A-01-02
B-01-01
```

Items:

```text
SKU-001 | Laptop Dell   | Electronics | PCS
SKU-002 | Mouse         | Electronics | PCS
SKU-003 | Office Chair  | Furniture   | PCS
```

#### Acceptance Criteria

- [ ] Seed dapat dijalankan.
- [ ] Data tersedia untuk development/testing.
- [ ] Seed tidak membuat duplicate data saat dijalankan sesuai strategy yang dipilih.

---

# 2. Item Management

## TASK-201 — Item Domain Model

**Priority:** P0  
**Module:** Item

#### Fields

```text
id
sku
name
category
unit
created_at
updated_at
deleted_at
```

#### Acceptance Criteria

- [ ] Model tersedia.
- [ ] UUID digunakan.
- [ ] Soft delete didukung.
- [ ] Field sesuai database.

---

## TASK-202 — Item Repository

**Priority:** P0  
**Module:** Item

#### Methods

```text
Create
FindByID
FindAll
Update
SoftDelete
ExistsBySKU
```

#### Acceptance Criteria

- [ ] Repository dapat melakukan CRUD.
- [ ] Repository menangani database access.
- [ ] Business logic tidak berada di repository.
- [ ] Query menggunakan parameterized query.

---

## TASK-203 — Create Item Service

**Priority:** P0  
**Module:** Item

#### Flow

```text
Request
   ↓
Validate
   ↓
Check SKU
   ↓
Create Item
```

#### Business Rules

```text
SKU duplicate → 409 Conflict
```

#### Acceptance Criteria

- [ ] Required fields divalidasi.
- [ ] Duplicate SKU ditolak.
- [ ] Item berhasil dibuat.
- [ ] Error dikembalikan secara konsisten.

---

## TASK-204 — Create Item API

**Priority:** P0  
**Module:** Item API

#### Endpoint

```http
POST /api/v1/items
```

#### Request

```json
{
  "sku": "SKU-001",
  "name": "Laptop Dell",
  "category": "Electronics",
  "unit": "PCS"
}
```

#### Acceptance Criteria

- [ ] Valid request menghasilkan `201`.
- [ ] Invalid request menghasilkan `400`.
- [ ] Duplicate SKU menghasilkan `409`.
- [ ] Response menggunakan standard envelope.

---

## TASK-205 — List Item Service

**Priority:** P0  
**Module:** Item

#### Support

```text
search
category
page
limit
```

#### Search

Search berdasarkan:

```text
SKU
Name
```

#### Acceptance Criteria

- [ ] Search SKU berjalan.
- [ ] Search name berjalan.
- [ ] Category filter berjalan.
- [ ] Pagination berjalan.
- [ ] Soft deleted item tidak muncul.

---

## TASK-206 — List Item API

**Priority:** P0  
**Module:** Item API

#### Endpoint

```http
GET /api/v1/items
```

#### Example

```http
GET /api/v1/items?page=1&limit=10&search=laptop&category=Electronics
```

#### Acceptance Criteria

- [ ] Response `200`.
- [ ] Data tersedia.
- [ ] Pagination metadata tersedia.
- [ ] Search bekerja.
- [ ] Category filter bekerja.

---

## TASK-207 — Get Item Detail

**Priority:** P0  
**Module:** Item API

#### Endpoint

```http
GET /api/v1/items/:id
```

#### Acceptance Criteria

- [ ] Existing item → `200`.
- [ ] Item tidak ditemukan → `404`.
- [ ] Soft deleted item tidak dianggap active.

---

## TASK-208 — Update Item

**Priority:** P0  
**Module:** Item

#### Endpoint

```http
PUT /api/v1/items/:id
```

#### Acceptance Criteria

- [ ] Item dapat di-update.
- [ ] Required fields divalidasi.
- [ ] SKU uniqueness tetap dijaga.
- [ ] Item tidak ditemukan → `404`.
- [ ] Success → `200`.

---

## TASK-209 — Soft Delete Item

**Priority:** P0  
**Module:** Item

#### Endpoint

```http
DELETE /api/v1/items/:id
```

#### Behavior

```text
deleted_at = current timestamp
```

#### Acceptance Criteria

- [ ] Item tidak dihapus secara physical.
- [ ] `deleted_at` terisi.
- [ ] Item tidak muncul pada list active.
- [ ] Item tidak dapat diperlakukan sebagai active item.

---

# 3. Location

> Endpoint CRUD Location tidak menjadi prioritas MVP karena location endpoint termasuk fitur yang dapat dihilangkan untuk menghemat waktu. Location tetap diperlukan sebagai master data untuk Stock dan dapat disediakan melalui seed.

## TASK-301 — Location Domain Model

**Priority:** P0  
**Module:** Location

#### Fields

```text
id
code
zone
type
```

#### Acceptance Criteria

- [ ] Model tersedia.
- [ ] UUID digunakan.
- [ ] Field sesuai database.

---

## TASK-302 — Location Repository

**Priority:** P0  
**Module:** Location

#### Methods

```text
FindByID
ExistsByID
```

#### Acceptance Criteria

- [ ] Repository dapat mencari location.
- [ ] Location tidak ditemukan menghasilkan error yang dapat diproses service.

---

## TASK-303 — Location Seed

**Priority:** P0  
**Module:** Location

#### Data

```text
A-01-01
A-01-02
B-01-01
```

#### Acceptance Criteria

- [ ] Location tersedia setelah seed.
- [ ] Stock dapat menggunakan location tersebut.

---

# 4. Stock Management

## TASK-401 — Stock Domain Model

**Priority:** P0  
**Module:** Stock

#### Fields

```text
id
item_id
location_id
qty
updated_at
```

#### Acceptance Criteria

- [ ] Model tersedia.
- [ ] Foreign key tersedia.
- [ ] Quantity dapat disimpan.
- [ ] Unique `(item_id, location_id)` tersedia.

---

## TASK-402 — Stock Repository

**Priority:** P0  
**Module:** Stock

#### Methods

```text
FindByItemAndLocation
Create
IncreaseQuantity
FindByItem
```

#### Acceptance Criteria

- [ ] Existing stock dapat ditemukan.
- [ ] Stock baru dapat dibuat.
- [ ] Quantity dapat ditambahkan.
- [ ] Stock item dapat diambil berdasarkan item.

---

## TASK-403 — Receive Stock Service

**Priority:** P0  
**Module:** Stock

#### Flow

```text
Receive Stock
      ↓
Validate Item
      ↓
Validate Location
      ↓
Validate Quantity
      ↓
BEGIN TRANSACTION
      ↓
Find Existing Stock
      ↓
Create / Increase Stock
      ↓
Write Log
      ↓
COMMIT
```

#### Business Rules

```text
qty <= 0 → reject
```

#### Acceptance Criteria

- [ ] Item harus valid.
- [ ] Location harus valid.
- [ ] Quantity harus lebih dari 0.
- [ ] Existing stock di-increase.
- [ ] Stock belum ada → create.
- [ ] Database transaction digunakan.
- [ ] Error → rollback.

---

## TASK-404 — Receive Stock API

**Priority:** P0  
**Module:** Stock API

#### Endpoint

```http
POST /api/v1/stock/receive
```

#### Request

```json
{
  "item_id": "uuid",
  "location_id": "uuid",
  "qty": 100
}
```

#### Acceptance Criteria

- [ ] Valid request berhasil.
- [ ] Item invalid → `404`.
- [ ] Location invalid → `404`.
- [ ] `qty <= 0` → `400`.
- [ ] Stock bertambah.
- [ ] Response standard.

---

## TASK-405 — Get Stock by Item

**Priority:** P0  
**Module:** Stock API

#### Endpoint

```http
GET /api/v1/stock/:item_id
```

#### Example

```text
SKU-001

A-01-01 → 100
A-01-02 → 50
B-01-01 → 25

Total → 175
```

#### Acceptance Criteria

- [ ] Stock semua location ditampilkan.
- [ ] Quantity benar.
- [ ] Item tidak ditemukan → `404`.
- [ ] Response standard.

---

## TASK-406 — Stock Transaction Logging

**Priority:** P0  
**Module:** Stock

#### Log

```text
RECEIVE_STOCK
item_id
location_id
qty
timestamp
```

#### Acceptance Criteria

- [ ] Receive stock menghasilkan log.
- [ ] Log berisi item.
- [ ] Log berisi location.
- [ ] Log berisi quantity.
- [ ] Log berisi timestamp.

---

# 5. API Infrastructure

## TASK-501 — Standard Response Envelope

**Priority:** P0  
**Module:** API Infrastructure

### Success

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

### Error

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": []
}
```

### Pagination

```json
{
  "success": true,
  "message": "Items retrieved successfully",
  "data": [],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 42
  }
}
```

#### Acceptance Criteria

- [ ] Semua endpoint menggunakan format konsisten.
- [ ] Error memiliki format konsisten.
- [ ] Pagination memiliki metadata.

---

## TASK-502 — Request Validation

**Priority:** P0  
**Module:** API Infrastructure

#### Acceptance Criteria

- [ ] Request body divalidasi.
- [ ] Required field divalidasi.
- [ ] Invalid data menghasilkan `400`.
- [ ] Validation error memiliki informasi field.

---

## TASK-503 — Centralized Error Handler

**Priority:** P0  
**Module:** API Infrastructure

### Mapping

```text
ValidationError → 400
NotFoundError   → 404
ConflictError   → 409
InternalError   → 500
```

#### Acceptance Criteria

- [ ] Tidak ada error response yang berbeda-beda antar endpoint.
- [ ] Internal error tidak membocorkan detail database.
- [ ] Error tercatat di log.

---

## TASK-504 — Request Logger

**Priority:** P1  
**Module:** API Infrastructure

### Log

```text
METHOD
PATH
STATUS
DURATION
```

Example:

```text
POST /api/v1/items 201 12ms
```

#### Acceptance Criteria

- [ ] Request tercatat.
- [ ] HTTP method tercatat.
- [ ] Path tercatat.
- [ ] Status tercatat.
- [ ] Duration tercatat.

---

# 6. OpenAPI

## TASK-601 — OpenAPI Setup

**Priority:** P0  
**Module:** Documentation

#### File

```text
backend/docs/swagger.yaml
```

#### Acceptance Criteria

- [ ] OpenAPI 3.0 tersedia.
- [ ] API server tercantum.
- [ ] API version tercantum.

---

## TASK-602 — Document Item API

**Priority:** P0  
**Module:** Documentation

Document:

```text
POST   /api/v1/items
GET    /api/v1/items
GET    /api/v1/items/:id
PUT    /api/v1/items/:id
DELETE /api/v1/items/:id
```

#### Acceptance Criteria

- [ ] Request schema tersedia.
- [ ] Response schema tersedia.
- [ ] Error response tersedia.
- [ ] Query parameter tersedia.

---

## TASK-603 — Document Stock API

**Priority:** P0  
**Module:** Documentation

Document:

```text
POST /api/v1/stock/receive
GET  /api/v1/stock/:item_id
```

#### Acceptance Criteria

- [ ] Request schema tersedia.
- [ ] Response schema tersedia.
- [ ] Error response tersedia.

---

# 7. Frontend — Item Dashboard

## TASK-701 — Item List Page

**Priority:** P0  
**Module:** Frontend Item

#### Route

```text
/items
```

#### UI

```text
Warehouse Items

[Search SKU / Name] [Category ▼] [+ Add Item]

------------------------------------------------
SKU       Name       Category       Unit   Action
------------------------------------------------
SKU-001   Laptop     Electronics    PCS    Edit
SKU-002   Mouse      Electronics    PCS    Edit
------------------------------------------------

             < 1 2 3 4 5 >
```

#### Acceptance Criteria

- [ ] Item list tampil.
- [ ] SKU tampil.
- [ ] Name tampil.
- [ ] Category tampil.
- [ ] Unit tampil.
- [ ] Edit action tersedia.
- [ ] Add Item tersedia.

---

## TASK-702 — Item API Client

**Priority:** P0  
**Module:** Frontend API

#### File

```text
src/services/itemApi.ts
```

#### Methods

```text
getItems()
getItem()
createItem()
updateItem()
deleteItem()
```

#### Acceptance Criteria

- [ ] API call terpisah dari component.
- [ ] TypeScript types digunakan.
- [ ] Error API dapat diproses.

---

## TASK-703 — useItems Hook

**Priority:** P0  
**Module:** Frontend State

#### Flow

```text
useItems()
    ↓
TanStack Query
    ↓
itemApi.getItems()
    ↓
REST API
```

#### Acceptance Criteria

- [ ] Server state menggunakan TanStack Query.
- [ ] Loading state tersedia.
- [ ] Error state tersedia.
- [ ] Refetch/invalidation dapat dilakukan.

---

## TASK-704 — Search Item

**Priority:** P0  
**Module:** Frontend Search

#### Requirement

```text
Search SKU
Search Name
Debounce >= 300ms
```

#### Acceptance Criteria

- [ ] Search SKU bekerja.
- [ ] Search Name bekerja.
- [ ] Input menggunakan debounce.
- [ ] API tidak dipanggil setiap keypress.
- [ ] Search dapat dikombinasikan dengan pagination.

---

## TASK-705 — Category Filter

**Priority:** P0  
**Module:** Frontend Filter

#### Acceptance Criteria

- [ ] Dropdown category tersedia.
- [ ] All Categories tersedia.
- [ ] Filter mengubah query API.
- [ ] Pagination di-reset ketika filter berubah.

---

## TASK-706 — Pagination

**Priority:** P0  
**Module:** Frontend Pagination

#### Acceptance Criteria

- [ ] Page number tersedia.
- [ ] Next tersedia.
- [ ] Previous tersedia.
- [ ] Current page terlihat.
- [ ] Total data digunakan untuk menentukan pagination.

---

# 8. Frontend — Item Form

## TASK-801 — Create Item Page

**Priority:** P0  
**Module:** Frontend Item

#### Route

```text
/items/new
```

#### Fields

```text
SKU
Name
Category
Unit
```

#### Acceptance Criteria

- [ ] Semua field tersedia.
- [ ] Required validation tersedia.
- [ ] Submit button tersedia.
- [ ] Cancel button tersedia.

---

## TASK-802 — Client Validation

**Priority:** P0  
**Module:** Frontend Validation

#### Rules

```text
SKU       → required
Name      → required
Category  → required
Unit      → required
```

#### Acceptance Criteria

- [ ] Empty field ditolak.
- [ ] Error ditampilkan dekat field.
- [ ] API tidak dipanggil jika form invalid.

---

## TASK-803 — Create Item Mutation

**Priority:** P0  
**Module:** Frontend Item

#### Flow

```text
Form
 ↓
Validation
 ↓
Mutation
 ↓
POST /api/v1/items
 ↓
Success
 ↓
Invalidate Items
 ↓
Toast
 ↓
Redirect
```

#### Acceptance Criteria

- [ ] Item berhasil dibuat.
- [ ] Success notification muncul.
- [ ] Item list ter-update.
- [ ] Duplicate SKU ditampilkan sebagai error.

---

## TASK-804 — Edit Item Page

**Priority:** P0  
**Module:** Frontend Item

#### Route

```text
/items/:id/edit
```

#### Flow

```text
GET item
   ↓
Populate form
   ↓
Edit
   ↓
PUT /api/v1/items/:id
```

#### Acceptance Criteria

- [ ] Existing data tampil.
- [ ] Data dapat diedit.
- [ ] Update berhasil.
- [ ] Error ditampilkan.
- [ ] List ter-update.

---

## TASK-805 — Loading & Error State

**Priority:** P1  
**Module:** Frontend UX

#### States

```text
Loading
Success
Error
Empty
```

#### Acceptance Criteria

- [ ] Loading state tersedia.
- [ ] Error state tersedia.
- [ ] Empty state tersedia.
- [ ] User mendapat feedback yang jelas.

---

# 9. UI/UX

## TASK-901 — Loading Skeleton

**Priority:** P1  
**Module:** UI

#### Acceptance Criteria

- [ ] Loading state tidak menampilkan blank screen.
- [ ] Table/form memiliki feedback loading.

---

## TASK-902 — Toast Notification

**Priority:** P1  
**Module:** UI

### Success

```text
Item created successfully
Item updated successfully
```

### Error

```text
SKU already exists
Failed to create item
```

#### Acceptance Criteria

- [ ] Success toast tersedia.
- [ ] Error toast tersedia.
- [ ] Toast tidak menghalangi interaction.

---

## TASK-903 — Responsive Layout

**Priority:** P1  
**Module:** UI

#### Target

```text
Desktop
Tablet
```

#### Acceptance Criteria

- [ ] Dashboard usable pada desktop.
- [ ] Dashboard usable pada tablet.
- [ ] Form usable pada desktop.
- [ ] Form usable pada tablet.
- [ ] Tidak ada horizontal overflow yang tidak diperlukan.

---

# 10. Docker & Environment

## TASK-1001 — Backend Dockerfile

**Priority:** P0  
**Module:** DevOps

#### Acceptance Criteria

- [ ] Backend dapat di-build.
- [ ] Backend dapat dijalankan dalam container.
- [ ] Environment variable dapat digunakan.

---

## TASK-1002 — Frontend Dockerfile

**Priority:** P0  
**Module:** DevOps

#### Acceptance Criteria

- [ ] Frontend dapat di-build.
- [ ] Frontend dapat dijalankan dalam container.
- [ ] API URL configurable.

---

## TASK-1003 — PostgreSQL Container

**Priority:** P0  
**Module:** DevOps

#### Acceptance Criteria

- [ ] PostgreSQL container berjalan.
- [ ] Database dibuat otomatis.
- [ ] Backend dapat connect ke PostgreSQL.

---

## TASK-1004 — Docker Compose

**Priority:** P0  
**Module:** DevOps

#### Services

```text
backend
frontend
postgres
```

#### Command

```bash
docker compose up --build
```

#### Acceptance Criteria

- [ ] Semua service berhasil start.
- [ ] Backend connect ke database.
- [ ] Frontend dapat mengakses backend.
- [ ] Tidak membutuhkan setup manual yang tidak terdokumentasi.

---

# 11. Testing

## TASK-1101 — Item Service Unit Test

**Priority:** P1  
**Module:** Testing

### Test Cases

```text
Create Item
├── valid item
├── duplicate SKU
└── invalid input
```

#### Acceptance Criteria

- [ ] Happy path tested.
- [ ] Duplicate SKU tested.
- [ ] Invalid input tested.

---

## TASK-1102 — Stock Service Unit Test

**Priority:** P1  
**Module:** Testing

### Test Cases

```text
Receive Stock
├── valid quantity
├── qty <= 0
├── item not found
└── location not found
```

#### Acceptance Criteria

- [ ] Happy path tested.
- [ ] Invalid quantity tested.
- [ ] Item not found tested.
- [ ] Location not found tested.

---

## TASK-1103 — API Integration Test

**Priority:** P1  
**Module:** Testing

### Endpoints

```text
POST   /api/v1/items
GET    /api/v1/items
GET    /api/v1/items/:id
PUT    /api/v1/items/:id
DELETE /api/v1/items/:id
POST   /api/v1/stock/receive
GET    /api/v1/stock/:item_id
```

#### Acceptance Criteria

- [ ] Endpoint success tested.
- [ ] Validation error tested.
- [ ] Not found tested.
- [ ] Conflict tested.

---

# 12. Documentation

## TASK-1201 — README

**Priority:** P0  
**Module:** Documentation

#### Sections

```text
Overview
Architecture
Tech Stack
Database
API
Setup
Environment
Docker
Testing
Technical Decisions
AI Declaration
```

#### Acceptance Criteria

- [ ] Developer baru dapat menjalankan project mengikuti README.
- [ ] Architecture dijelaskan.
- [ ] Setup dijelaskan.
- [ ] Environment variables dijelaskan.
- [ ] Testing dijelaskan.

---

## TASK-1202 — Technical Decisions

**Priority:** P0  
**Module:** Documentation

Minimal jelaskan:

```text
1. Architecture
2. Database
3. Transaction strategy
```

Format:

```markdown
## Decision

### Problem

### Decision

### Alternative

### Trade-off
```

---

## TASK-1203 — AI Usage Declaration

**Priority:** P0  
**Module:** Documentation

#### Acceptance Criteria

- [ ] AI tools yang digunakan disebutkan.
- [ ] Bagian yang dibantu AI disebutkan.
- [ ] Bagian yang ditulis/review manual disebutkan.
- [ ] Kandidat memahami seluruh code yang disubmit.

---

# 13. Final QA

## TASK-1301 — Backend QA

**Priority:** P0  
**Module:** QA

### Checklist

- [ ] Health check
- [ ] Create Item
- [ ] Duplicate SKU
- [ ] List Item
- [ ] Search
- [ ] Category Filter
- [ ] Pagination
- [ ] Get Item
- [ ] Update Item
- [ ] Soft Delete
- [ ] Receive Stock
- [ ] Get Stock
- [ ] Error Handling
- [ ] Request Logging
- [ ] OpenAPI

---

## TASK-1302 — Frontend QA

**Priority:** P0  
**Module:** QA

### Checklist

- [ ] Item List
- [ ] Search
- [ ] Debounce
- [ ] Category Filter
- [ ] Pagination
- [ ] Create Item
- [ ] Edit Item
- [ ] Validation
- [ ] Loading State
- [ ] Empty State
- [ ] Error State
- [ ] Toast
- [ ] Responsive

---

## TASK-1303 — Docker QA

**Priority:** P0  
**Module:** QA

### Fresh Environment Test

```bash
git clone <repository>

cd warehouse-management

docker compose up --build
```

### Verify

```text
Frontend → OK
Backend  → OK
Database → OK
API      → OK
```

#### Acceptance Criteria

- [ ] Fresh clone dapat dijalankan.
- [ ] Docker build berhasil.
- [ ] Database migration berhasil.
- [ ] Frontend dapat dibuka.
- [ ] API dapat dipanggil.
- [ ] CRUD Item berjalan.
- [ ] Stock berjalan.

---

# 14. Task Priority

## P0 — Mandatory

```text
TASK-001  Initialize Repository
TASK-002  Backend Setup
TASK-003  Frontend Setup

TASK-101  Database Design
TASK-102  Database Migration

TASK-201  Item Domain
TASK-202  Item Repository
TASK-203  Create Item Service
TASK-204  Create Item API
TASK-205  List Item Service
TASK-206  List Item API
TASK-207  Get Item
TASK-208  Update Item
TASK-209  Soft Delete

TASK-301  Location Model
TASK-302  Location Repository
TASK-303  Location Seed

TASK-401  Stock Model
TASK-402  Stock Repository
TASK-403  Receive Stock Service
TASK-404  Receive Stock API
TASK-405  Get Stock
TASK-406  Stock Logging

TASK-501  Response Envelope
TASK-502  Validation
TASK-503  Error Handler

TASK-601  OpenAPI Setup
TASK-602  Item API Documentation
TASK-603  Stock API Documentation

TASK-701  Item Dashboard
TASK-702  API Client
TASK-703  useItems
TASK-704  Search
TASK-705  Category Filter
TASK-706  Pagination

TASK-801  Create Item
TASK-802  Validation
TASK-803  Create Mutation
TASK-804  Edit Item

TASK-1001 Backend Docker
TASK-1002 Frontend Docker
TASK-1003 PostgreSQL
TASK-1004 Docker Compose

TASK-1201 README
TASK-1202 Technical Decisions
TASK-1203 AI Declaration

TASK-1301 Backend QA
TASK-1302 Frontend QA
TASK-1303 Docker QA
```

---

# 15. P1 — Quality

```text
TASK-103   Database Seed
TASK-504   Request Logger

TASK-805   Loading/Error State
TASK-901   Loading Skeleton
TASK-902   Toast
TASK-903   Responsive

TASK-1101  Item Unit Test
TASK-1102  Stock Unit Test
TASK-1103  API Integration Test
```

---

# 16. P2 — Optional / Bonus

```text
Stock Transfer
Location CRUD API
Item Detail + Stock Page
Redis Cache
Global Error Boundary
Deployment
CI/CD
```

---

# 17. Recommended Execution Order

```text
PHASE 1 — FOUNDATION
│
├── TASK-001
├── TASK-002
└── TASK-003
        ↓
PHASE 2 — DATABASE
│
├── TASK-101
├── TASK-102
└── TASK-103
        ↓
PHASE 3 — ITEM
│
├── TASK-201
├── TASK-202
├── TASK-203
├── TASK-204
├── TASK-205
├── TASK-206
├── TASK-207
├── TASK-208
└── TASK-209
        ↓
PHASE 4 — LOCATION
│
├── TASK-301
├── TASK-302
└── TASK-303
        ↓
PHASE 5 — STOCK
│
├── TASK-401
├── TASK-402
├── TASK-403
├── TASK-404
├── TASK-405
└── TASK-406
        ↓
PHASE 6 — API INFRASTRUCTURE
│
├── TASK-501
├── TASK-502
├── TASK-503
└── TASK-504
        ↓
PHASE 7 — OPENAPI
│
├── TASK-601
├── TASK-602
└── TASK-603
        ↓
PHASE 8 — FRONTEND
│
├── TASK-701
├── TASK-702
├── TASK-703
├── TASK-704
├── TASK-705
├── TASK-706
├── TASK-801
├── TASK-802
├── TASK-803
└── TASK-804
        ↓
PHASE 9 — UI/UX
│
├── TASK-805
├── TASK-901
├── TASK-902
└── TASK-903
        ↓
PHASE 10 — DOCKER
│
├── TASK-1001
├── TASK-1002
├── TASK-1003
└── TASK-1004
        ↓
PHASE 11 — TESTING
│
├── TASK-1101
├── TASK-1102
└── TASK-1103
        ↓
PHASE 12 — DOCUMENTATION
│
├── TASK-1201
├── TASK-1202
└── TASK-1203
        ↓
PHASE 13 — FINAL QA
│
├── TASK-1301
├── TASK-1302
└── TASK-1303
        ↓
      DONE
```

---

# 18. Definition of Done

Sebuah task dianggap **DONE** jika:

- [ ] Requirement task sudah diimplementasikan.
- [ ] Acceptance criteria terpenuhi.
- [ ] Tidak ada error pada happy path.
- [ ] Error case yang relevan sudah ditangani.
- [ ] Code mengikuti architecture yang ditentukan.
- [ ] Tidak ada secret hardcoded.
- [ ] Tidak menggunakan `any`.
- [ ] Tidak menggunakan `@ts-ignore`.
- [ ] Test dibuat jika task masuk scope testing.
- [ ] Dokumentasi diperbarui jika behavior/API berubah.

---

# 19. Final Submission Checklist

```text
[ ] Repository dapat di-clone
[ ] README tersedia
[ ] .env.example tersedia
[ ] Docker Compose tersedia
[ ] Backend berjalan
[ ] Frontend berjalan
[ ] PostgreSQL berjalan
[ ] Migration berjalan
[ ] Item CRUD berjalan
[ ] Search berjalan
[ ] Filter berjalan
[ ] Pagination berjalan
[ ] Soft delete berjalan
[ ] Receive stock berjalan
[ ] Stock query berjalan
[ ] Stock transaction log berjalan
[ ] Error handling berjalan
[ ] OpenAPI tersedia
[ ] Unit tests tersedia
[ ] Integration tests tersedia
[ ] Responsive UI
[ ] AI declaration tersedia
[ ] Fresh Docker setup berhasil
```

---

# 20. MVP Completion

MVP dianggap selesai apabila seluruh P0 berikut terpenuhi:

```text
Foundation
     ✓
Database
     ✓
Item CRUD
     ✓
Location
     ✓
Stock Receive
     ✓
Stock Query
     ✓
API Infrastructure
     ✓
OpenAPI
     ✓
Frontend Item Management
     ✓
Docker
     ✓
Documentation
     ✓
Final QA
     ✓
```

**Target akhir:**

```bash
docker compose up --build
```

Kemudian user dapat:

```text
Create Item
     ↓
View Items
     ↓
Search / Filter
     ↓
Edit Item
     ↓
Soft Delete Item
     ↓
Receive Stock
     ↓
View Stock by Location
```
