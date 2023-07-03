package main

import (
	"go-echo-sample/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userGroup := e.Group("/users")
	routes.UserRoutes(userGroup)

	messageGroup := e.Group("/messages")
	routes.MessageRoutes(messageGroup)

	e.GET("/", getTop)
	e.Logger.Fatal(e.Start(":8080"))
}

func getTop(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
