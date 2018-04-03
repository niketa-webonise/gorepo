package main

import ("fmt"
		"log"
		"net/http")


func allBooks(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Get request on books")
}


func handleRequests(){

	http.HandleFunc("/",homePage)
	http.HandleFunc("/books",allBooks)
	log.Fatal(http.ListenAndServe(":8084",nil))
}
func main() {

	handleRequests()
}
