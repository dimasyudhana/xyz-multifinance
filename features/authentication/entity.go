package authentication

import "time"

type CustomerCore struct {
	CustomerId      string    `json:"user_id"`
	Email           string    `json:"email" validate:"required,email"`
	Phone           string    `json:"phone"`
	Password        string    `json:"password" validate:"required"`
	Fullname        string    `json:"fullname"`
	Legalname       string    `json:"legalname"`
	Nik             string    `json:"nik"`
	Pob             string    `json:"pob"`
	Dob             time.Time `json:"dob"`
	Salary          float64   `json:"salary"`
	OCRPicture      string    `json:"ocr_picture"`
	LivenessPicture string    `json:"liveness_picture"`
	Limit           float64   `json:"limit"`
}

type AuthDataInterface interface {
	Create(data *CustomerCore) (string, error)
	GetCustomerByEmail(email string) (*CustomerCore, error)
}

type AuthServiceInterface interface {
	Register(data *CustomerCore) (string, error)
	Login(email string, password string) (*CustomerCore, string, error)
}
