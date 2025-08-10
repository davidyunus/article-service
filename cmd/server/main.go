package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/davidyunus/article-service/handler"
	repository "github.com/davidyunus/article-service/repository/mysql"
	"github.com/davidyunus/article-service/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// example: user:pass@tcp(host:3306)/dbname?parseTime=true
		dsn = "root:password@tcp(localhost:3306)/articles?parseTime=true"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	repo := repository.NewArticleMySQLRepo(db)
	uc := usecase.NewArticleUsecase(repo)

	e := echo.New()
	handler.RegisterArticleRoutes(e, uc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
