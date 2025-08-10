package usecase

import (
	"context"

	"github.com/davidyunus/article-service/domain"
)

type articleUsecase struct {
	repo domain.ArticleRepository
}

func NewArticleUsecase(repo domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{repo: repo}
}

func (u *articleUsecase) CreateArticle(ctx context.Context, req *domain.CreateArticle) (*domain.Article, error) {
	a := &domain.Article{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
	}
	return u.repo.Create(ctx, a)
}

func (u *articleUsecase) ListArticles(ctx context.Context, limit, offset int) ([]*domain.Article, error) {
	return u.repo.List(ctx, limit, offset)
}
