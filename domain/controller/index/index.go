package index

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

const (
	appName = "go-boilerplate"
	version = "1.0.0"
)

func Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, fmt.Sprintf("AppName: %s / Version: %s", appName, version))
}
