package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/davidyunus/article-service/domain"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type stubUC struct {
	createFn func(ctx context.Context, req *domain.CreateArticle) (*domain.Article, error)
	listFn   func(ctx context.Context, limit, offset int) ([]*domain.Article, error)
}

func (s *stubUC) CreateArticle(ctx context.Context, req *domain.CreateArticle) (*domain.Article, error) {
	return s.createFn(ctx, req)
}
func (s *stubUC) ListArticles(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
	return s.listFn(ctx, limit, offset)
}

func TestCreateArticle_Handler(t *testing.T) {
	e := echo.New()
	now := time.Now()
	stub := &stubUC{
		createFn: func(ctx context.Context, req *domain.CreateArticle) (*domain.Article, error) {
			return &domain.Article{
				ID:        10,
				Title:     req.Title,
				Content:   req.Content,
				Author:    req.Author,
				CreatedAt: now,
			}, nil
		},
	}
	h := NewArticleHandler(stub)

	body := `{"title":"t","content":"c","author":"a"}`
	req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreateArticle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var got domain.Article
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, int64(10), got.ID)
	}
}

func TestListArticles_Handler(t *testing.T) {
	e := echo.New()
	stub := &stubUC{
		listFn: func(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
			return []*domain.Article{
				{ID: 1, Title: "t1"},
				{ID: 2, Title: "t2"},
			}, nil
		},
	}
	h := NewArticleHandler(stub)

	req := httptest.NewRequest(http.MethodGet, "/articles?limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.ListArticles(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var got []*domain.Article
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Len(t, got, 2)
	}
}
