package storage

import (
	"context"
	"fmt"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/jackc/pgx/v5"
)

func GetUserStorage() []models.User {
	dsn := postgresql.GetDSN()
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	rows, err := conn.Query(context.Background(), `SELECT id, name, email FROM "users"`)
	if err != nil {
		fmt.Println("Query error")
	}
	defer rows.Close()

	userStorage := []models.User{}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Println("Scan error")
		}
		userStorage = append(userStorage, user)
	}

	return userStorage
}

func GetAchievementsByID(id string) []models.Achievement {
	dsn := postgresql.GetDSN()
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	rows, err := conn.Query(context.Background(), `SELECT title, info, image FROM "achievements" WHERE user_id = $1`, id)
	if err != nil {
		fmt.Println("Query error")
	}
	defer rows.Close()

	achievementStorage := []models.Achievement{}

	for rows.Next() {
		var ach models.Achievement

		err = rows.Scan(&ach.Title, &ach.Image, &ach.Info)
		if err != nil {
			fmt.Println("Scan error")
		}
		achievementStorage = append(achievementStorage, ach)
	}

	return achievementStorage
}
