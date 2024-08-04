package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
)

type registerRequest struct {
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	Fullname        string `json:"fullname"`
	Legalname       string `json:"legalname"`
	Nik             string `json:"nik"`
	Pob             string `json:"pob"`
	Dob             string `json:"dob"`
	Salary          string `json:"salary"`
	OCRPicture      string `json:"ocr_picture"`
	LivenessPicture string `json:"liveness_picture"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(request registerRequest) *authentication.CustomerCore {
	dob, _ := time.Parse("02/01/2006", request.Dob)
	salary, _ := strconv.ParseFloat(strings.Trim(request.Salary, "."), 64)
	return &authentication.CustomerCore{
		Email:           request.Email,
		Phone:           request.Phone,
		Password:        request.Password,
		Fullname:        request.Fullname,
		Legalname:       request.Legalname,
		Nik:             request.Nik,
		Pob:             request.Pob,
		Dob:             dob,
		Salary:          salary,
		OCRPicture:      request.OCRPicture,
		LivenessPicture: request.LivenessPicture,
	}
}
