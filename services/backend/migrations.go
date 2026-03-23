package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// runMigrations applies all pending SQL migration files from the migrations/ directory.
// Files must be named NNN_description.sql (e.g., 001_init.sql).
// Each migration is applied in a transaction and tracked in the schema_migrations table.
func runMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	// Ensure migration tracking table exists
	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			filename TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("create schema_migrations table: %w", err)
	}

	// Read migration files
	entries, err := os.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations directory: %w", err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	// Apply each migration that hasn't been applied yet
	for _, filename := range files {
		var applied bool
		err := pool.QueryRow(ctx,
			"SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE filename = $1)",
			filename,
		).Scan(&applied)
		if err != nil {
			return fmt.Errorf("check migration %s: %w", filename, err)
		}

		if applied {
			continue
		}

		content, err := os.ReadFile(filepath.Join("migrations", filename))
		if err != nil {
			return fmt.Errorf("read migration %s: %w", filename, err)
		}

		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("begin transaction for %s: %w", filename, err)
		}

		if _, err := tx.Exec(ctx, string(content)); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("apply migration %s: %w", filename, err)
		}

		if _, err := tx.Exec(ctx,
			"INSERT INTO schema_migrations (filename) VALUES ($1)",
			filename,
		); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("record migration %s: %w", filename, err)
		}

		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("commit migration %s: %w", filename, err)
		}

		slog.Info("applied migration", "file", filename)
	}

	return nil
}
