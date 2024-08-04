package service

import (
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
	"github.com/dimasyudhana/xyz-multifinance/utils/encrypt"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
	"github.com/go-playground/validator/v10"
)

type authService struct {
	authData authentication.AuthDataInterface
	validate *validator.Validate
}

func New(data authentication.AuthDataInterface) authentication.AuthServiceInterface {
	return &authService{
		authData: data,
		validate: validator.New(),
	}
}

func (service *authService) Login(email string, password string) (*authentication.CustomerCore, string, error) {
	dataLogin := authentication.CustomerCore{
		Email:    email,
		Password: password,
	}
	errValidate := response.ValidateStruct(service.validate, dataLogin)
	if errValidate != nil {
		return nil, "", errValidate
	}

	data, err := service.authData.GetCustomerByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if encrypt.CheckPasswordHash(data.Password, password) {
		token, err := encrypt.GenerateToken(data.CustomerId)
		if err != nil {
			return nil, "", errors.New(response.JWT_FailedCreateToken)
		}
		return data, token, nil
	}
	return nil, "", errors.New(response.ERR_AuthWrongCredentials)
}
