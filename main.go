package main

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "POST called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "PUT called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "DELETE called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}


type Book struct {
	Id    string    `json:"id"`
	Title string    `json:"title"`
	Price float64   `json:"price"`
	Author *Author  `json:"author"`
	Year  int       `json:"year"`
	Rating float64  `json:"rating"`
}

type Author struct {
	Id   string    `json:"id"`
	Name string `json:"name"`
	Rating float64 `json:"rating"`
}

func main() {
	// create a new router
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)
	// define routes
	// http.HandleFunc("/books", getBooks)
	// http.HandleFunc("/authors", getAuthors)
	// http.HandleFunc("/books/add", addBook)
	// http.HandleFunc("/authors/add", addAuthor)

	// start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handler to get all books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	// books := []Book{
// 	// 	{Id: "1", Title: "Book 1", Price: 29.99, Author: &Author{Id: "1", Name: "Author 1", Rating: 4.5}, Year: 2020, Rating: 4.5},
// 	// 	{Id: "2", Title: "Book 2", Price: 39.99, Author: &Author{Id: "2", Name: "Author 2", Rating: 4.0}, Year: 2021, Rating: 4.0},
// 	// }
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// // Handler to get all authors
// func getAuthors(w http.ResponseWriter, r *http.Request) {
// 	// authors := []Author{
// 	// 	{Id: "1", Name: "Author 1", Rating: 4.5},
// 	// 	{Id: "2", Name: "Author 2", Rating: 4.0},
// 	// }
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(authors)
// }

// // Handler to add a book
// func addBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	err = json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	book.Id = "3"
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }
