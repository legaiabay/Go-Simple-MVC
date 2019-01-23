package config

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DG database
var DG *gorm.DB

//DBConnect -> connect with gorm
func DBConnect() {
	var err error

	if DG, err = gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?parseTime=true"); err != nil {
		DG.LogMode(true)
		log.Println(err.Error())
	} else {
		DG = DG.Set("gorm:auto_preload", true)
	}
}
