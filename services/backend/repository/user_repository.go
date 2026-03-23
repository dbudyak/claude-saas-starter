package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yourorg/yourapp/models"
)

// ErrNotFound is returned when a requested resource does not exist.
var ErrNotFound = errors.New("not found")

// UserRepository handles database operations for users.
type UserRepository struct {
	pool *pgxpool.Pool
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

// Create inserts a new user and returns the created record.
func (r *UserRepository) Create(ctx context.Context, email, passwordHash string) (*models.User, error) {
	query := `
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, $3)
		RETURNING id, email, password_hash, role, created_at
	`

	var u models.User
	var role string
	err := r.pool.QueryRow(ctx, query, email, passwordHash, string(models.UserRoleMember)).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &role, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	u.Role = models.UserRole(role)
	return &u, nil
}

// FindByEmail returns the user with the given email, or ErrNotFound.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash, role, created_at FROM users WHERE email = $1`

	var u models.User
	var role string
	err := r.pool.QueryRow(ctx, query, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &role, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	u.Role = models.UserRole(role)
	return &u, nil
}

// FindByID returns the user with the given ID, or ErrNotFound.
func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, email, password_hash, role, created_at FROM users WHERE id = $1`

	var u models.User
	var role string
	err := r.pool.QueryRow(ctx, query, id).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &role, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	u.Role = models.UserRole(role)
	return &u, nil
}

// EmailExists returns true if a user with the given email already exists.
func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		email,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check email existence: %w", err)
	}
	return exists, nil
}
