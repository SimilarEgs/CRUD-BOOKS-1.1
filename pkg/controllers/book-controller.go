package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/models"
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/responses"
)

var Book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

	//reading request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	book := models.Book{}
	//unmarshaling body
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//validate request data
	validate := validator.New()
	err = validate.Struct(book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//insert request data into DB via «CreateBook» method
	b := book.CreateBook()
	//testing
	res, _ := json.Marshal(b)
	w.Write(res)
	responses.JSON(w, http.StatusCreated, b)

}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {

}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {

}
