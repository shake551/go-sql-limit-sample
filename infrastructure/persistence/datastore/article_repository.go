package datastore

import (
	"context"

	"github.com/shake551/go-sql-limit-sample/db"
	"github.com/shake551/go-sql-limit-sample/domain/model"
)

type ArticleRepository struct{}

func NewCocktailRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (r ArticleRepository) GetLimit(ctx context.Context, limit int64, offset int64) ([]model.Article, error) {

	query := `SELECT * FROM articles LIMIT ? OFFSET ?`
	rows, err := db.DB.QueryContext(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		a := model.Article{}
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}

		articles = append(articles, a)
	}

	if len(articles) == 0 {
		return []model.Article{}, nil
	}

	return articles, nil
}
