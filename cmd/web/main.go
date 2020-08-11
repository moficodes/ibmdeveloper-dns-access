package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/ibmdeveloper-domain/internals/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "{'meesage': 'success'}")
	})

	e.GET("/zones", handlers.ListZoneHandler)
	e.GET("/records/:zone", handlers.ZoneRecordsHandler)

	e.POST("/records", handlers.CreateNewRecord)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	e.Logger.Fatal(e.Start(":7777"))
}
