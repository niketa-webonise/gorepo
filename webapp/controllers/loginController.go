package controllers

import (
		"net/http"
		"html/template"
		_ "github.com/go-sql-driver/mysql"
		"github.com/niketa/gorepo/webapp/model"

        )


func LoginPage(w http.ResponseWriter,r *http.Request) {

	
	if r.Method == "GET" {

		t,_ := template.ParseFiles("./view/login.gtpl")
		t.Execute(w,nil)

	} else if r.Method == "POST"{

			
		r.ParseForm()

 		email := r.FormValue("email")
 		password := r.FormValue("pwd")

 		if model.CheckUsers(email,password) {

 			model.CurrentUserEmail =  email
	        http.Redirect(w,r,"/dashboard",301)
		
		} else {
			
			http.Redirect(w,r,"/loginerror",301)
		
		}
	}
}


func Logout(w http.ResponseWriter,r *http.Request) {


if r.Method == "POST"{

model.CurrentUserEmail = " "

http.Redirect(w,r,"/login",301)


}
}

