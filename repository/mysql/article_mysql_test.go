package mysql

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/davidyunus/article-service/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestArticleMySQLRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewArticleMySQLRepo(db)

	input := &domain.Article{
		Title:   "T",
		Content: "C",
		Author:  "A",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO articles (title, content, author, created_at) VALUES (?, ?, ?, ?)`)).
		WithArgs(input.Title, input.Content, input.Author, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(123, 1))

	ctx := context.Background()
	got, err := repo.Create(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, int64(123), got.ID)
	assert.False(t, got.CreatedAt.IsZero())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestArticleMySQLRepo_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewArticleMySQLRepo(db)

	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at"}).
		AddRow(int64(1), "T1", "C1", "A1", now).
		AddRow(int64(2), "T2", "C2", "A2", now)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, content, author, created_at FROM articles ORDER BY created_at DESC LIMIT ? OFFSET ?`)).
		WithArgs(10, 0).
		WillReturnRows(rows)

	ctx := context.Background()
	res, err := repo.List(ctx, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, res, 2)
	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, "T2", res[1].Title) // Note ordering depends on returned rows

	assert.NoError(t, mock.ExpectationsWereMet())
}
