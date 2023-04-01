package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_Driver = "mysql"
var DB_Host = "127.0.0.1"
var DB_Port = "3306"
var DB_Name = "btpn-finaltask"
var DB_User = "root"
var DB_Password = ""

var DB *gorm.DB

var errConnection error
func ConnectDatabase() {
	

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",DB_User,DB_Password,DB_Host,DB_Port, DB_Name)
		DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	

	if errConnection != nil {
		panic("Can't connect to database")
	}

	log.Println("Connecting to database")
}