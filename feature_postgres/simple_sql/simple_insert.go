package simplesql

import (
	"context"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, user models.User) (pgconn.CommandTag, error) {
	sqlQuery := `
		INSERT INTO users (full_name, phone_number)
		VALUES ($1, $2)
	`
	return conn.Exec(ctx, sqlQuery,
		user.FullName,
		user.PhoneNumber,
	)
}
