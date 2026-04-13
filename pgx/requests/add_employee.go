package requests

import (
	"context"
	"study/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, employee models.Employee) (pgconn.CommandTag, error) {
	sqlQuery := `
		INSERT INTO employees (full_name, position)
		VALUES ($1, $2)
	`
	return conn.Exec(ctx, sqlQuery,
		employee.FullName,
		employee.Position,
	)
}
