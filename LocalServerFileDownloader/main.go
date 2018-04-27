package main

import ("net/http"
		"log"
		"github.com/niketa/gorepo/LocalServerFileDownloader/routes")

func main(){

	routes.CreateRoute()

	
	log.Fatal(http.ListenAndServe(":8300",nil))
}
