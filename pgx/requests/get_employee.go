package requests

import (
	"context"
	"errors"
	"study/models"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn, employeeID int) (models.Employee, error) {
	sqlQuery := `
		SELECT id, full_name, position 
		FROM  employees
		WHERE id = $1
	`
	var employee models.Employee
	err := conn.QueryRow(ctx, sqlQuery, employeeID).Scan(
		&employee.ID,
		&employee.FullName,
		&employee.Position,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Employee{}, errors.New("сотрудник не найден")
		}
		return models.Employee{}, err
	}

	return employee, nil
}
