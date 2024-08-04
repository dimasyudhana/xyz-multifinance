package router

import (
	"database/sql"
	"net/http"

	_authFactory "github.com/dimasyudhana/xyz-multifinance/features/authentication/factory"
	_authAPI "github.com/dimasyudhana/xyz-multifinance/features/authentication/handler"

	_customerFactory "github.com/dimasyudhana/xyz-multifinance/features/customers/factory"
	_customerAPI "github.com/dimasyudhana/xyz-multifinance/features/customers/handler"

	"github.com/dimasyudhana/xyz-multifinance/app/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type appsFactory struct {
	authHandler     *_authAPI.AuthHandler
	customerHandler *_customerAPI.CustomerHandler
}

func InitRouter(db *sql.DB, e *echo.Echo) {

	sysRoute := appsFactory{
		authHandler:     _authFactory.New(db),
		customerHandler: _customerFactory.New(db),
	}

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.POST("/login", sysRoute.authHandler.Login())
	e.POST("/registrasi", sysRoute.customerHandler.Insert(), middlewares.JWTMiddleware())

}
