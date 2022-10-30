package main

import (
	"github.com/farridkun/go-echo-challenge/app/routes"
	"github.com/farridkun/go-echo-challenge/cmd"
	"github.com/farridkun/go-echo-challenge/infra/database"

	"github.com/labstack/echo/v4"
)

func main() {
	resp := echo.New()

	database.ConnectDB()

	routes.RONasabah(resp)

	cmd.RunServer(resp)
}
