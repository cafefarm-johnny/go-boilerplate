package error

import (
	"github.com/labstack/echo"
	"net/http"
)

func CustomHttpErrorHandler(err error, ctx echo.Context) {
	httpError, ok := err.(*echo.HTTPError)

	if ok {
		switch httpError.Code {
		case http.StatusNotFound:
			_ = defaultHttpErrorResponse(ctx, http.StatusNotFound)
			break
		case http.StatusMethodNotAllowed:
			_ = defaultHttpErrorResponse(ctx, http.StatusMethodNotAllowed)
			break
		default:
			_ = defaultHttpErrorResponse(ctx, http.StatusInternalServerError)
			break
		}
	}
}

func defaultHttpErrorResponse(ctx echo.Context, httpStatusCode int) error {
	em := newErrorMessage(httpStatusCode, http.StatusText(httpStatusCode))
	ctx.Logger().Error(em)

	return ctx.JSON(httpStatusCode, em)
}

func BusinessErrorResponse(ctx echo.Context, code StatusCode) error {
	em := newErrorMessage(castInt(code), statusText(code))
	ctx.Logger().Error(em)

	return ctx.JSON(httpCode(code), em)
}

func InternalServerErrorResponse(ctx echo.Context, err error) error {
	em := newErrorMessage(castInt(InternalServerError), statusText(InternalServerError))
	ctx.Logger().Error(err.Error())

	return ctx.JSON(httpCode(InternalServerError), em)
}
