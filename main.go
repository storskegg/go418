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

type teapotJSON struct {
	Status string `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}

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
	msg := &teapotJSON{
		Status:"error",
		Message:"I'm a teapot",
	}
	return c.JSON(http.StatusTeapot, msg)
}
