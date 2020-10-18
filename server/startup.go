package server

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-boilerplate/domain/route"
	customError "go-boilerplate/server/configure/error"
	"log"
	"os"
)

const (
	port = ":8080"
)

func init() {
	validate()
}

func Startup() {
	e := echo.New()
	e.Use(middleware.RequestID(), middleware.MethodOverride(), middleware.Recover())
	e.HTTPErrorHandler = customError.CustomHttpErrorHandler

	route.Register(e)

	log.Fatal(e.Start(port))
}

func validate() {
	valid := false

	profile := os.Getenv("profile")

	switch profile {
	case "local", "dev", "stage", "prod":
		valid = !valid
		break
	default:
		break
	}

	if !valid {
		fmt.Println("Startup Failed. Cause: 'profile' environment variable is nil.")
		os.Exit(1)
	}
}
