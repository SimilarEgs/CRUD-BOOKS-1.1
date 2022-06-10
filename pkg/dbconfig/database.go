package dbconfig

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/utils"
)

func CreateDbConection() *gorm.DB {
	db, err := gorm.Open("mysql", utils.ConnectionString)
	if err != nil {
		panic(err)
	}
	log.Print("DB connection was created")
	return db
}
