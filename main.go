package main

import (
	"fmt"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() { // все мтеоды get post

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")
	router.Static("styles", "./assets/styles")

	testUser := models.CreateTestUser("admin", "admin")

	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	// config := models.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// }

	// Подключаем БД

	// Подключаем маршруты
	routes.IndexRoutes(router, testUser)
	routes.AuthRoutes(router)

	router.Run(":8080")
}
