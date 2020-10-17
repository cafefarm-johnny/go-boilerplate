package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-boilerplate/domain/route"
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

	route.Register(e)

	log.Fatal(e.Start(port))
}
