package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadKey() []byte {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var key = []byte(os.Getenv("KEY"))

	return key
}
