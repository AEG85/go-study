package simplesql

import (
	"context"
	"study/feature_postgres/models"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]models.User, error) {
	sqlQuery := `
		SELECT id, full_name, phone_number  
		FROM  users
		ORDER BY full_name ASC
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		rows.Scan(
			&user.ID,
			&user.FullName,
			&user.PhoneNumber,
		)
		users = append(users, user)
	}
	return users, nil
}
