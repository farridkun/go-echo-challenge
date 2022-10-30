package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Environment() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment", err)
	}

	return os.Getenv("MONGOURI")
}
