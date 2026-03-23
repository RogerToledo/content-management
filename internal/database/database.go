package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go/content-management/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgrsPool(ctx context.Context) (*pgxpool.Pool, error) {
	cfg := config.Get()

	poolConfig, err := pgxpool.ParseConfig(cfg.DbUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing Dburl config: %w", err)
	}

	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MinConns = int32(cfg.MinConns)
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating new pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging pool: %w", err)
	}

	return pool, nil

}
