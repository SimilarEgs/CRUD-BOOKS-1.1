package dbconfig

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// this function will serve us to create connection to mySQL db
func CreateDbConnection() {
	d, err := gorm.Open("mysql", "root:Eretika12@tcp(127.0.0.1:3306)/golang_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// purpose of this function to transfer the db connection in another part of the program
func GetDB() *gorm.DB {
	return db
}
