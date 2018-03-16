package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

var db *gorm.DB

func init() {
	//连接MySQL
	connect()
}

func connect() {
	dbHost := beego.AppConfig.String("mysql::host")
	dbPort := beego.AppConfig.String("mysql::port")
	dbUser := beego.AppConfig.String("mysql::user")
	dbPassword := beego.AppConfig.String("mysql::password")
	dbName := beego.AppConfig.String("mysql::dbname")
	dbCharset := beego.AppConfig.String("mysql::charset")

	dbSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)
	dbObj, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Println("Falied to connect database")
		panic(err)
	}
	defer dbObj.Close()

	db = dbObj
}