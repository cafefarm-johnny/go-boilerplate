package server

import (
	"github.com/labstack/echo"
	"log"
)

const (
	port = ":8080"
)

func Startup() {
	e := echo.New()

	log.Fatal(e.Start(port))
}
