package config

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

// TODO this is temporary file, db config handling should be improved

func GetDB() (*sql.DB, error) {
	dbHost := "127.0.0.1" // TODO read from config file
	dbPort := "3308"
	dbUser := "root"
	dbPass := "root"
	dbName := "charlie"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	//fmt.Println(dsn)
	db, err := sql.Open(`mysql`, dsn)
	//	defer db.Close()
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
