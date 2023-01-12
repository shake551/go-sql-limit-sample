package repository

import (
	"context"

	"github.com/shake551/go-sql-limit-sample/domain/model"
)

type ArticleRepository interface {
	GetLimit(ctx context.Context, limit int64, offset int64) ([]model.Article, error)
}
