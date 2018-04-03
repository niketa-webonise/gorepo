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
        
        

    } else if r.Method == "POST"{

       
        r.ParseForm()


        email := r.FormValue("email")
        password := r.FormValue("psw")
        repassword := r.FormValue("psw-repeat")
      

           if m, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,email);
        

           !m {

            //fmt.Println("no insertion")

            http.Redirect(w,r,"/signuperror",301)
                }else{

                        if strings.Compare(password,repassword)!= 0 {

                           
                            http.Redirect(w,r,"/signuperror",301)
                        }else{

                             
                            model.CreateUsers(email,password)

                            http.Redirect(w,r,"/login",301)

                        }
}
}
}

