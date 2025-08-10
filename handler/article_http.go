package handler

import (
	"net/http"
	"strconv"

	"github.com/davidyunus/article-service/domain"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	uc domain.ArticleUsecase
}

func NewArticleHandler(uc domain.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{uc: uc}
}

func RegisterArticleRoutes(e *echo.Echo, uc domain.ArticleUsecase) {
	h := NewArticleHandler(uc)
	e.POST("/articles", h.CreateArticle)
	e.GET("/articles", h.ListArticles)
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	var req domain.CreateArticle
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	art, err := h.uc.CreateArticle(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, art)
}

func (h *ArticleHandler) ListArticles(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if limit <= 0 {
		limit = 10
	}
	articles, err := h.uc.ListArticles(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articles)
}
