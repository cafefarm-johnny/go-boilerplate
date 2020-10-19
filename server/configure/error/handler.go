package error

import (
	"github.com/labstack/echo"
	"net/http"
	"runtime"
)

const (
	errFormat        = "location: %s, \t line: %d, \t cause: %v"
	unknownErrFormat = "location: unknown, \t line: unknown, \t cause: %v"
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
	_, file, line, ok := runtime.Caller(1)
	if ok {
		ctx.Logger().Errorf(errFormat, file, line, em)
	} else {
		ctx.Logger().Errorf(unknownErrFormat, em)
	}

	return ctx.JSON(httpCode(code), em)
}

func InternalServerErrorResponse(ctx echo.Context, err error) error {
	em := newErrorMessage(castInt(InternalServerError), statusText(InternalServerError))
	_, file, line, ok := runtime.Caller(1)
	if ok {
		ctx.Logger().Errorf(errFormat, file, line, err)
	} else {
		ctx.Logger().Errorf(unknownErrFormat, err)
	}

	return ctx.JSON(httpCode(InternalServerError), em)
}
