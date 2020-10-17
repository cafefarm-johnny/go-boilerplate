package user

import (
	"github.com/labstack/echo"
	"go-boilerplate/domain/dao/user"
	customError "go-boilerplate/server/configure/error"
	"net/http"
)

var userDao *user.Dao

func init() {
	userDao = user.NewDao()
}

func Users(ctx echo.Context) error {
	users, err := userDao.Find()

	if err != nil {
		return customError.InternalServerErrorResponse(ctx, err)
	}
	if users == nil || len(users) == 0 {
		return customError.BusinessErrorResponse(ctx, customError.NotFound)
	}

	return ctx.JSON(http.StatusOK, users)
}

func User(ctx echo.Context) error {
	name := ctx.Param("name")
	if name == "" {
		return customError.BusinessErrorResponse(ctx, customError.BadRequest)
	}

	findU, err := userDao.FindOne(name)
	if err != nil {
		return customError.BusinessErrorResponse(ctx, customError.NotFound)
	}

	return ctx.JSON(http.StatusOK, findU)
}
