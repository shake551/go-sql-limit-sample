package usecase

import (
	"context"

	"github.com/shake551/go-sql-limit-sample/domain/model"
	"github.com/shake551/go-sql-limit-sample/domain/repository"
)

type ArticleUseCase interface {
	GetLimit(ctx context.Context, limit int64, offset int64) ([]model.Article, error)
}

type articleUseCase struct {
	repository.ArticleRepository
}

func NewArticleUseCase(r repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{r}
}

func (u *articleUseCase) GetLimit(ctx context.Context, limit int64, offset int64) ([]model.Article, error) {
	return u.ArticleRepository.GetLimit(ctx, limit, offset)
}
