package main

import ("fmt"
		"log"
		"net/http"
		"controllers")

func homePage(w http.ResponseWriter,r *http.Request) {

	fmt.Fprintf(w,"homepage end url")
}
func allBooks(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Get request on books")
}


func handleRequests(){

	http.HandleFunc("/",controllers.homePageController.go)
	http.HandleFunc("/books",controllers.allBooksController.go)
	log.Fatal(http.ListenAndServe(":8086",nil))
}
func main() {

	handleRequests()
}