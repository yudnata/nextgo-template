package feature2

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, d *Data) error {
	// Contoh query boilerplate
	_, err := r.db.Exec(ctx, `INSERT INTO feature2 (id, title, content) VALUES ($1, $2, $3)`, d.ID, d.Title, d.Content)
	return err
}
