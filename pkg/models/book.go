package models

import (
	dbconfig "github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/dbconfig"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	BookID     uint32 `gorm:"primary_key;auto_increment" json:"id"           validate:"isdefault"`
	BookName   string `gorm:"size:255;not null;unique"   json:"bookname"     validate:"required"`
	BookDate   string `gorm:"size:4;not null"            json:"bookdate"     validate:"required"`
	AuthorName string `gorm:"size:255;not null;unique"   json:"authorname"   validate:"required"`
	AuthorDate string `gorm:"size:4;not null"            json:"authordate"   validate:"required"`
}

// this function will initialize db connection
func Init() {

	dbconfig.CreateDbConnection()
	db = dbconfig.GetDB()

	// migrate schema
	db.AutoMigrate(&Book{})
}

/////////////////////////////////////////////////////////////////////////
//                     	sql handlers                                   //
/////////////////////////////////////////////////////////////////////////

func (b *Book) CreateBook() *Book {

	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db

}

func DeleteBookByID(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}

// i leave update function inside controllers.go

// UpdateBookByID will occur as follows:
// 1. find book by sent id
// 2. delete that book
// 3. use data from request to create new book (which will be our updated book)
// 1 - GET 2 - DELETE 3 - POST
