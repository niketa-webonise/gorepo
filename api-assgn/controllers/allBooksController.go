package controllers

import ("net/http"
		"fmt")


func AllBooks(w http.ResponseWriter,r *http.Request){

	fmt.Fprintf(w,"Get request on books")
}
