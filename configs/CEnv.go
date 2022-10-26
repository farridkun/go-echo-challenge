package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func COenvURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment", err)
	}

	return os.Getenv("MONGOURI")
}
