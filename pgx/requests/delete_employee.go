package requests

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, employeeIDs []int) (pgconn.CommandTag, error) {
	sqlQuery := `
		DELETE from employees 
		WHERE id = ANY($1)
	`
	return conn.Exec(ctx, sqlQuery, employeeIDs)
}
