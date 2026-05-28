package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)
// to add the required module to your go.mod/go.sum.

func ConnectDB() (*pgxpool.Pool, error) {
	connString := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
