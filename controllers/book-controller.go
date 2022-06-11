package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/models"
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/responses"
)

// todo:
// it's better to isolate db logic
// and get db connection through the functin call
// from a coresponding package
var dbcon *gorm.DB

func InitDB(db *gorm.DB) {

	dbcon = db
	db.AutoMigrate(&models.Book{})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book = models.GetBook()

	// processed request body
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// validate request data
	validate := validator.New()
	err = validate.Struct(book)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// insert request entity into DB
	err = dbcon.Create(&book).Error
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "[Info] book was sucesfully created")
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book = models.GetBook()

	// extracting URL params
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		responses.JSON(w, http.StatusUnprocessableEntity, "[Error] id convertion failed")
		return
	}

	// look into db with extracted ID, handle errors
	res := dbcon.First(&book, "id = ?", id)

	if res.Error != nil {
		log.Println(res.Error)

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			responses.JSON(w, http.StatusNotFound, "[Info] record not found")
		} else {
			responses.JSON(w, http.StatusInternalServerError, "[Error] failed while extracting book object")
		}
		return
	}
	responses.JSON(w, http.StatusOK, book)

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var books = models.GetBooks()

	rw := dbcon.Find(&books)

	// checking existing records
	if rw.RowsAffected == 0 {
		log.Println("[Info] no entries in the DB table")
		responses.JSON(w, http.StatusNotFound, "[Info] no entries in the DB table")
		return
	}

	responses.JSON(w, http.StatusOK, books)
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book = models.GetBook()

	// extracting URL params
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		responses.JSON(w, http.StatusUnprocessableEntity, "[Error] id convertion failed")
		return
	}

	// look into db with extracted ID, handle errors
	res := dbcon.First(&book, "id = ?", id)

	if res.Error != nil {
		log.Println(res.Error)

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			responses.JSON(w, http.StatusNotFound, "[Info] record not found")
		} else {
			responses.JSON(w, http.StatusInternalServerError, "[Error] occurred while deleting book entity")
		}
		return
	}

	rw := dbcon.Delete(&book)
	// checking deletion resualt
	if rw.RowsAffected != 1 {
		responses.JSON(w, http.StatusInternalServerError, "[Error] occurred while deleting book entity")
		return
	}

	responses.JSON(w, http.StatusNoContent, "[Info] book was sucesfully deleted")
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
}
