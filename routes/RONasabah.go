package routes

import (
	"github.com/farridkun/go-echo-challenge/controllers"

	"github.com/labstack/echo/v4"
)

func RONasabah(api *echo.Echo) {
	api.POST("/nasabah", controllers.CreateDataNasabah)
	api.GET("/nasabah/:nasabahId", controllers.GetDataNasabah)
	api.PUT("/nasabah/:nasabahId", controllers.UpdateDataNasabah)
	api.DELETE("/nasabah/:nasabahId", controllers.DeleteDataNasabah)
	api.GET("/nasabah", controllers.GetAllNasabah)
	api.GET("/", controllers.Index)
}
