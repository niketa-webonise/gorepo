package routes

import ("github.com/niketa/gorepo/LocalServerFileDownloader/controllers"
		"net/http")

func CreateRoute(){

	http.HandleFunc("/download",controllers.FileURL)
}