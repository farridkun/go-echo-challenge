package main

import (
	"github.com/farridkun/go-echo-challenge/configs"
	"github.com/farridkun/go-echo-challenge/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	resp := echo.New()

	configs.ConnectDB()

	routes.RONasabah(resp)

	resp.Logger.Fatal(resp.Start(":9999"))
}
