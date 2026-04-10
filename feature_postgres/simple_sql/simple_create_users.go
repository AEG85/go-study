package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(250) NOT NULL,
			phone_number VARCHAR(200)
		);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
