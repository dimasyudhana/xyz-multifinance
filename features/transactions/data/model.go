package data

import "time"

type Transaction struct {
	TransactionId   string
	CustomerId      string
	NomorKontrak    string
	OTR             float32
	JumlahCicilan   float32
	JumlahBunga     float32
	AdminFee        float32
	TransactionDate time.Time
}

type TransactionProduct struct {
	TransactionId string
	ProductId     string
}
