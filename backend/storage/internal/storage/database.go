package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDatabase interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Close()
}

type postgresDatabase struct {
	pool *pgxpool.Pool
}

func NewPostgresDatabase(pool *pgxpool.Pool) (IDatabase, error) {
	if pool == nil {
		return nil, fmt.Errorf("pool is nil")
	}
	return &postgresDatabase{pool: pool}, nil
}

func (p *postgresDatabase) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *postgresDatabase) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, sql, args...)
}

func (p *postgresDatabase) Close() {
	p.pool.Close()
}
