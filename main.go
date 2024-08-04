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
