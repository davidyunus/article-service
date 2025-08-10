package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidyunus/article-service/domain"
)

type articleMySQLRepo struct {
	db *sql.DB
}

func NewArticleMySQLRepo(db *sql.DB) domain.ArticleRepository {
	return &articleMySQLRepo{db: db}
}

func (r *articleMySQLRepo) Create(ctx context.Context, a *domain.Article) (*domain.Article, error) {
	now := time.Now().UTC()
	query := `INSERT INTO articles (title, content, author, created_at) VALUES (?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, a.Title, a.Content, a.Author, now)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	a.ID = id
	a.CreatedAt = now
	return a, nil
}

func (r *articleMySQLRepo) List(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
	query := `SELECT id, title, content, author, created_at FROM articles ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*domain.Article
	for rows.Next() {
		var a domain.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Author, &a.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
