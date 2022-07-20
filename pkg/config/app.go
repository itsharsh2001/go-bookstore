package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() { //Connect function helps to connect to Database
	d, err := gorm.Open("mysql", "root:itsharsh2001/simplerest?charset=utf8&parseTime=True&loc=Local") //username:password/tablename?iske baad ka mysql mangta hai

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
