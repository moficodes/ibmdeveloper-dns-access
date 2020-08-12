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

	api := e.Group("/api")

	api.GET("/zones", handlers.ListZoneHandler)
	api.GET("/records/:zone", handlers.ZoneRecordsHandler)
	api.POST("/records", handlers.CreateNewRecord)

	api.GET("/user", handlers.UserInfoHandler)
	api.GET("/auth/login", handlers.LoginHandler)
	api.GET("/auth/authenticate", handlers.AuthenticationHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3333"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Logger.Fatal(e.Start(":7777"))
}
