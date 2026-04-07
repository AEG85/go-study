package simplesql

import (
	"context"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, book models.Book) (pgconn.CommandTag, error) {
	sqlQuery := `
		INSERT INTO books (title, author, review, publication_year, is_read, date_readed)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	return conn.Exec(ctx, sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.PublicationYear,
		book.IsRead,
		book.DateReaded,
	)
}
