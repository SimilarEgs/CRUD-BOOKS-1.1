package models

import (
	dbconfig "github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/DBconfig"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	BookID     string `gorm:"primary_key" json:"id"`
	BookName   string `json:"bookname"`
	BookDate   string `json:"bookdate"`
	AuthorName string `json:"authorname"`
	AuthorDate string `json:"authordate"`
}

// this function will initialize db connection
func init() {
	dbconfig.CreateDbConnection()
	db = dbconfig.GetDB()
	// migrate schema
	db.AutoMigrate(&Book{})
}
