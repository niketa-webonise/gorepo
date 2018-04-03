package main

import ("log"
		"net/http"
		"./controllers")

func handleRequests(){
	http.HandleFunc("/",controllers.HomePage)
	http.HandleFunc("/articles",controllers.AllArticles)
	log.Fatal(http.ListenAndServe(":8088",nil))
}
func main() {

	handleRequests()
}