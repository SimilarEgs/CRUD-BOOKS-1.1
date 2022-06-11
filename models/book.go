package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookID     uint32 `gorm:"primary_key;auto_increment" json:",omitempty"   validate:"isdefault"`
	BookName   string `gorm:"size:255;not null;unique"   json:"bookname"     validate:"required"`
	BookDate   string `gorm:"size:4;not null"            json:"bookdate"     validate:"required"`
	AuthorName string `gorm:"size:255;not null;unique"   json:"authorname"   validate:"required"`
	AuthorDate string `gorm:"size:4;not null"            json:"authordate"   validate:"required"`
	//CreatedAt  time.Time      `gorm:"autoCreateTime:milli"       json:"created_at"`
	//UpdatedAt  time.Time      `gorm:"autoUpdateTime:milli"       json:"updated_at"`
	//DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func GetBook() Book {
	var book Book
	return book
}
func GetBooks() []Book {
	var book []Book
	return book
}
