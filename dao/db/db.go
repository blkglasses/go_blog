package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

var connStr string = "root:Xxf970524.@tcp(192.168.0.111:3306)/blog"

func InitDB() (err error) {
	DB, err = sqlx.Open("mysql", connStr)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	return nil
}