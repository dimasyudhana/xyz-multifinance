package handler

import (
	"github.com/dimasyudhana/xyz-multifinance/features/customers"
)

type customerRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func ReqToCore(data customerRequest) *customers.CustomersCore {
	return &customers.CustomersCore{
		Email:    data.Email,
		Phone:    data.Phone,
		Password: data.Password,
	}
}
