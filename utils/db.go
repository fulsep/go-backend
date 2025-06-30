package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func DBConnect() (*pgxpool.Conn, error) {
	godotenv.Load()

	// connectionString := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s?sslmode=require&channel_binding=require",
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGHOST"),
	// 	os.Getenv("PGPORT"),
	// 	os.Getenv("PGDATABASE"),
	// )
	pool, _ := pgxpool.New(
		context.Background(),
		// connectionString,
		os.Getenv("DATABASE_URL"),
	)

	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return conn, err
	}

	return conn, err

}
