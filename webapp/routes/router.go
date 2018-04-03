package routes

import ("github.com/niketa/gorepo/webapp/controllers"
		"net/http")

func CreateRoute() {



		http.HandleFunc("/signup", controllers.SignUpPage)
		http.HandleFunc("/login",controllers.LoginPage)
		http.HandleFunc("/dashboard",controllers.DashboardPage)
		http.HandleFunc("/loginerror",controllers.LoginError)
		http.HandleFunc("/dashboard_error",controllers.DashboardError)
		http.HandleFunc("/signuperror",controllers.SignupError)

}