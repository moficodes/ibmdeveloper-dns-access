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
	api.GET("/token", handlers.TokenEndpointHandler)

	api.GET("/user", handlers.UserInfoHandler)
	api.GET("/user/:userID/preference", handlers.UserPreferenceHandler)
	api.GET("/auth/login", handlers.LoginHandler)
	api.POST("/auth/authenticate", handlers.AuthenticationHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} [${status}] ${latency_human}\t${uri}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3333"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Logger.Fatal(e.Start(":7777"))
}
