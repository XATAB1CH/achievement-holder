package main

import (
	"fmt"

	"github.com/XATAB1CH/achievement-holder/config"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/XATAB1CH/achievement-holder/routes"
	"github.com/gin-gonic/gin"
)

func main() { // все мтеоды get post

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")
	router.Static("styles", "./assets/styles")

	// Подключаем БД
	config := config.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	conn := postgresql.NewConnect(dsn)
	defer postgresql.CloseConnect(conn)

	// Подключаем маршруты
	routes.IndexRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":8080")
}
