package filters

import (
	"errors"
	"lazy/common/types"
	"lazy/domain/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HttpFilter(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	httpCode := http.StatusInternalServerError
	code := err.Error()

	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		httpCode = echoErr.Code
		code = echoErr.Message.(string)
	} else {
		switch {
		case errors.Is(err, user.ErrInvalidEmailFormat), errors.Is(err, user.ErrNameTooShort), errors.Is(err, user.ErrPasswordTooShort):
			httpCode = http.StatusUnprocessableEntity

		case errors.Is(err, user.ErrEmailAlreadyInUse):
			httpCode = http.StatusConflict

		case errors.Is(err, user.ErrUserDoesNotExist):
			httpCode = http.StatusNotFound
		}
	}

	ctx.JSON(httpCode, types.ErrorResponse{
		TraceId: ctx.Response().Header().Get(echo.HeaderXRequestID),
		Error: types.ErrorDetail{
			Code: code,
		},
	})
}
