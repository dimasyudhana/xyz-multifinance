package customers

import (
	"time"
)

type CustomersCore struct {
	CustomerId      string     `json:"user_id"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Password        string     `json:"password"`
	Fullname        string     `json:"fullname"`
	Legalname       string     `json:"legalname"`
	Nik             string     `json:"nik"`
	Pob             string     `json:"pob"`
	Dob             time.Time  `json:"dob"`
	Salary          float64    `json:"salary"`
	OCRPicture      string     `json:"ocr_picture"`
	LivenessPicture string     `json:"liveness_picture"`
	Limit           float64    `json:"limit"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	TransactionId   string     `json:"transaction_id"`
	NomorKontrak    string     `json:"nomor_kontrak"`
	NamaAsset       string     `json:"nama_asset"`
	OTR             string     `json:"otr"`
	JumlahCicilan   string     `json:"jumlah_cicilan"`
	JumlahBunga     string     `json:"jumlah_bunga"`
	AdminFee        string     `json:"admin_fee"`
	TransactionDate string     `json:"transaction_date"`
}

type Service interface {
	CustomerTransactions(customerId string) (<-chan *CustomersCore, error)
}

type Data interface {
	CustomerTransactions(customerId string) ([]*CustomersCore, error)
}
