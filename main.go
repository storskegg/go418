package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	address = ""
	port = "3500"
)

func main() {
	srv := echo.New()

	// Middleware
	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())

	// Routes
	srv.Any("*", teapot)

	srv.Logger.Fatal(srv.Start(address + ":" + port))
}

func teapot(c echo.Context) error {
	return c.JSON(http.StatusTeapot, "{\"status\":\"error\",\"msg\":\"I'm a teapot.\"}")
}
