package requests

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(250) NOT NULL,
			position VARCHAR(200) NOT NULL,

			UNIQUE(position, full_name)
		);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
