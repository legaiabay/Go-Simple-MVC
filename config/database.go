package config

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DG database
var DG *gorm.DB

//DBConnect -> connect with gorm
func DBConnect() {
	var err error

	if DG, err = gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME); err != nil {
		DG.LogMode(true)
		log.Println(err.Error())
	} else {
		DG = DG.Set("gorm:auto_preload", true)
	}
}

//DBConnectLama -> connect with native mysql
func DBConnectLama() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
