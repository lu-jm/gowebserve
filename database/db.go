package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySql() error {
	dsn := "root:123456@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("db connect fail:%v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("db ping error:%v", err)
	}

	log.Println("db connect success")
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
