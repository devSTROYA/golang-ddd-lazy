package guards

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "NO_TOKEN_PROVIDED")
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "INVALID_TOKEN_FORMAT")
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed):
				return echo.NewHTTPError(http.StatusUnauthorized, "TOKEN_MALFORMED")
			case errors.Is(err, jwt.ErrTokenSignatureInvalid):
				return echo.NewHTTPError(http.StatusUnauthorized, "INVALID_SIGNATURE")
			case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
				return echo.NewHTTPError(http.StatusUnauthorized, "TOKEN_EXPIRED")
			case parsedToken.Valid:
				userId, err := parsedToken.Claims.GetSubject()
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, err)
				}

				ctx.Set("USER_ID", userId)

				return next(ctx)
			default:
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}

		}
	}
}
