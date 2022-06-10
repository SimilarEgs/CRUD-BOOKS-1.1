package dbconfig

import (
	"log"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDBconnection() *gorm.DB {
	dsn := utils.ConnectionString
	log.Fatal(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("DBconnection was created")

	return db
}
