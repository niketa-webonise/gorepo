package controllers

import ("net/http"
		"html/template"
		_ "github.com/go-sql-driver/mysql"
		)

type UserError struct {

	Msg string
}

type SgnError struct {

	Msg string
	Password string
	Firstname string
	Lastname string

}




func LoginError(w http.ResponseWriter,r *http.Request){

	if r.Method == "GET"{


		t,_ := template.ParseFiles("./view/loginerror.gtpl")

		message:=UserError{}

		message.Msg = "please enter correct email/password"

		t.ExecuteTemplate(w,"loginerror.gtpl",message)


	}
}


func DashboardError(w http.ResponseWriter,r *http.Request){

	if r.Method == "GET"{

		t,_ := template.ParseFiles("./view/dashboarderror.gtpl")

		dashboard_msg := UserError{}

		dashboard_msg.Msg = "no user logged in"

		t.ExecuteTemplate(w,"dashboarderror.gtpl",dashboard_msg)


	}}


	func SignupError(w http.ResponseWriter,r *http.Request){

	if r.Method == "GET"{

		t,_ := template.ParseFiles("./view/signuperror.gtpl")

		signup_msg := SgnError{}

		signup_msg.Msg = "please enter correct email"

		signup_msg.Password = "password doesn't match"

		signup_msg.Firstname = "firstname should contains alphabets only"

		signup_msg.Lastname = "lastname should contains alphabets only"


		t.ExecuteTemplate(w,"signuperror.gtpl",signup_msg)

		




	}
}