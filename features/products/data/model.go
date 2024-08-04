package data

import "time"

type Products struct {
	ProductId string
	Category  string
	SKU       string
	Name      string
	Amount    float32
	Stock     uint32
	Tenor     uint32
	Bunga     uint32
	Fee       uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
