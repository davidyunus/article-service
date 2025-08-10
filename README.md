# Article Service

This is a simple Article Service built using Golang with Clean Architecture (domain, usecase, repository, handler) and Echo framework.  
The repository layer uses MySQL with `database/sql`.  

## Features

- Create an article
- List all articles
- Clean Architecture design
- Unit tests for usecase, repository, and handlers
- Docker and docker-compose for deployment

## Project Structure

```
.
├── domain
│   └── article.go
├── handler
│   ├── article_handler.go
│   └── article_handler_test.go
├── repository
│   ├── article_mysql.go
│   └── article_mysql_test.go
├── usecase
│   ├── article_usecase.go
│   └── article_usecase_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

## Requirements

- Go 1.22+
- MySQL 8.0+
- Docker & Docker Compose

## How to Run Locally

1. **Clone the repository**
```
git clone https://github.com/davidyunus/article-service.git
cd article-service
```

2. **Setup environment variables**
Create a `.env` file or set directly in your terminal:
```
DB_USER=article_user
DB_PASSWORD=article_pass
DB_NAME=article_db
DB_HOST=localhost
DB_PORT=3306
```

3. **Run using Docker Compose**
```
docker-compose up --build
```

4. **Run without Docker (requires MySQL running locally)**
```
go mod tidy
go run main.go
```

5. **Access API**
- Create article: `POST http://localhost:8080/articles`
- List articles: `GET http://localhost:8080/articles`

## Running Tests

```
go test ./...
```