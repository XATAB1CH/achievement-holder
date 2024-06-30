package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/XATAB1CH/achievement-holder/config"
	"github.com/jackc/pgx/v5"
)

func GetDSN() string {
	config := config.GetConfig()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
}

// Добавление пользователя в базу данных
func InsertUser(conn *pgx.Conn, name, email, password string) int {
	var id int
	err := conn.QueryRow(context.Background(), `INSERT INTO "users" (name, email, password) VALUES ($1, $2, $3) RETURNING id`, name, email, password).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return id
}
