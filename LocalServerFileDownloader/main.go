package main

import ("net/http"
		"log"
		"github.com/niketa/gorepo/file/routes")

func main(){

	routes.CreateRoute()

	
	log.Fatal(http.ListenAndServe(":8500",nil))
}
