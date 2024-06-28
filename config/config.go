package config

import (
	"fmt"
	"os"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/joho/godotenv"
)

func GetConfig() (config models.Config) {

	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	config = models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return
}
