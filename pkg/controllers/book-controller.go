package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/models"
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/utils"
)


var Book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request){

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	book := Book
	//unmurshal body
	err = json.Unmarshal(body, &book)
	if err != nil{
		
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request){

}

func GetAllBooks(w http.ResponseWriter, r *http.Request){

}

func DeleteBookByID(w http.ResponseWriter, r *http.Request){

}

func UpdateBookByID(w http.ResponseWriter, r *http.Request){
	
}