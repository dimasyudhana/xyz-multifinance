package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/dimasyudhana/xyz-multifinance/app/config"
)

func InitDatabase(cfg *config.AppConfig) *sql.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Hostname, cfg.DB.Port, cfg.DB.Name)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	return db

}
