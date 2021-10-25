package main

//importing packages for working with json, logging errors, working with http, creating ids as random numbers, converting types to the string
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//we will add third party package called mux, for routes

//Book struct(Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as slice Book struct

var books []Book

//Get all books(every route handling function must get request and response)

func getBooks(w http.ResponseWriter, r *http.Request) {
	//Heder value of content-type will be set to json, because will be sent as text elsewhere
	w.Header().Set("Content-Type", "application/json")
	//w reposne variable will be books slice
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params

	//Loop throug books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	//strcon convertin integer to a string, and getting random number between 0-10000000
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID, because can generate same ID multiple times
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "444332", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "555432", Title: "Book Two", Author: &Author{Firstname: "Jim", Lastname: "Bach"}})
	books = append(books, Book{ID: "3", Isbn: "6464737", Title: "Book Three", Author: &Author{Firstname: "Jane", Lastname: "Kao"}})

	//Creating rout handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//Creating server which will listen for requests
	log.Fatal(http.ListenAndServe(":8000", r))

}
