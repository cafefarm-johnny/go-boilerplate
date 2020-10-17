package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	customError "go-boilerplate/server/configure/error"
	"log"
)

const (
	port = ":8080"
)

func Startup() {
	e := echo.New()
	e.Use(middleware.RequestID(), middleware.MethodOverride(), middleware.Recover())
	e.HTTPErrorHandler = customError.CustomHttpErrorHandler

	log.Fatal(e.Start(port))
}
