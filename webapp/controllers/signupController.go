package controllers

import (
		"net/http"
		"html/template"
        "strings"
        "regexp"
        "github.com/niketa/gorepo/webapp/model"
        )




func SignUpPage(w http.ResponseWriter,r *http.Request){


    //get request method
    if r.Method == "GET" {
    	        t, _ := template.ParseFiles("./view/signup.gtpl")
        t.Execute(w, nil)
        
        

    } else if r.Method == "POST" {

       
        r.ParseForm()

        firstname := r.FormValue("firstname")
        lastname := r.FormValue("lastname")
        email := r.FormValue("email")
        password := r.FormValue("psw")
        repassword := r.FormValue("psw-repeat")

         

        var fname = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

        var lname = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
        
        if m, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,email);
           
            !m {
                http.Redirect(w,r,"/signuperror",301)
            }else{

                if strings.Compare(password,repassword)!= 0 {
                http.Redirect(w,r,"/signuperror",301)

        }else{

            if fname(firstname) && lname(lastname){

                model.CreateUsers(email,password,firstname,lastname)

                http.Redirect(w,r,"/login",301)

            }else{

                http.Redirect(w,r,"/signuperror",301)
            }

        }
    }

}
}
