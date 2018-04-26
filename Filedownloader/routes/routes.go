package routes

import ("github.com/niketa/gorepo/download/controllers"
		"net/http")

func CreateRoute(){

	http.HandleFunc("/download",controllers.FileURL)
}