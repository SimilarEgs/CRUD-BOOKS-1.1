package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/models"
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/responses"
)

var dbcon *gorm.DB

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book = models.GetBook()

	// processed request body
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// validate request data
	validate := validator.New()
	err = validate.Struct(book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// insert request entity into DB
	err = dbcon.Create(&book).Error
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "book was sucesfully created")
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
}

func InitDB(db *gorm.DB) {

	dbcon = db
	db.AutoMigrate(&models.Book{})
}
