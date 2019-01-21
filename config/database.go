package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
