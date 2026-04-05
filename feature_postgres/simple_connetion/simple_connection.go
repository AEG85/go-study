package simpleconnetion

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:1985@localhost:5432/")

	if err != nil {
		return nil, err
	}

	return conn, nil
}
