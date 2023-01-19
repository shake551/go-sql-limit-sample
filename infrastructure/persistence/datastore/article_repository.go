package datastore

import (
	"context"
	"time"

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

func (r ArticleRepository) Create(ctx context.Context, param model.ArticleParam) (*model.Article, error) {
	now := time.Now().Unix()
	query := `INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	res, err := db.DB.ExecContext(ctx, query, param.Title, param.Content, now, now)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	article := &model.Article{
		ID:        id,
		Title:     param.Title,
		Content:   param.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return article, nil
}
