package main

import (
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

	// Подключаем маршруты
	routes.IndexRoutes(router, nil)
	routes.AuthRoutes(router)

	router.Run(":8080")
}
