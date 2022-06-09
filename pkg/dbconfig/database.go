package dbconfig

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/utils"
)

var (
	db *gorm.DB
)

// this function will serve us to create connection to mySQL db
func CreateDbConnection() {

	d, err := gorm.Open("mysql", utils.ConnectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

// purpose of this function to transfer the db connection in another part of the program
func GetDB() *gorm.DB {
	return db
}
