package controllers

import ("net/http"
		"html/template"
		_ "github.com/go-sql-driver/mysql"
		"github.com/niketa/gorepo/webapp/model"
		)

type User struct{
	Email string
}
func DashboardPage(w http.ResponseWriter,r *http.Request){


	if r.Method == "GET"{

if model.CurrentUserEmail!=" " {

t,_ := template.ParseFiles("./view/dashboard.gtpl")


email := model.CurrentUserEmail

usrEID := User{}
usrEID.Email = email

t.ExecuteTemplate(w,"dashboard.gtpl",usrEID)


} else {
http.Redirect(w,r,"/dashboarderror",301)
}
}
}