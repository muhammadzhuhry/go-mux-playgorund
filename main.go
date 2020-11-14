package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

// slice book
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get detail book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params (id)
	// loop through the books and find with id that passed by user
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// insert a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range books {
		if item.ID == param["id"] {
			// delete first
			books = append(books[:index], books[index+1:]...)

			// insert the new one
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = param["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init route
	//r := mux.NewRouter()
	var r *mux.Router = mux.NewRouter()

	// mock data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "476671",
		Title: "Book One",
		Author: &Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "390188",
		Title: "Book Two",
		Author: &Author{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	// route hanlders / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":9000", r))
}
