package simplesql

import (
	"context"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, book models.Book) (pgconn.CommandTag, error) {
	sqlQuery := `
		UPDATE books 
		SET 
			title = $1, 
			author = $2, 
			review = $3, 
			publication_year = $4, 
			is_read = $5, 
			date_readed = $6
		WHERE id = $7
	`
	return conn.Exec(ctx, sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.PublicationYear,
		book.IsRead,
		book.DateReaded,
		book.ID,
	)
}
