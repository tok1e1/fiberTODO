package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	InitDB(databaseURL string) error
	CloseDB()
	ExecQuery(ctx context.Context, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

type Postgres struct {
	conn *pgxpool.Pool
}

func (p *Postgres) InitDB(databaseURL string) error {

	conn, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("Can't connect to database: %v", err)
	}

	p.conn = conn

	sqlStmt := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY, 
		title TEXT NOT NULL, 
		description TEXT, 
		status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new', 
		created_at TIMESTAMP DEFAULT now(), 
		updated_at TIMESTAMP DEFAULT now()
	)`

	_, err = conn.Exec(context.Background(), sqlStmt)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %v", err)
	}

	return nil
}

func (p *Postgres) CloseDB() {
	if p.conn != nil {
		p.conn.Close()
	}
}

func (p *Postgres) ExecQuery(ctx context.Context, query string, args ...interface{}) error {
	if p.conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	_, err := p.conn.Exec(ctx, query, args...)
	return err
}

func (p *Postgres) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if p.conn == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	return p.conn.Query(ctx, query, args...)
}
