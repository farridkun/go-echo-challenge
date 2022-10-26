package routes

import (
	"github.com/farridkun/go-echo-challenge/controllers"

	"github.com/labstack/echo/v4"
)

func RONasabah(api *echo.Echo) {
	api.POST("/nasabah", controllers.COCreateNasabah)
	api.GET("/nasabah/:nasabahId", controllers.GetDataNasabah)
	api.PUT("/nasabah/:nasabahId", controllers.UpdateDataNasabah)
}
