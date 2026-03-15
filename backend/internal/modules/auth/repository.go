package auth

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

func (r *Repository) Create(ctx context.Context, u *User) error {
	_, err := r.db.Exec(ctx, `INSERT INTO users (id, email, name, password) VALUES ($1, $2, $3, $4)`, u.ID, u.Email, u.Name, u.Password)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := r.db.QueryRow(ctx, `SELECT id, email, name, password FROM users WHERE email = $1`, email).Scan(&u.ID, &u.Email, &u.Name, &u.Password)
	return &u, err
}
