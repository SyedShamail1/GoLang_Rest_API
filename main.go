package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Books Struct
//(A Slice Book Struct is a variable length array
//i.e the length of the array is flexible and not defined on initialization)
var Books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	//Loop through books and find with id
	for _, item := range Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	//If The given ID record does not exist return a new Book Struct with the Title with value "DEFAULT"
	json.NewEncoder(w).Encode(&Book{Title: "DEFAULT"})

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

	//Setting Header content type
	w.Header().Set("Content-Type", "application/json")

	//New Book Struct Type Variable for insertion
	var book Book

	//Decoding the json request body values
	_ = json.NewDecoder(r.Body).Decode(&book)

	//Setting the iD randomly with "rand.Intn" function and converting it to string with "strconv.Itoa" function
	//as it will not be provided in the request
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID not safe

	//Adding the book in the declared Book Struct with Append function
	Books = append(Books, book)

	//Returning the newly added book as a response
	json.NewEncoder(w).Encode(book)
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	//Setting Header content type
	w.Header().Set("Content-Type", "application/json")

	//Get the id with paranmeters
	params := mux.Vars(r)

	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			//New Book Struct Type Variable for insertion
			var book Book

			//Decoding the json request body values
			_ = json.NewDecoder(r.Body).Decode(&book)

			//Setting the iD randomly with "rand.Intn" function and converting it to string with "strconv.Itoa" function
			//as it will not be provided in the request
			book.ID = params["id"]

			//Adding the book in the declared Book Struct with Append function
			Books = append(Books, book)

			//Returning the newly added book as a response
			json.NewEncoder(w).Encode(book)
			return

		}

	}

}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//Setting Header content type
	w.Header().Set("Content-Type", "application/json")

	//Get the id with paranmeters
	params := mux.Vars(r)

	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}

	}

	//Returning the list of books
	json.NewEncoder(w).Encode(Books)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	Books = append(Books, Book{ID: "1", Isbn: "122424", Title: "Flowers For Algernon",
		Author: &Author{Firstname: "Daniel", Lastname: "Keys"}})

	Books = append(Books, Book{ID: "2", Isbn: "1255125", Title: "By Way Of Deception",
		Author: &Author{Firstname: "Simon", Lastname: "Lahav"}})

	//Route handler/ End points
	r.HandleFunc("/api/Books", getBooks).Methods("GET")
	r.HandleFunc("/api/Books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/Books", createBook).Methods("POST")
	r.HandleFunc("/api/Books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/Books/{id}", deleteBook).Methods("DELETE")

	//Server for listening
	log.Fatal(http.ListenAndServe(":8000", r))

}
