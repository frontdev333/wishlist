package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	*pgx.Conn
}

func New(ctx context.Context, credentials string) (*DB, error) {
	conn, err := pgx.Connect(ctx, credentials)
	if err != nil {
		return nil, err
	}
	return &DB{
		conn,
	}, nil
}
