package routes

import (
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/controllers"
	"github.com/gorilla/mux"
)

var (
	RoutesHandlers = func(r *mux.Router) {
		r.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")
		r.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
		r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
		r.HandleFunc("/book/{id}", controllers.UpdateBookByID).Methods("PUT")
		r.HandleFunc("/book/{id}", controllers.DeleteBookByID).Methods("DELETE")
	}
)