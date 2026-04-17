# OVC Management Information System — Project Specification

## 1. Project Overview

**Project Name:** Orphan and Vulnerable Children Management Information System (OVC-MIS)

**Purpose:** A web-based management information system for Helen Kaunda Secondary School to replace the current manual (paper-based) record-keeping of the Orphan and Vulnerable Children grant support program. The system enables school administrators to register, track, and report on children receiving OVC support — improving data integrity, transparency, and accountability.

**Original Project Author:** Teddy Sichivula — Copperbelt University, Zambia

**Tech Stack (Reimplementation):**

| Layer    | Technology                        |
| -------- | --------------------------------- |
| Frontend | Vue 3 + Vite + Pinia + Vue Router |
| Backend  | Go (Golang) with Gin/Chi router   |
| Database | PostgreSQL                        |
| Auth     | JWT (email + password)            |

---

## 2. Problem Statement

Helen Kaunda Secondary School currently manages the OVC grant program manually using paper records stored in files and cupboards. This system suffers from:

- Risk of damage to records (fire, water, rodents)
- Misplacement and loss of records
- Difficulty searching, reproducing, or modifying data
- Unauthorized access to sensitive information
- Lack of transparency — records can be altered to benefit ineligible children
- High cost of maintaining individual physical files per child
- No reporting capability for program evaluation

The OVC-MIS replaces this with a secure, digital system that provides CRUD operations on children records, report generation, and role-based access.

---

## 3. User Roles

### 3.1 Administrator (Database Admin)

- Creates and manages user accounts
- Full system access
- Manages system configuration

### 3.2 User (School Guidance Staff)

- Logs into the system with email/password
- Manages children records (add, edit, delete, search)
- Creates, edits, deletes, filters, and prints reports
- Views dashboard and program highlights

### 3.3 Government Inspector (Future/Read-Only)

- Logs in with credentials
- Views children records and reports (read-only)
- Monitors fund utilization and program compliance

---

## 4. Functional Requirements

### 4.1 Authentication & Authorization

| ID     | Requirement                                              | Priority  |
| ------ | -------------------------------------------------------- | --------- |
| AUTH-1 | Users must log in with email and password                | Must Have |
| AUTH-2 | Display error when user submits empty or invalid input   | Must Have |
| AUTH-3 | Only authenticated users can access the system           | Must Have |
| AUTH-4 | Auto-logout when session expires or connection is lost   | Must Have |
| AUTH-5 | Password recovery via email                              | Should Have |
| AUTH-6 | Admin can create/manage user accounts                    | Must Have |
| AUTH-7 | Student user login (view-only access)                    | Won't Have (future) |

### 4.2 Children Management (CRUD)

| ID      | Requirement                                                        | Priority    |
| ------- | ------------------------------------------------------------------ | ----------- |
| CHILD-1 | Add a new child record with full details                           | Must Have   |
| CHILD-2 | Edit an existing child record                                      | Must Have   |
| CHILD-3 | Delete a child record (with confirmation)                          | Must Have   |
| CHILD-4 | View list of all children in a data table                          | Must Have   |
| CHILD-5 | Search for a child by Pupil ID                                     | Must Have   |
| CHILD-6 | Upload and display a child's photo                                 | Should Have |
| CHILD-7 | Assign OVC category to a child                                     | Must Have   |
| CHILD-8 | Assign school requisites (with quantities) to a child              | Must Have   |
| CHILD-9 | Assign sponsor(s) to a child                                       | Must Have   |

#### 4.2.1 Child Information Fields

```
Child Information:
  - first_name        (string, required)
  - last_name         (string, required)
  - address           (string, required)
  - class_name        (string, required, e.g. "12B")
  - image_url         (string, optional)
  - pupil_id          (string/integer, unique identifier)

Guardian Information:
  - guardian_first_name   (string, required)
  - guardian_last_name    (string, required)
  - guardian_address      (string, required)
  - guardian_phone        (string, required)
```

#### 4.2.2 OVC Categories (Multi-Select — one or more per child)

```
- SINGLE ORPHAN
- DOUBLE ORPHAN
- EXTENDED FAMILY UNABLE TO MEET SCHOOL COST
- CHILD HEADED HOUSEHOLD
- ELDERLY PERSON HEADED HOUSEHOLDS AGED 65
- PARENT WITH NO RELIABLE SOURCE OF INCOME
- ON SCHOOL RE-ENTRY PROGRAM WITHOUT FAMILY SUPPORT
- POOR AS DEFINED BY CWAC
```

#### 4.2.3 School Requisites (Checklist with quantities)

Each requisite has a `checked` boolean and a `quantity` integer (adjustable with +/- controls):

```
- SCHOOL UNIFORM
- SCHOOL BAGS
- SANITARY TOWELS
- NOTE BOOKS
- PENS AND PENCILS
- MATHEMATICAL SETS
- CALCULATORS
- TEXT BOOKS
```

Additional requisites may include: shoes, stipends, pamphlets, and feeding programs.

Each requisite also has a `price_per_item` field for cost tracking.

#### 4.2.4 Sponsors (Multi-Select)

```
- GOVERNMENT REPUBLIC OF ZAMBIA (GRZ)
- NON GOVERNMENTAL ORGANIZATION (NGO)
- OTHER SPONSORS
```

### 4.3 Reports Module

| ID    | Requirement                                        | Priority    |
| ----- | -------------------------------------------------- | ----------- |
| RPT-1 | Create a new report (title + body)                 | Must Have   |
| RPT-2 | Edit an existing report                            | Must Have   |
| RPT-3 | Delete a report                                    | Must Have   |
| RPT-4 | List all reports                                   | Must Have   |
| RPT-5 | Filter reports by Term (1, 2, 3) and Year          | Must Have   |
| RPT-6 | Print/download a report as PDF                     | Must Have   |
| RPT-7 | Reports include author name and signature field     | Should Have |

#### 4.3.1 Report Fields

```
- title        (string, required)
- body         (text/rich text, required)
- term         (enum: TERM_1, TERM_2, TERM_3)
- year         (integer, e.g. 2025)
- created_by   (foreign key to users)
- created_at   (timestamp)
- updated_at   (timestamp)
```

### 4.4 Dashboard / Home Page

| ID     | Requirement                                                  | Priority    |
| ------ | ------------------------------------------------------------ | ----------- |
| DSH-1  | Display welcome message and mission statement                | Must Have   |
| DSH-2  | Show "Highlights of the Term" (images/carousel)              | Should Have |
| DSH-3  | Display summary statistics (total children, by category)     | Should Have |

### 4.5 About Page

| ID    | Requirement                                    | Priority    |
| ----- | ---------------------------------------------- | ----------- |
| ABT-1 | Display system information and version          | Could Have  |
| ABT-2 | Display developer/school contact information    | Could Have  |

---

## 5. Non-Functional Requirements

| ID     | Requirement                                                                                 |
| ------ | ------------------------------------------------------------------------------------------- |
| NFR-1  | **Security:** Sensitive child data must not be leaked or shared without authorization        |
| NFR-2  | **Reliability:** System must be readily available when required by end users                 |
| NFR-3  | **Performance:** Data fetching (children, reports) should complete in under 2 seconds        |
| NFR-4  | **Usability:** Simple, intuitive UI that non-technical school staff can operate              |
| NFR-5  | **Data Integrity:** Prevent unauthorized modification of records                             |
| NFR-6  | **Privacy:** Only authorized users can view child and guardian personal data                 |
| NFR-7  | **Maintainability:** Code should follow MVVM-like separation (API / State / UI)             |

---

## 6. System Requirements

**Minimum client requirements:**

- Modern web browser (Chrome, Firefox, Edge)
- Internet connection
- Screen resolution: 1280x720 or higher

**Server requirements:**

- Linux server (Ubuntu 22+ recommended)
- Go 1.21+
- PostgreSQL 15+
- Node.js 18+ (for building Vue frontend)
- 2 GB RAM minimum
- 10 GB storage minimum

---

## 7. Database Schema

### 7.1 Entity Relationship Summary

```
users  ──< children
children >──< ovc_categories      (via child_categories)
children >──< requisites          (via child_requisites)
children >──< sponsors            (via child_sponsors)
users  ──< reports
```

### 7.2 Table Definitions

#### `users`

```sql
CREATE TABLE users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name     VARCHAR(255) NOT NULL,
    role          VARCHAR(50) DEFAULT 'user',  -- 'admin', 'user', 'inspector'
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW()
);
```

#### `children`

```sql
CREATE TABLE children (
    id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pupil_id              VARCHAR(50) UNIQUE NOT NULL,
    first_name            VARCHAR(100) NOT NULL,
    last_name             VARCHAR(100) NOT NULL,
    address               VARCHAR(255) NOT NULL,
    class_name            VARCHAR(20) NOT NULL,
    image_url             VARCHAR(500),
    guardian_first_name   VARCHAR(100) NOT NULL,
    guardian_last_name    VARCHAR(100) NOT NULL,
    guardian_address      VARCHAR(255) NOT NULL,
    guardian_phone        VARCHAR(20) NOT NULL,
    created_by            UUID REFERENCES users(id),
    created_at            TIMESTAMP DEFAULT NOW(),
    updated_at            TIMESTAMP DEFAULT NOW()
);
```

#### `ovc_categories`

```sql
CREATE TABLE ovc_categories (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100) UNIQUE NOT NULL
);

-- Seed data:
-- SINGLE ORPHAN, DOUBLE ORPHAN, EXTENDED FAMILY UNABLE TO MEET SCHOOL COST,
-- CHILD HEADED HOUSEHOLD, ELDERLY PERSON HEADED HOUSEHOLDS AGED 65,
-- PARENT WITH NO RELIABLE SOURCE OF INCOME,
-- ON SCHOOL RE-ENTRY PROGRAM WITHOUT FAMILY SUPPORT,
-- POOR AS DEFINED BY CWAC
```

#### `child_categories`

```sql
CREATE TABLE child_categories (
    child_id    UUID REFERENCES children(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES ovc_categories(id) ON DELETE CASCADE,
    PRIMARY KEY (child_id, category_id)
);
```

#### `requisites`

```sql
CREATE TABLE requisites (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) UNIQUE NOT NULL,
    default_price   DECIMAL(10,2) DEFAULT 0.00
);

-- Seed data:
-- SCHOOL UNIFORM, SCHOOL BAGS, SANITARY TOWELS, NOTE BOOKS,
-- PENS AND PENCILS, MATHEMATICAL SETS, CALCULATORS, TEXT BOOKS
```

#### `child_requisites`

```sql
CREATE TABLE child_requisites (
    child_id      UUID REFERENCES children(id) ON DELETE CASCADE,
    requisite_id  INTEGER REFERENCES requisites(id) ON DELETE CASCADE,
    quantity      INTEGER DEFAULT 1,
    checked       BOOLEAN DEFAULT false,
    price_per_item DECIMAL(10,2) DEFAULT 0.00,
    PRIMARY KEY (child_id, requisite_id)
);
```

#### `sponsors`

```sql
CREATE TABLE sponsors (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100) UNIQUE NOT NULL
);

-- Seed data:
-- GOVERNMENT REPUBLIC OF ZAMBIA, NON GOVERNMENTAL ORGANIZATION, OTHER SPONSORS
```

#### `child_sponsors`

```sql
CREATE TABLE child_sponsors (
    child_id    UUID REFERENCES children(id) ON DELETE CASCADE,
    sponsor_id  INTEGER REFERENCES sponsors(id) ON DELETE CASCADE,
    PRIMARY KEY (child_id, sponsor_id)
);
```

#### `reports`

```sql
CREATE TABLE reports (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title       VARCHAR(255) NOT NULL,
    body        TEXT NOT NULL,
    term        VARCHAR(10) NOT NULL,  -- 'TERM_1', 'TERM_2', 'TERM_3'
    year        INTEGER NOT NULL,
    created_by  UUID REFERENCES users(id),
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);
```

#### `highlights` (optional — for home page images)

```sql
CREATE TABLE highlights (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_url   VARCHAR(500) NOT NULL,
    caption     VARCHAR(255),
    term        VARCHAR(10),
    year        INTEGER,
    created_at  TIMESTAMP DEFAULT NOW()
);
```

---

## 8. API Endpoints

### 8.1 Authentication

```
POST   /api/auth/register       — Create a new user (admin only)
POST   /api/auth/login          — Login, returns JWT token
POST   /api/auth/forgot-password — Send password reset email
POST   /api/auth/reset-password  — Reset password with token
GET    /api/auth/me             — Get current user profile
```

### 8.2 Children

```
GET    /api/children             — List all children (supports ?search=pupil_id)
POST   /api/children             — Create a new child record
GET    /api/children/:id         — Get a single child with categories, requisites, sponsors
PUT    /api/children/:id         — Update a child record
DELETE /api/children/:id         — Delete a child record
POST   /api/children/:id/image   — Upload child photo
```

### 8.3 Child Relationships

```
PUT    /api/children/:id/categories   — Set OVC categories for a child
PUT    /api/children/:id/requisites   — Set requisites for a child
PUT    /api/children/:id/sponsors     — Set sponsors for a child
```

### 8.4 Lookup Data

```
GET    /api/categories       — List all OVC categories
GET    /api/requisites       — List all requisites
GET    /api/sponsors         — List all sponsors
```

### 8.5 Reports

```
GET    /api/reports           — List reports (supports ?term=TERM_1&year=2025)
POST   /api/reports           — Create a new report
GET    /api/reports/:id       — Get a single report
PUT    /api/reports/:id       — Update a report
DELETE /api/reports/:id       — Delete a report
GET    /api/reports/:id/pdf   — Download report as PDF
```

### 8.6 Dashboard

```
GET    /api/dashboard/stats       — Summary counts (children, by category, by sponsor)
GET    /api/highlights            — Get highlight images for home page
POST   /api/highlights            — Upload a highlight image (admin only)
DELETE /api/highlights/:id        — Remove a highlight (admin only)
```

---

## 9. Frontend Pages & Navigation

### 9.1 Sidebar Navigation

```
├── Home         (dashboard, mission statement, highlights)
├── Children     (data table, CRUD operations)
├── Reports      (list, create, filter, print)
├── About        (system info)
└── Log Out
```

### 9.2 Page Breakdown

| Page              | Route              | Description                                              |
| ----------------- | ------------------ | -------------------------------------------------------- |
| Login             | `/login`           | Email + password form with validation errors             |
| Home / Dashboard  | `/`                | Welcome message, mission, highlights carousel, stats     |
| Children List     | `/children`        | Data table with search bar, "Add New Child" button       |
| Add/Edit Child    | `/children/new`    | Multi-step form: Info → Requisites → Category → Sponsors |
| Child Detail      | `/children/:id`    | View full child record with all relationships            |
| Reports List      | `/reports`         | List of reports with filter (term + year) and print      |
| Add/Edit Report   | `/reports/new`     | Report title + body form                                 |
| About             | `/about`           | System and contact information                           |

### 9.3 Key UI Components

- **Sidebar:** Persistent left sidebar with icon + label navigation
- **Data Table:** Sortable, searchable table for children (columns: Photo, ID, First Name, Last Name, Sponsor Name, Category, Address)
- **OVC Category Dialog:** Modal with checkboxes for each category
- **Requisites Dialog:** Modal with item name, quantity (+/- controls), and checkbox per item
- **Sponsors Dialog:** Modal with checkboxes for each sponsor type
- **Filter Reports Dialog:** Modal with Term radio buttons (1/2/3) and Year dropdown
- **Error Dialogs:** Toast/snackbar notifications for validation errors
- **Confirmation Dialog:** "Are you sure?" prompt before deleting records

---

## 10. Build Phases

### Phase 1 — Foundation & Auth (Week 1–2)

- [ ] Initialize Go project with router and PostgreSQL connection
- [ ] Create database migrations for `users` table
- [ ] Build register/login endpoints with JWT
- [ ] Scaffold Vue 3 project with Vite, Pinia, Vue Router
- [ ] Build login page with validation
- [ ] Implement token storage and route guards

### Phase 2 — Children CRUD (Week 3–4)

- [ ] Create `children` table migration
- [ ] Build children API endpoints (list, create, read, update, delete)
- [ ] Build Children list page with data table and search
- [ ] Build Add/Edit Child form (child info + guardian info)
- [ ] Implement image upload for child photos

### Phase 3 — Categories, Requisites & Sponsors (Week 5–6)

- [ ] Create lookup tables and seed data migrations
- [ ] Create junction table migrations
- [ ] Build relationship API endpoints
- [ ] Build OVC Category selection modal
- [ ] Build Requisites selection modal with quantity controls
- [ ] Build Sponsors selection modal
- [ ] Integrate modals into child creation/edit flow

### Phase 4 — Reports Module (Week 7–8)

- [ ] Create `reports` table migration
- [ ] Build reports API endpoints with term/year filtering
- [ ] Build PDF generation endpoint
- [ ] Build Reports list page
- [ ] Build report create/edit form
- [ ] Build filter dialog (term + year)
- [ ] Implement print/download PDF functionality

### Phase 5 — Dashboard & Polish (Week 9–10)

- [ ] Build dashboard stats endpoint
- [ ] Build Home page with mission statement and stats
- [ ] Build highlights carousel
- [ ] Build About page
- [ ] Add form validation across all forms
- [ ] Add error/success toast notifications
- [ ] Implement auto-logout on session expiry

### Phase 6 — Testing & Deployment (Week 11–12)

- [ ] Write Go unit tests for auth and CRUD handlers
- [ ] Test all frontend flows end-to-end
- [ ] Set up CORS and security headers
- [ ] Dockerize backend + database
- [ ] Build and serve Vue frontend via Nginx
- [ ] Write deployment documentation

---

## 11. OVC Grant Program Context

The OVC grant program at Helen Kaunda Secondary School operates three times per year (once per term). Children are evaluated and categorized by vulnerability level. Based on their category, they receive various supplies funded by the Government of Zambia (GRZ), NGOs, or individual sponsors.

The system must track:

- **Who** received support (child identity and guardian)
- **What** they received (requisites with quantities)
- **Why** they qualify (OVC category)
- **Who funded** the support (sponsors)
- **When** it happened (term and year via reports)

This data enables the school to make evidence-based decisions, allocate resources fairly, and demonstrate accountability to government inspectors and donors.

---

## 12. Future Enhancements (Won't Have — Current Scope)

- Student login portal (view-only access to their own records)
- Health and mental welfare tracking module
- Public-facing webpage showcasing program impact
- Mobile app version (Android/iOS)
- Multi-school support (deploy for multiple schools)
- Bulk import of children data (CSV/Excel)
- Audit trail / activity logging
- SMS notifications to guardians
