# Article Service

A simple clean-architecture-based Go service for managing articles.  
Uses [Echo](https://echo.labstack.com/) as the HTTP framework and PostgreSQL as the database.

---

## 📂 Project Structure

```
.
├── cmd/
│   └── server/              # Main application entrypoint
├── domain/                  # Entities and interfaces
├── repository/              # Database layer
├── usecase/                 # Business logic
├── handler/                 # HTTP handlers
├── schema.sql               # Database schema
├── Dockerfile               # Container build file
├── docker-compose.yml       # Local dev with Postgres
└── README.md
```

---

## 🚀 How to Run Locally

### 1. Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (if not using Docker)

---

### 2. Clone the Repository

```bash
git clone https://github.com/yourusername/article-service.git
cd article-service
```

---

### 3. Set up Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/articles?sslmode=disable
PORT=8080
```

---

### 4. Run with Docker Compose

We provide a `docker-compose.yml` that starts:
- **PostgreSQL** database with schema initialization.
- **Go API service** that connects to the database.

Run:

```bash
docker compose up --build
```

Once running:
- API available at: `http://localhost:8080`
- PostgreSQL available at: `localhost:5432`

---

### 5. Run without Docker

1. Create database & run schema:
```bash
createdb articles
psql -d articles -f schema.sql
```

2. Run API:
```bash
go run cmd/server/main.go
```

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 📦 Example API Endpoints

### Create an Article
```bash
curl -X POST http://localhost:8080/articles   -H "Content-Type: application/json"   -d '{
    "title": "My First Article",
    "content": "Hello, world!",
    "author": "David Yunus"
  }'
```

### List Articles
```bash
curl -X GET "http://localhost:8080/articles?author=John&page=1&limit=10"
```

---

## 🛠 Tech Stack

- **Language**: Go
- **Framework**: Echo
- **Database**: PostgreSQL
- **Architecture**: Clean Architecture (Domain, Usecase, Repository, Handler)
- **Testing**: `testing` + `testify`

---

## 🐳 Docker Setup

- **Dockerfile**: Builds a minimal Go binary and runs the service.
- **docker-compose.yml**: Defines the database service, mounts schema, and connects it to the API container.

Run `docker compose up --build` to start the stack.