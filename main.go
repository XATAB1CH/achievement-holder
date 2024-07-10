package main

import (
	"context"

	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/XATAB1CH/achievement-holder/routes"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Подключаем БД

	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	// Подключаем маршруты
	router := routes.InitRoutes()

	router.Run(":8080")

}
