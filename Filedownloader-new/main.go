package main

import ("net/http"
		"log"
		"github.com/niketa/gorepo/download/routes")

func main(){

	routes.CreateRoute()


	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	
	log.Fatal(http.ListenAndServe(":8050",nil))
}
