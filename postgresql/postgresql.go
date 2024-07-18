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

func GetUserByID(conn *pgx.Conn, id int) (user models.User, err error) {
	err = conn.QueryRow(context.Background(), `SELECT id, name, email, password FROM "users" WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == pgx.ErrNoRows {
		return user, err
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return user, nil
}

func InsertAchievement(conn *pgx.Conn, title, image, info string, userID int) (id int) {
	if title == "" || image == "" || info == "" {
		return 0
	}

	err := conn.QueryRow(context.Background(), `INSERT INTO "achievements" (title, image, info, user_id) VALUES ($1, $2, $3, $4) ON CONFLICT (title) DO NOTHING RETURNING id `, title, image, info, userID).Scan(&id)

	if err == pgx.ErrNoRows {
		return 0
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return id

}

func GetAchievementsByUserID(conn *pgx.Conn, userID int) ([]models.Achievement, error) {
	var achievements []models.Achievement

	rows, err := conn.Query(context.Background(), `SELECT id, title, image, info, user_id FROM "achievements" WHERE user_id = $1`, userID)

	if err == pgx.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var achievement models.Achievement
		err = rows.Scan(&achievement.ID, &achievement.Title, &achievement.Image, &achievement.Info, &achievement.UserID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

func GetAchievementByID(conn *pgx.Conn, id int) (models.Achievement, error) {
	var achievement models.Achievement

	err := conn.QueryRow(context.Background(), `SELECT id, title, image, info, user_id FROM "achievements" WHERE id = $1`, id).Scan(&achievement.ID, &achievement.Title, &achievement.Image, &achievement.Info, &achievement.UserID)

	if err == pgx.ErrNoRows {
		return achievement, err
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return achievement, nil
}

func DeleteAchievement(conn *pgx.Conn, id int) error {
	_, err := conn.Exec(context.Background(), `DELETE FROM "achievements" WHERE id = $1`, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return nil
}
