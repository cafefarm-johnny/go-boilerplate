package route

import (
	"github.com/labstack/echo"
	"go-boilerplate/domain/controller/index"
	"net/http"
)

func Register(e *echo.Echo) {
	e.GET("/", index.Index)

	v1(e)
}

func v1(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.GET("/users", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNoContent)
	})
}
