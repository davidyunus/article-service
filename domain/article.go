package domain

import (
	"context"
	"time"
)

// Use int64 for DB compatibility with LastInsertId
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Repository interface (for DI)
type ArticleRepository interface {
	Create(ctx context.Context, a *Article) (*Article, error)
	List(ctx context.Context, limit, offset int) ([]*Article, error)
}

// Usecase interface
type ArticleUsecase interface {
	CreateArticle(ctx context.Context, req *CreateArticle) (*Article, error)
	ListArticles(ctx context.Context, limit, offset int) ([]*Article, error)
}
