package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type DB interface {
	InitDB()
	CloseDB()
	ExecQuery(ctx context.Context, query string, args ...interface{}) error
}

type Postgres struct {
	conn *pgxpool.Pool
}

func (p *Postgres) InitDB() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Can't connect to database: %v", err)
		os.Exit(1)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS tasks (id SERIAL NOT NULL PRIMARY KEY, title TEXT NOT NULL, description TEXT, status TEXT CHECK('new', 'in_progress', 'done') DEFAULT 'new', created_at TIMESTAMP DEFAULT now(), updated_at TIMESTAMP DEFAULT now())`
	_, err = conn.Exec(context.Background(), sqlStmt)
	if err != nil {
		log.Fatalf("Error executing SQL statement: %v", err)
	}
}

func (p *Postgres) CloseDB() {
	if p.conn != nil {
		p.conn.Close()
	}
}
func (p *Postgres) ExecQuery(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.conn.Exec(ctx, query, args...)
	return err
}
