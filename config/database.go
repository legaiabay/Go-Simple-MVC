package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"gotest.com/go-sandbox/structs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DG database
var DB *gorm.DB

//DBConnect -> connect with gorm
func DBConnect() {
	var err error

	if DB, err = gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?parseTime=true"); err != nil {
		DB.LogMode(true)
		log.Println(err.Error())
	} else {
		DB = DB.Set("gorm:auto_preload", true)
		DBMigration()
	}
}

//DBMigration -> Migration
func DBMigration() {
	DB.AutoMigrate(
		&structs.Todo{},
		&structs.TodoGroup{},
	)
}
