package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func NewPostgresDB(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, nil
}
