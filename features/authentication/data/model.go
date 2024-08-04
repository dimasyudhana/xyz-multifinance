package data

import (
	"time"

	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
)

type Customers struct {
	CustomerId      string
	Email           string
	Phone           string
	Password        string
	Fullname        string
	Legalname       string
	Nik             string
	Pob             string
	Dob             time.Time
	Salary          float64
	OCRPicture      string
	LivenessPicture string
	Limit           float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func CoreToData(data *authentication.CustomerCore) *Customers {
	return &Customers{
		CustomerId:      data.CustomerId,
		Email:           data.Email,
		Phone:           data.Phone,
		Password:        data.Password,
		Fullname:        data.Fullname,
		Legalname:       data.Legalname,
		Nik:             data.Nik,
		Pob:             data.Pob,
		Dob:             data.Dob,
		Salary:          data.Salary,
		OCRPicture:      data.OCRPicture,
		LivenessPicture: data.LivenessPicture,
		Limit:           data.Limit,
	}
}
