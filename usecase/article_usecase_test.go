package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/davidyunus/article-service/domain"

	"github.com/stretchr/testify/assert"
)

// simple stub repo to test usecase behavior
type stubRepo struct {
	createFn func(ctx context.Context, a *domain.Article) (*domain.Article, error)
	listFn   func(ctx context.Context, limit, offset int) ([]*domain.Article, error)
}

func (s *stubRepo) Create(ctx context.Context, a *domain.Article) (*domain.Article, error) {
	return s.createFn(ctx, a)
}
func (s *stubRepo) List(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
	return s.listFn(ctx, limit, offset)
}

func TestCreateArticle_Success(t *testing.T) {
	now := time.Now()
	stub := &stubRepo{
		createFn: func(ctx context.Context, a *domain.Article) (*domain.Article, error) {
			a.ID = 11
			a.CreatedAt = now
			return a, nil
		},
	}
	uc := NewArticleUsecase(stub)

	req := &domain.CreateArticle{
		Title:   "title",
		Content: "content",
		Author:  "author",
	}
	got, err := uc.CreateArticle(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int64(11), got.ID)
	assert.Equal(t, "title", got.Title)
}

func TestCreateArticle_Error(t *testing.T) {
	stub := &stubRepo{
		createFn: func(ctx context.Context, a *domain.Article) (*domain.Article, error) {
			return nil, errors.New("fail")
		},
	}
	uc := NewArticleUsecase(stub)

	req := &domain.CreateArticle{Title: "x"}
	_, err := uc.CreateArticle(context.Background(), req)
	assert.Error(t, err)
}

func TestListArticles(t *testing.T) {
	stub := &stubRepo{
		listFn: func(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
			return []*domain.Article{
				{ID: 1, Title: "t1"},
				{ID: 2, Title: "t2"},
			}, nil
		},
	}
	uc := NewArticleUsecase(stub)
	res, err := uc.ListArticles(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, res, 2)
}
