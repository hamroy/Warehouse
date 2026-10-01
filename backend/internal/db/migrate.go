// Package db provides database connection, migration, and seed helpers.
package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"regexp"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

// migrationFiles are ordered by the numeric prefix in the filename.
type migrationFile struct {
	version int
	up      string
	down    string
}

// migrationRecord tracks applied versions in the database.
const migrationTable = `
CREATE TABLE IF NOT EXISTS schema_migrations (
    version   INTEGER PRIMARY KEY,
    name      TEXT NOT NULL,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
)`

var versionRe = regexp.MustCompile(`^(\d+)_(.+)\.(up|down)\.sql$`)

// Connect opens a connection to PostgreSQL using the given DSN and pings it.
func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return db, nil
}

// loadMigrations reads embedded migration files and pairs up/down by version.
func loadMigrations() ([]migrationFile, error) {
	entries, err := fs.ReadDir(migrationFS, "migrations")
	if err != nil {
		return nil, fmt.Errorf("read embedded migrations: %w", err)
	}

	ups := map[int]string{}
	downs := map[int]string{}

	for _, e := range entries {
		m := versionRe.FindStringSubmatch(e.Name())
		if m == nil {
			continue
		}
		var version int
		if _, err := fmt.Sscanf(m[1], "%d", &version); err != nil {
			return nil, fmt.Errorf("parse migration version %q: %w", m[1], err)
		}
		content, err := fs.ReadFile(migrationFS, "migrations/"+e.Name())
		if err != nil {
			return nil, fmt.Errorf("read migration %s: %w", e.Name(), err)
		}
		switch m[3] {
		case "up":
			ups[version] = string(content)
		case "down":
			downs[version] = string(content)
		}
	}

	versions := make([]int, 0, len(ups))
	for v := range ups {
		versions = append(versions, v)
	}
	sort.Ints(versions)

	files := make([]migrationFile, 0, len(versions))
	for _, v := range versions {
		up, ok := ups[v]
		if !ok {
			continue
		}
		files = append(files, migrationFile{version: v, up: up, down: downs[v]})
	}
	return files, nil
}

// appliedVersions returns the set of migration versions already recorded.
func appliedVersions(ctx context.Context, db *sql.DB) (map[int]bool, error) {
	if _, err := db.ExecContext(ctx, migrationTable); err != nil {
		return nil, fmt.Errorf("ensure schema_migrations: %w", err)
	}
	rows, err := db.QueryContext(ctx, `SELECT version FROM schema_migrations`)
	if err != nil {
		return nil, fmt.Errorf("read schema_migrations: %w", err)
	}
	defer rows.Close()

	applied := map[int]bool{}
	for rows.Next() {
		var v int
		if err := rows.Scan(&v); err != nil {
			return nil, fmt.Errorf("scan version: %w", err)
		}
		applied[v] = true
	}
	return applied, rows.Err()
}

// Migrate applies all pending migrations in order inside a transaction.
func Migrate(ctx context.Context, db *sql.DB) error {
	files, err := loadMigrations()
	if err != nil {
		return err
	}
	applied, err := appliedVersions(ctx, db)
	if err != nil {
		return err
	}

	for _, f := range files {
		if applied[f.version] {
			continue
		}
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("begin migration %d: %w", f.version, err)
		}
		if _, err := tx.ExecContext(ctx, f.up); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %d failed: %w", f.version, err)
		}
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO schema_migrations (version, name) VALUES ($1, $2)`,
			f.version, fmt.Sprintf("%03d", f.version)); err != nil {
			tx.Rollback()
			return fmt.Errorf("record migration %d: %w", f.version, err)
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit migration %d: %w", f.version, err)
		}
		fmt.Printf("migration %03d applied\n", f.version)
	}
	return nil
}

// Rollback reverts the last applied migration.
func Rollback(ctx context.Context, db *sql.DB) error {
	files, err := loadMigrations()
	if err != nil {
		return err
	}
	applied, err := appliedVersions(ctx, db)
	if err != nil {
		return err
	}

	// find highest applied version
	last := -1
	for _, f := range files {
		if applied[f.version] && f.version > last {
			last = f.version
		}
	}
	if last < 0 {
		return errors.New("no migrations to roll back")
	}

	var target migrationFile
	for _, f := range files {
		if f.version == last {
			target = f
		}
	}
	if target.down == "" {
		return fmt.Errorf("migration %03d has no down script", last)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin rollback %d: %w", last, err)
	}
	if _, err := tx.ExecContext(ctx, target.down); err != nil {
		tx.Rollback()
		return fmt.Errorf("rollback %d failed: %w", last, err)
	}
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM schema_migrations WHERE version = $1`, last); err != nil {
		tx.Rollback()
		return fmt.Errorf("unrecord migration %d: %w", last, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit rollback %d: %w", last, err)
	}
	fmt.Printf("migration %03d rolled back\n", last)
	return nil
}

// EnsureString is a small helper to keep strings import used in future edits.
var _ = strings.TrimSpace
