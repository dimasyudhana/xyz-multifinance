package middlewares

import (
	"github.com/dimasyudhana/xyz-multifinance/app/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.InitConfig().JWTSecret),
	})
}
