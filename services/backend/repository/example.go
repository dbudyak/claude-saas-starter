package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yourorg/yourapp/models"
)

var ErrNotFound = errors.New("not found")

// ExampleRepository shows the repository pattern. Replace with your domain repositories.
type ExampleRepository struct {
	pool *pgxpool.Pool
}

func NewExampleRepository(pool *pgxpool.Pool) *ExampleRepository {
	return &ExampleRepository{pool: pool}
}

func (r *ExampleRepository) GetByID(ctx context.Context, id string) (*models.Example, error) {
	query := `SELECT id, status, created_at FROM examples WHERE id = $1`

	var e models.Example
	err := r.pool.QueryRow(ctx, query, id).Scan(&e.ID, &e.Status, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get example: %w", err)
	}

	return &e, nil
}
