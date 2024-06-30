package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewConnect(dsn string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func CloseConnect(conn *pgx.Conn) {
	conn.Close(context.Background())
}

// Добавление пользователя в базу данных
func InsertUser(conn *pgx.Conn, name, email, password string) int {
	var id int
	err := conn.QueryRow(context.Background(), `INSERT INTO "user" (name, email, password) VALUES ($1, $2, $3) RETURNING id`, name, email, password).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return id
}
