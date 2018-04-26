package routes

import ("github.com/niketa/gorepo/file/controllers"
		"net/http")

func CreateRoute(){

	http.HandleFunc("/download",controllers.FileURL)
}