package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	fmt.Println(os.Getenv("CONN_STRING"))
	return pgx.Connect(ctx, os.Getenv("CONN_STRING"))
}
