package main

import (
	"context"

	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/XATAB1CH/achievement-holder/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() { // все мтеоды get post

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")
	router.Static("styles", "./assets/styles")

	// Подключаем БД

	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	// Подключаем маршруты
	routes.IndexRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":8080")

}
