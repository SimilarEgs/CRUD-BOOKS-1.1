package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SimilarEgs/CRUD-BOOKS-1.1/pkg/models"
	"github.com/SimilarEgs/CRUD-BOOKS-1.1/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	var db = db.CreateDBconnection()
	models.InitDB(db)

	r := mux.NewRouter()
	routes.RoutesHandlers(r)

	fmt.Println("Server started on port - 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
