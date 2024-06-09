package provider

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type Category struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}

type CategoryProvider struct {
	db *sqlx.DB
}

func NewCategoryProvider(db *sqlx.DB) *CategoryProvider {
	return &CategoryProvider{
		db: db,
	}
}

func (cp *CategoryProvider) Store(ctx context.Context, name string) (uint64, error) {
	var id uint64
	err := cp.db.QueryRowContext(ctx, "INSERT INTO categories (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		log.Println("Error storing category:", err)
		return 0, err
	}
	return id, nil
}

func (cp *CategoryProvider) GetByName(ctx context.Context, name string) (*Category, error) {
	var category Category
	err := cp.db.GetContext(ctx, &category, "SELECT id, name FROM categories WHERE name = $1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		log.Println("Error getting category by name:", err)
		return nil, err
	}
	return &category, nil
}
