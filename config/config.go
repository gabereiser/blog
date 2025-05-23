package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {

		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Print("Error loading .env file\n")
		}
	}

	return os.Getenv(key)
}
