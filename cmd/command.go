package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func RunServer(cmd *echo.Echo) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Port", err)
	}

	cmd.Start(":" + os.Getenv("PORT"))
}
