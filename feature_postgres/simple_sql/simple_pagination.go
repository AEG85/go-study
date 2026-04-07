package simplesql

import (
	"context"
	"strconv"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp/v3"
)

func ListPages(ctx context.Context, conn *pgx.Conn, limit int64) error {

	booksStore := make(map[string][]models.Book, 1)
	counter := 1
	var offset int64 = 0

	for {

		books, err := getPageRows(ctx, conn, limit, offset)
		if err != nil {
			panic(err)
		}

		if len(books) == 0 {
			break
		}

		pageName := "Страница " + strconv.Itoa(counter)
		booksStore[pageName] = books
		if len(books) < int(limit) {
			break
		} else {
			counter += 1
			offset += limit
		}
	}
	if len(booksStore) > 0 {
		pp.Println(booksStore)
	}
	return nil
}

func getPageRows(ctx context.Context, conn *pgx.Conn, limit int64, offset int64) ([]models.Book, error) {
	sqlQuery := `
		SELECT id, title, author, is_read, review, date_added, date_readed, publication_year  
		FROM  books
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
	`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return []models.Book{}, err
	}
	defer rows.Close()

	booksOnPage := []models.Book{}
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
		booksOnPage = append(booksOnPage, book)
	}

	return booksOnPage, nil
}
