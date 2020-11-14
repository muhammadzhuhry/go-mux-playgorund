package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// model book
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// model author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	// init route
	//r := mux.NewRouter()
	var r *mux.Router = mux.NewRouter()

	// route hanlders / endpoints
	//r.HandleFunc("/api/books", getBooks).Methods("GET")
	//r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	//r.HandleFunc("/api/book", createBook).Methods("POST")
	//r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":9000", r))
}
