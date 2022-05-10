package main

import (
	"awesomeProject/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:kartikeya@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//db.AutoMigrate(&models.Book{})

	return db
}

func main() {

	//Initializing database

	DB := Init()
	h := Handlers.New(DB)

	// API ROUTES

	router := mux.NewRouter()

	router.HandleFunc("/product", h.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	router.HandleFunc("/product", h.AddProducts).Methods("POST")
	router.HandleFunc("/product/{id}/{qty}", h.BuyProduct).Methods("PUT")
	router.HandleFunc("/products/top", h.TopProducts).Methods("GET")

	fmt.Printf("Starting server at port :8080 \n")
	log.Fatal(http.ListenAndServe(":8080", router))

}
