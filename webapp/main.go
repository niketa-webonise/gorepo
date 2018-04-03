package main


import ("log"
		"net/http"
		"github.com/niketa/gorepo/webapp/routes"
		"github.com/niketa/gorepo/webapp/database"
		)



func main() {

	database.InitialiseDB()

	routes.CreateRoute()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Fatal(http.ListenAndServe(":8080",nil))
}