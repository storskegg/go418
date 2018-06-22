package main

import (
	"net/http"

	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shirou/gopsutil/host"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	address = "127.0.0.1"
	port    = "3500"
)

type systemJSON struct {
	Make            string `json:"make" xml:"make"`
	Model           string `json:"model" xml:"model"`
	Year            int    `json:"year" xml:"year"`
	Uptime          uint64 `json:"uptime" xml:"uptime"`
	UptimeFormatted string `json:"uptimeFormatted" xml:"uptime-formatted"`
}

type teapotJSON struct {
	Status  string     `json:"status" xml:"status"`
	Message string     `json:"message" xml:"message"`
	System  systemJSON `json:"system" xml:"system"`
}

func main() {
	srv := echo.New()

	// Middleware
	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())

	// Routes
	srv.Any("*", func(c echo.Context) error {
		return teapot(c, srv.Logger)
	})

	// Get the teapot started
	srv.Logger.Fatal(srv.Start(address + ":" + port))
}

func teapot(c echo.Context, logger echo.Logger) error {
	uptime, uptimeFormatted := uptime()

	logger.Print("Uptime: ", uptime)

	system := &systemJSON{
		Make:            "Fiestaware",
		Model:           "Teapot (Fiesta Red)",
		Year:            1958,
		Uptime:          uptime,
		UptimeFormatted: uptimeFormatted,
	}
	msg := &teapotJSON{
		Status:  "error",
		Message: "I'm a teapot",
		System:  *system,
	}
	return c.JSON(http.StatusTeapot, msg)
}

func uptime() (uint64, string) {
	uptimeSeconds, _ := host.Uptime()
	uptimeSeconds = uptimeSeconds * 1000000000
	uptime := time.Duration(uptimeSeconds).String()
	split := strings.Split(uptime, "h")
	hours, _ := strconv.ParseInt(split[0], 10, 64)
	split = strings.Split(split[1], "m")
	minutes, _ := strconv.ParseInt(split[0], 10, 64)
	seconds, _ := strconv.ParseInt(strings.TrimRight(split[1], "s"), 10, 64)

	days := math.Floor(float64(hours) / 24)
	hours = hours % 24

	return uptimeSeconds, fmt.Sprintf("%v days, %v hours, %v minutes, %v seconds", days, hours, minutes, seconds)
}
