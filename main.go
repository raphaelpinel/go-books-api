package main

import (
	// "encoding/json"
	"log"
	"net/http"
)

type server struct {}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "PUT called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "DELETE called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
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
	// define routes
	// http.HandleFunc("/books", getBooks)
	// http.HandleFunc("/authors", getAuthors)
	// http.HandleFunc("/books/add", addBook)
	// http.HandleFunc("/authors/add", addAuthor)

	// start server
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
