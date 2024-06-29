package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Author   string
	Password string
	DBName   string
	SSLMode  string
}

func GetConfig() (config Config) {

	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	config = Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Author:   os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return
}
