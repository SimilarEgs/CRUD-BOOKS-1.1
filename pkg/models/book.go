package models

import (
	"gorm.io/gorm"
)

var dbcon *gorm.DB

type Book struct {
	gorm.Model
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"           validate:"isdefault"`
	BookName   string `gorm:"size:255;not null;unique"   json:"bookname"     validate:"required"`
	BookDate   string `gorm:"size:4;not null"            json:"bookdate"     validate:"required"`
	AuthorName string `gorm:"size:255;not null;unique"   json:"authorname"   validate:"required"`
	AuthorDate string `gorm:"size:4;not null"            json:"authordate"   validate:"required"`
}

// this function will initialize db connection
func InitDB(db *gorm.DB) {
	dbcon = db
	// migrate schema
	db.AutoMigrate(&Book{})
}

/////////////////////////////////////////////////////////////////////////
//                     	sql handlers                                   //
/////////////////////////////////////////////////////////////////////////

func (b *Book) CreateBook() *Book {
	dbcon.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	dbcon.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var book Book
	db := dbcon.Where("ID=?", id).Find(&book)
	return &book, db

}

func DeleteBookByID(id int64) Book {
	var book Book
	dbcon.Where("ID=?", id).Delete(book)
	return book
}

// i leave update function inside controllers.go

// UpdateBookByID will occur as follows:
// 1. find book by sent id
// 2. delete that book
// 3. use data from request to create new book (which will be our updated book)
// 1 - GET 2 - DELETE 3 - POST
