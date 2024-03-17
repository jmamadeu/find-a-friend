package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	*pgxpool.Pool
}

func New(ctx context.Context) (*Db, error) {
	url := os.Getenv("DB_URL")
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Db{
		db,
	}, nil
}
