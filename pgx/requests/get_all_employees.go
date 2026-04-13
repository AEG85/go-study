package requests

import (
	"context"
	"study/models"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]models.Employee, error) {
	sqlQuery := `
		SELECT id, full_name, position 
		FROM  employees
		ORDER BY full_name ASC
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return []models.Employee{}, err
	}
	defer rows.Close()

	employees := []models.Employee{}
	for rows.Next() {
		var employee models.Employee
		rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Position,
		)
		employees = append(employees, employee)
	}
	return employees, nil
}
