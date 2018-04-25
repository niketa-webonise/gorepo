package main

import ("net/http" 
	    "log"
	    "github.com/niketa/gorepo/downloadFile/routes")

func main() {

	routes.CreateRoute()

	log.Fatal(http.ListenAndServe(":8080",nil))

    
}