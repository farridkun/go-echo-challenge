package main

import (
	"github.com/farridkun/go-echo-challenge/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	resp := echo.New()

	configs.ConnectDB()

	resp.Logger.Fatal(resp.Start(":9999"))
}
