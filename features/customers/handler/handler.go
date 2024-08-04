package handler

import (
	"log"

	"github.com/dimasyudhana/xyz-multifinance/features/customers"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	srv customers.Service
}

func New(srv customers.Service) *CustomerHandler {
	return &CustomerHandler{
		srv: srv,
	}
}

func (ch *CustomerHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		/*
			_, role, errToken := middlewares.ExtractTokenJWT(c)
			log.Println("role:", role)
			if errToken != nil {
				jsonResponse, httpCode := response.WebResponseError(errors.New(response.JWT_FailedCastingJwtToken), response.FEAT_USER_CODE)
				return c.JSON(httpCode, jsonResponse)
			}
		*/
		requestBody := new(customerRequest)
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(err, response.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		res, err := ch.srv.Insert(ReqToCore(*requestBody))
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(err, response.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		log.Println(res)
		mapResponse, httpCode := response.WebResponseSuccess("[success] create data", response.FEAT_USER_CODE, nil)
		return c.JSON(httpCode, mapResponse)
	}
}
