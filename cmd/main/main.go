package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RoutesHandlers(r)
	fmt.Println("Server started on port - 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
