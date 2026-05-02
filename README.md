# Go + Vue + PostgreSQL Template

A full-stack starter template using **Go (Gin)** for the backend API, **Vue 3 (Vite + Tailwind)** for the frontend, and **PostgreSQL** as the database.

---

## Stack

| Layer    | Technology                              |
|----------|-----------------------------------------|
| Backend  | Go 1.22, Gin, GORM, JWT                 |
| Frontend | Vue 3, Vite, Pinia, Vue Router, Tailwind CSS |
| Database | PostgreSQL                              |

---

## Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) >= 1.22
- [Node.js](https://nodejs.org/) >= 20.19
- [PostgreSQL](https://www.postgresql.org/) >= 14
- `make` (optional but recommended)

---

## 1. Clone the repo

```bash
git clone git@github.com:Gymnott1/go_vue__postgres_tamplete.git
cd go_vue__postgres_tamplete
```

---

## 2. Set up PostgreSQL

### Create the database

Connect as the postgres superuser and create the database:

```bash
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'yourpassword';"
sudo -u postgres psql -c "CREATE DATABASE clientcompany;"
```

### (Optional) Save credentials to avoid password prompts

```bash
echo "localhost:5432:clientcompany:postgres:yourpassword" > ~/.pgpass
chmod 600 ~/.pgpass
```

---

## 3. Configure environment variables

Copy the example env and fill in your values:

```bash
cp .env.example .env
```

Edit `.env`:

```env
APP_PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=clientcompany

JWT_SECRET=change-me-in-production
SESSION_SECRET=change-me-in-production
```

> Never commit `.env` to version control. It is already in `.gitignore`.

---

## 4. Install Go dependencies

```bash
go mod download
```

> If you are behind a restricted network, set the Go proxy:
> ```bash
> export GOPROXY=https://goproxy.io,direct
> ```

---

## 5. Run database migrations + seed data

The app auto-migrates on startup. To seed the default users:

```bash
go run cmd/seed/main.go
```

This creates:

| Name       | Email                    | Password    | Role  |
|------------|--------------------------|-------------|-------|
| Admin User | admin@clientcompany.com  | Admin@1234  | admin |
| Jane Smith | jane@clientcompany.com   | User@1234   | user  |
| John Doe   | john@clientcompany.com   | User@1234   | user  |

The seeder is idempotent — safe to run multiple times.

---

## 6. Install frontend dependencies

```bash
cd ui && npm install && cd ..
```

---

## 7. Build the Vue frontend

```bash
make ui
# or manually:
cd ui && npm run build && cd ..
```

This builds Vue into `dist/` at the project root. Go serves it automatically on startup.

---

## 8. Run the server

```bash
make run
# or:
go run main.go
```

Open **http://localhost:8080** in your browser.

---

## Development mode (hot reload)

Run Go and the Vite dev server concurrently:

```bash
make dev
```

| Service    | URL                      |
|------------|--------------------------|
| Go API     | http://localhost:8080    |
| Vite (Vue) | http://localhost:5173    |

In dev mode, Vite proxies all `/api` requests to the Go server automatically.

---

## Project structure

```
.
├── main.go                   # Entry point
├── .env                      # Environment variables (not committed)
├── Makefile                  # Dev/build shortcuts
├── config/config.go          # Loads config from .env
├── db/db.go                  # PostgreSQL connection + auto-migrate
├── models/user.go            # GORM models
├── middleware/
│   ├── auth.go               # JWT auth + role guard middleware
│   └── session.go            # Request logger
├── handlers/auth.go          # Register, login, logout, me
├── api/
│   ├── apiinternal/routes.go # Internal API route groups
│   └── external/client.go   # Reusable external HTTP client
├── routes/routes.go          # Wires all routes together
├── templates/                # Fallback HTML templates (no Vue dist)
├── static/                   # Fallback CSS/JS (no Vue dist)
├── cmd/seed/main.go          # Database seeder
└── ui/                       # Vue 3 frontend
    ├── src/
    │   ├── api/index.js      # Fetch wrapper for Go API
    │   ├── stores/auth.js    # Pinia auth store
    │   ├── router/index.js   # Vue Router + auth guards
    │   └── views/
    │       ├── LoginView.vue
    │       └── DashboardView.vue
    └── vite.config.js        # Builds to ../dist
```

---

## API endpoints

| Method | Path                    | Auth     | Description        |
|--------|-------------------------|----------|--------------------|
| POST   | /api/v1/auth/register   | No       | Register new user  |
| POST   | /api/v1/auth/login      | No       | Login, returns JWT |
| POST   | /api/v1/auth/logout     | No       | Clear token cookie |
| GET    | /api/v1/me              | Required | Current user info  |
| GET    | /api/v1/admin/users     | Admin    | List all users     |

---

## Makefile commands

| Command      | Description                              |
|--------------|------------------------------------------|
| `make run`   | Run the Go server                        |
| `make ui`    | Build Vue into dist/                     |
| `make build` | Build Vue then compile Go binary         |
| `make dev`   | Run Go + Vite dev server concurrently    |
# KFS
