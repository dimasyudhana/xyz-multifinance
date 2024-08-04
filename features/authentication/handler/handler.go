package handler

import (
	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService authentication.AuthServiceInterface
}

func New(service authentication.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (handler *AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		authInput := new(loginRequest)
		errBind := c.Bind(&authInput)
		if errBind != nil {
			jsonResponse, httpCode := response.WebResponseError(errBind, response.FEAT_AUTH_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		dataUser, token, err := handler.authService.Login(authInput.Email, authInput.Password)
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(err, response.FEAT_AUTH_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := response.WebResponseSuccess("[success] login", response.FEAT_AUTH_CODE, map[string]any{
			"token": token,
			"email": &dataUser.Email,
		})
		return c.JSON(httpCode, mapResponse)
	}
}
