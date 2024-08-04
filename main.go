package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dimasyudhana/xyz-multifinance/app/config"
	"github.com/dimasyudhana/xyz-multifinance/app/database"
	"github.com/dimasyudhana/xyz-multifinance/app/server"
	migrator "github.com/dimasyudhana/xyz-multifinance/migration"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()

	if args := os.Args; len(args) > 1 {
		switch args[1] {
		case "help":
			fmt.Print(`
Usage:
go run main.go <command>
        
The commands are:
help         Show helper
migrateup    Migrate the DB to the most recent version available
migratedown  Rollback migration

WARNING : please make sure the query in rollback / down file, its may contains DROP TABLE, the data stored will lost
            `)

		case "migrateup":
			migrator.Migrate(*cfg, "up")
		case "migratedown":
			migrator.Migrate(*cfg, "down")
		}
		return
	}

	db := database.InitDatabase(cfg)
	e := echo.New()

	s := server.NewServer(db, e, cfg)
	if err := s.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}

/*
type Customer struct {
	UserId       string
	NIK          string
	Fullname     string
	Legalname    string
	Pob          string
	Dob          string
	Gaji         float64
	FotoKTP      string
	FotoSelfie   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	Limits       []Limit
	Transactions []Transaction
}

type Fee struct {
	FeeId  uint64
	Name   string
	Amount float64
}

type Interest struct {
	InterestId uint64
	Name       string
	Rate       float32
	ProductId  uint64
}

type Limit struct {
	LimitId   uint64
	Name      string
	Amount    float64
	UserId    string
	ProductId uint64
}

type Product struct {
	ProductId uint64
	SKU       string
	Name      string
	Type      string
	Amount    float64
	Stock     uint64
	FeeId     uint64
	Interests []Interest
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
}

type Transaction struct {
	TrxId         string
	NomorKontrak  string
	OTR           string
	AdminFee      float64
	JumlahCicilan uint64
	JumlahBunga   float32
	CreatedAt     time.Time
	CreatedBy     string
	NamaAssets    []Product
}

type TransactionProducts struct {
	ProductId string
	TrxId     string
	Qty       uint64
	Subtotal  float64
	CreatedAt time.Time
	CreatedBy string
}
*/
