package service

import (
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
	"github.com/dimasyudhana/xyz-multifinance/utils/encrypt"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
	"github.com/dimasyudhana/xyz-multifinance/utils/toolkit"
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

const (
	SALARY_01 = 10000000
	SALARY_02 = 7000000

	LIMIT_A1 = 200000000
	LIMIT_A2 = 20000000
	LIMIT_A3 = 2000000
)

func (service *authService) Register(data *authentication.CustomerCore) (string, error) {

	data.CustomerId = toolkit.RandomString(15)

	hashedPassword, err := encrypt.HashPassword(data.Password)
	if err != nil {
		return "", err
	}

	data.Password = hashedPassword

	if data.Salary > SALARY_01 {
		data.Limit = LIMIT_A1
	} else if data.Salary <= SALARY_01 && data.Salary >= SALARY_02 {
		data.Limit = LIMIT_A2
	} else if data.Salary < SALARY_02 {
		data.Limit = LIMIT_A3
	}

	registeredId, err := service.authData.Create(data)
	if err != nil {
		return "", err
	}

	return registeredId, nil
}

func (service *authService) Login(email string, password string) (*authentication.CustomerCore, string, error) {

	if err := response.ValidateStruct(service.validate, authentication.CustomerCore{
		Email:    email,
		Password: password,
	}); err != nil {
		return nil, "", err
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

		data.Password = ""

		return data, token, nil
	}

	return nil, "", errors.New(response.ERR_AuthWrongCredentials)
}
