package main


import ("log"
		"net/http"
		"./controllers")


func handleRequests(){

	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/books",controllers.AllBooks)
	log.Fatal(http.ListenAndServe(":8686",nil))
}
func main() {

	handleRequests()
}
