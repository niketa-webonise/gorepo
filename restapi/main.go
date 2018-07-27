package main

import ("log"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux")

//Book struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books[] Book

//get all books
func getBooks(w http.ResponseWriter,r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//get a single book
func getBook(w http.ResponseWriter,r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	//loop through books and find ID
	for _,item := range books {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}


	}

	json.NewEncoder(w).Encode(&Book{})
}

//create a single book
func createBook(w http.ResponseWriter,r *http.Request) {

}

//update a single book
func updateBook(w http.ResponseWriter,r *http.Request) {

}

//delete a single book
func deleteBook(w http.ResponseWriter,r *http.Request) {

}





func main() {

router := mux.NewRouter()

books = append(books,Book{ID:"1",Isbn:"12345",Title:"MAGIC",Author:&Author{Firstname:"Rohnde",Lastname:"Bynre"}})
books = append(books,Book{ID:"2",Isbn:"6789",Title:"SECRET",Author:&Author{Firstname:"Rohnde",Lastname:"Bynre"}})


router.HandleFunc("/api/books",getBooks).Methods("GET")
router.HandleFunc("/api/book/{id}",getBook).Methods("GET")
router.HandleFunc("/api/book/{id}",createBook).Methods("POST")
router.HandleFunc("/api/book/{id}",updateBook).Methods("PUT")
router.HandleFunc("/api/book/{id}",deleteBook).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8500",router))
}