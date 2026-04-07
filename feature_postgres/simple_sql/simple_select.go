package simplesql

import (
	"context"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]models.Book, error) {
	sqlQuery := `
		SELECT id, title, author, is_read, review, date_added, date_readed, publication_year  
		FROM  books
		ORDER BY title ASC
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return []models.Book{}, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.IsRead,
			&book.Review,
			&book.DateAdded,
			&book.DateReaded,
			&book.PublicationYear,
		)
		books = append(books, book)
	}
	return books, nil

}
