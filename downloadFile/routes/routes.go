package routes

import ("net/http"
		"github.com/niketa/gorepo/downloadFile/controllers")

func CreateRoute(){

	http.HandleFunc("/download",controllers.FileURL)
}