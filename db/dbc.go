package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type dbconn interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
}

var DB dbconn

func Connect(s string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), s)
	if err != nil {
		return nil, fmt.Errorf("не смог подключиться к БД: %w", err)
	}
	fmt.Println("успешный коннект к БД")
	return db, nil
}
