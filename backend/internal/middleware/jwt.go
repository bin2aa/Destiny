package middleware

import (
	"destiny-backend/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTConfig returns the Echo JWT middleware configuration.
func JWTConfig(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &utils.Claims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(401, map[string]interface{}{
				"success": false,
				"error": map[string]interface{}{
					"code":    401,
					"message": "unauthorized",
				},
			})
		},
	})
}
