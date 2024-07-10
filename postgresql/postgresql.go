package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/XATAB1CH/achievement-holder/config"
	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/jackc/pgx/v5"
)

func GetDSN() string {
	config := config.GetConfig()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
}

func InsertUser(conn *pgx.Conn, name, email, password string) (id int) {

	if name == "" || email == "" || password == "" {
		return 0
	}

	err := conn.QueryRow(context.Background(), `INSERT INTO "users" (name, email, password) VALUES ($1, $2, $3) ON CONFLICT (name) DO NOTHING RETURNING id `, name, email, password).Scan(&id)

	if err == pgx.ErrNoRows {
		return 0
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return id
}

func GetUserByName(conn *pgx.Conn, name string) (user models.User, err error) {
	err = conn.QueryRow(context.Background(), `SELECT id, name, email, password FROM "users" WHERE name = $1`, name).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == pgx.ErrNoRows {
		return user, err
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return user, nil
}
