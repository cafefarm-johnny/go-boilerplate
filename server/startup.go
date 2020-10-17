package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

const (
	port = ":8080"
)

func Startup() {
	e := echo.New()
	e.Use(middleware.RequestID(), middleware.MethodOverride(), middleware.Recover())

	log.Fatal(e.Start(port))
}
