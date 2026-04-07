package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOT NULL,
			author VARCHAR(200) NOT NULL,
			review VARCHAR(1024),
			publication_year SMALLINT NOT NULL,
			is_read BOOLEAN NOT NULL,
			date_added TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			date_readed TIMESTAMPTZ

		);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
